package appointment

import (
	"errors"
	"github.com/DamianSformo/proyecto-entregable-final/internal/domain"
	"github.com/DamianSformo/proyecto-entregable-final/pkg/appointment_store"
)

type Repository interface {
	CreateAppointment(p domain.Appointment) (domain.Appointment, error)
	GetAppointmentByDni(dni int) ([]domain.Appointment, error)
}

type repository struct {
	storage appointment_store.StoreInterface
}

func NewRepository(storage appointment_store.StoreInterface) Repository {
	return &repository{storage}
}

func (r *repository) GetAppointmentByDni(dni int) ([]domain.Appointment, error) {

	appointments, err := r.storage.GetAppointmentByDni(dni)
	if err != nil {
		return []domain.Appointment{}, err
	}
	return appointments, nil
}

func (r *repository) CreateAppointment(a domain.Appointment) (domain.Appointment, error) {
	id, err := r.storage.CreateAppointment(a)
	if err != nil {
		return domain.Appointment{}, errors.New("error creating appointment")
	}

	a.Id = int(id)

	return a, nil
} 
