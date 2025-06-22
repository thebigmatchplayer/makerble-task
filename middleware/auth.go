package middleware

import (
	"net/http"
	"slices"
	"strings"

	"github.com/thebigmatchplayer/markerble-task/config"
	"github.com/thebigmatchplayer/markerble-task/utils"
	"go.uber.org/zap"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		config.Log.Debug("reached /patients")
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			config.LogAndRespond(w, r, "Missing or invalid Authorization header", http.StatusUnauthorized, "DEBUG")
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := utils.ParseJWT(tokenStr)
		if err != nil {
			config.LogAndRespond(w, r, "Invalid Token", http.StatusUnauthorized, "DEBUG")
			return
		}

		// ABAC check
		methodMap, ok := AccessMatrix[r.URL.Path]
		if !ok {
			config.LogAndRespond(w, r, "Invalid Path", http.StatusNotFound, "DEBUG")
			return
		}

		roles, ok := methodMap[r.Method]
		if !ok {
			config.LogAndRespond(w, r, "Method Not Allowed", http.StatusMethodNotAllowed, "DEBUG")
			return
		}

		if !slices.Contains(roles, claims.Role) {
			config.LogAndRespond(w, r, "Forbidden", http.StatusForbidden, "DEBUG")
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		config.Log.Info("Access granted",
			zap.String("path", r.URL.Path),
			zap.String("method", r.Method),
			zap.String("role", claims.Role),
			zap.Int("user_id", claims.UserID),
		)

		next(w, r)
	}
}
