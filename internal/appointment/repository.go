package appointment

import (
	"errors"
	"github.com/DamianSformo/proyecto-entregable-final/internal/domain"
	"github.com/DamianSformo/proyecto-entregable-final/pkg/appointment_store"
)

type Repository interface {
	GetAppointmentById(id int) (domain.Appointment, error)
	CreateAppointment(p domain.Appointment) (domain.Appointment, error)
	GetAppointmentByDni(dni int) ([]domain.Appointment, error)
	DeleteAppointment(id int) error
}

type repository struct {
	storage appointment_store.StoreInterface
}

func NewRepository(storage appointment_store.StoreInterface) Repository {
	return &repository{storage}
}


func (repository *repository) GetAppointmentById(id int) (domain.Appointment, error) {
	 
	if !repository.storage.ExistsAppointment(id) {
		return domain.Appointment{}, errors.New("There is no appointment with this Id")
	} 

	appointment, err := repository.storage.GetAppointmentById(id)
	if err != nil {
		return domain.Appointment{}, err
	}
	return appointment, nil
}

func (repository *repository) GetAppointmentByDni(dni int) ([]domain.Appointment, error) {

	appointments, err := repository.storage.GetAppointmentByDni(dni)
	if err != nil {
		return []domain.Appointment{}, err
	}
	return appointments, nil
}

func (repository *repository) CreateAppointment(a domain.Appointment) (domain.Appointment, error) {
	id, err := repository.storage.CreateAppointment(a)
	if err != nil {
		return domain.Appointment{}, errors.New("Error creating appointment")
	}

	a.Id = int(id)

	return a, nil
} 

func (repository *repository) DeleteAppointment(id int) error {
	
	if !repository.storage.ExistsAppointment(id) {
		return errors.New("There is no appointment with this Id")
	} 

	err := repository.storage.DeleteAppointment(id)
	if err != nil {
		return err
	}
	return nil
}
