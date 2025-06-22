package handlers

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/thebigmatchplayer/markerble-task/config"
	"github.com/thebigmatchplayer/markerble-task/middleware"
)

var validate = validator.New()

func SetupRoutes() {
	http.HandleFunc("/patients", middleware.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		config.Log.Debug("reached /patients")
		switch r.Method {
		case "POST":
			CreatePatientHandler(w, r)
		case "GET":
			id := r.URL.Query().Get("id")
			if id != "" {
				GetPatientByIDHandler(w, r)
			} else {
				GetAllPatientsHandler(w, r)
			}
		case "PUT":
			UpdatePatientHandler(w, r)
		case "DELETE":
			DeletePatientHandler(w, r)
		default:
			config.Log.Warn("Method not allowed")
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}))
	http.HandleFunc("/register", RegisterHandler)
	http.HandleFunc("/login", LoginHandler)

}
