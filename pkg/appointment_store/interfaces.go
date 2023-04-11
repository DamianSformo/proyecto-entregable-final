package appointment_store

import "github.com/DamianSformo/proyecto-entregable-final/internal/domain"

type StoreInterface interface {
	GetAppointmentById(id int) (domain.Appointment, error)
	GetAppointmentByDni(dni int) ([]domain.Appointment, error)
	CreateAppointment(appointment domain.Appointment) (int64, error)
	//UpdatePatient(p domain.Patient, id int) error 
	DeleteAppointment(id int) error
	ExistsAppointment(id int) bool
}
