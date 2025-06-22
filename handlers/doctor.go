package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/thebigmatchplayer/markerble-task/models"
)

func UpdatePatientHandler(w http.ResponseWriter, r *http.Request) {
	var p models.Patient
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := validate.Struct(p); err != nil {
		http.Error(w, "Validation failed", http.StatusBadRequest)
		return
	}

	if ok, _ := models.IsValidDoctorID(p.DoctorID); !ok {
		http.Error(w, "Invalid doctor ID", http.StatusBadRequest)
		return
	}

	if err := models.UpdatePatient(&p); err != nil {
		http.Error(w, "Failed to update patient", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(p)
}

func GetPatientByIDHandler(w http.ResponseWriter, r *http.Request) {
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

	patient, err := models.GetPatientByID(id)
	if err != nil {
		http.Error(w, "Patient not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(patient)
}
