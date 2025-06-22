package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/thebigmatchplayer/markerble-task/config"
	"github.com/thebigmatchplayer/markerble-task/models"
)

func CreatePatientHandler(w http.ResponseWriter, r *http.Request) {
	config.Log.Debug("reached CreatePatientHandler")
	var p models.Patient
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// validate struct (optional)
	if err := validate.Struct(p); err != nil {
		http.Error(w, "Validation failed", http.StatusBadRequest)
		return
	}

	// check doctor ID is valid
	if ok, _ := models.IsValidDoctorID(p.DoctorID); !ok {
		http.Error(w, "Invalid doctor ID", http.StatusBadRequest)
		return
	}

	if err := models.CreatePatient(&p); err != nil {
		http.Error(w, "Could not create patient", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(p)
}

func DeletePatientHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing patient ID", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid patient ID", http.StatusBadRequest)
		return
	}

	if err := models.DeletePatient(id); err != nil {
		http.Error(w, "Failed to delete patient", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent) // 204 No Content
}
