package patient_store

import "github.com/DamianSformo/proyecto-entregable-final/internal/domain"

type StoreInterface interface {
	GetPatientById(id int) (domain.Patient, error)
	GetPatientByDni(dni int) (domain.Patient, error)
	CreatePatient(patient domain.Patient) (int64, error)
	UpdatePatient(p domain.Patient, id int) error 
	ExistsPatient(id int) bool
	DeletePatient(id int) error
}
