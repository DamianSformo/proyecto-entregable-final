package patient

import (
	"errors"
	"github.com/DamianSformo/proyecto-entregable-final/internal/domain"
	"github.com/DamianSformo/proyecto-entregable-final/pkg/patient_store"
)

type Repository interface {
	CreatePatient(p domain.Patient) (domain.Patient, error)
	GetPatientByID(id int) (domain.Patient, error)
	GetPatientByDni(dni int) (domain.Patient, error)
	UpdatePatient(id int, p domain.Patient)  (domain.Patient, error)
	DeletePatient(id int) error
}

type repository struct {
	storage patient_store.StoreInterface
}

func NewRepository(storage patient_store.StoreInterface) Repository {
	return &repository{storage}
}

func (repository *repository) GetPatientByID(id int) (domain.Patient, error) {
	 
	if !repository.storage.ExistsPatient(id) {
		return domain.Patient{}, errors.New("There is no patient with this Id")
	}

	patient, err := repository.storage.GetPatientById(id)
	if err != nil {
		return domain.Patient{}, err
	}
	return patient, nil
}

func (repository *repository) GetPatientByDni(dni int) (domain.Patient, error) {

	patient, err := repository.storage.GetPatientByDni(dni)
	if err != nil {
		return domain.Patient{}, err
	}
	return patient, nil
}

func (repository *repository) CreatePatient(p domain.Patient) (domain.Patient, error) {

	_, err := repository.storage.GetPatientByDni(p.DNI) 

	if err == nil {
		return domain.Patient{}, errors.New("Exist patient with this dni")
	}

	id, err := repository.storage.CreatePatient(p)
	if err != nil {
		return domain.Patient{}, errors.New("Error creating patient")
	} 

	p.Id = int(id)

	return p, nil
}  

func (repository *repository) UpdatePatient(id int, p domain.Patient) (domain.Patient, error) {

	if !repository.storage.ExistsPatient(id) {
		return domain.Patient{}, errors.New("There is no patient with this Id") 
	}

	pres, err := repository.storage.GetPatientByDni(p.DNI) 

	if pres.Id != 0 && pres.Id != id {
		return domain.Patient{}, errors.New("Exist dentist with this license")
	}

	err = repository.storage.UpdatePatient(p, id)
	if err != nil {
		return domain.Patient{}, errors.New("Error updating patient")
	}

	p.Id = id

	return p, nil
}

func (r *repository) DeletePatient(id int) error {

	if !r.storage.ExistsPatient(id) {
		return errors.New("There is no patient with this Id")
	}

	err := r.storage.DeletePatient(id)
	if err != nil {
		return err
	}
	return nil
}