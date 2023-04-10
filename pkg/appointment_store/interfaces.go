package appointment_store

import "github.com/DamianSformo/proyecto-entregable-final/internal/domain"

type StoreInterface interface {
	//GetPatientById(id int) (domain.Patient, error)
	GetAppointmentByDni(id int) ([]domain.Appointment, error)
	CreateAppointment(appointment domain.Appointment) (int64, error)
	//UpdatePatient(p domain.Patient, id int) error
	//ExistsPatient(id int) bool
	//DeletePatient(id int) error
}
