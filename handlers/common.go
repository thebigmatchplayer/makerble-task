package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/thebigmatchplayer/markerble-task/models"
)

func GetAllPatientsHandler(w http.ResponseWriter, r *http.Request) {
	patients, err := models.GetAllPatients()
	if err != nil {
		http.Error(w, "Failed to retrieve patients", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(patients)
}
