package config

import (
	"net/http"
	"strings"

	"go.uber.org/zap"
)

var Log *zap.Logger

func InitLogger() {
	Log, _ = zap.NewProduction()
}

func LogAndRespond(w http.ResponseWriter, r *http.Request, msg string, code int, level string) {
	logFields := []zap.Field{
		zap.String("path", r.URL.Path),
		zap.String("method", r.Method),
		zap.Int("status", code),
		zap.String("clientIP", getClientAddress(r)),
	}

	switch level {
	case "DEBUG":
		Log.Debug(msg, logFields...)
	case "ERROR":
		Log.Error(msg, logFields...)
	default:
		Log.Info(msg, logFields...)
	}

	http.Error(w, msg, code)
}

func getClientAddress(r *http.Request) string {
	//if found in headers
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		return strings.Split(xff, ",")[0]
	}
	//fallback: direct connection
	return strings.Split(r.RemoteAddr, ":")[0]
}
