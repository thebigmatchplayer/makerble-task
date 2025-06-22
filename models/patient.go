package models

import (
	"time"

	"github.com/thebigmatchplayer/markerble-task/config"
)

type Patient struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" validate:"required"`
	Age       int       `json:"age" validate:"gte=0,lte=130"`
	Gender    string    `json:"gender" validate:"oneof=male female other"`
	Diagnosis string    `json:"diagnosis"`
	DoctorID  int       `json:"doctor_id" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
}

func CreatePatient(p *Patient) error {
	query := `INSERT INTO patients (name, age, gender, diagnosis, doctor_id)
		          VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at`
	return config.DB.QueryRow(query, p.Name, p.Age, p.Gender, p.Diagnosis, p.DoctorID).Scan(&p.ID, &p.CreatedAt)
}

func IsValidDoctorID(doctorID int) (bool, error) {
	query := `SELECT EXISTS (
			SELECT 1 FROM users WHERE id = $1 AND role = 'doctor'
		)`
	var exists bool
	err := config.DB.QueryRow(query, doctorID).Scan(&exists)
	return exists, err
}

func GetPatientByID(id int) (*Patient, error) {
	query := `SELECT id, name, age, gender, diagnosis, doctor_id, created_at FROM patients WHERE id = $1`
	var p Patient
	err := config.DB.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Age, &p.Gender, &p.Diagnosis, &p.DoctorID, &p.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func UpdatePatient(p *Patient) error {
	query := `UPDATE patients SET name = $1, age = $2, gender = $3, diagnosis = $4, doctor_id = $5 WHERE id = $6`
	_, err := config.DB.Exec(query, p.Name, p.Age, p.Gender, p.Diagnosis, p.DoctorID, p.ID)
	return err
}

func DeletePatient(id int) error {
	query := `DELETE FROM patients WHERE id = $1`
	_, err := config.DB.Exec(query, id)
	return err
}

func GetAllPatients() ([]Patient, error) {
	rows, err := config.DB.Query(`SELECT id, name, age, gender, diagnosis, doctor_id, created_at FROM patients`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var patients []Patient
	for rows.Next() {
		var p Patient
		if err := rows.Scan(&p.ID, &p.Name, &p.Age, &p.Gender, &p.Diagnosis, &p.DoctorID, &p.CreatedAt); err != nil {
			return nil, err
		}
		patients = append(patients, p)
	}
	return patients, nil
}
