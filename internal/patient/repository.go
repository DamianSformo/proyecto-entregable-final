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
	DeletePatient(id int) error
	UpdatePatient(p domain.Patient, id int) (domain.Patient, error)
}

type repository struct {
	storage patient_store.StoreInterface
}

func NewRepository(storage patient_store.StoreInterface) Repository {
	return &repository{storage}
}

func (r *repository) GetPatientByID(id int) (domain.Patient, error) {
	 
	if !r.storage.ExistsPatient(id) {
		return domain.Patient{}, errors.New("There is no patient with this Id")
	}

	product, err := r.storage.GetPatientById(id)
	if err != nil {
		return domain.Patient{}, err
	}
	return product, nil
}

func (r *repository) GetPatientByDni(dni int) (domain.Patient, error) {

	product, err := r.storage.GetPatientByDni(dni)
	if err != nil {
		return domain.Patient{}, err
	}
	return product, nil
}

func (r *repository) CreatePatient(p domain.Patient) (domain.Patient, error) {

	_, err := r.storage.GetPatientByDni(p.DNI) 
	if err == nil {
		return domain.Patient{}, errors.New("exist patient with this dni")
	}

	id, err := r.storage.CreatePatient(p)
	if err != nil {
		return domain.Patient{}, errors.New("error creating product")
	} 

	p.Id = int(id)

	return p, nil
}  

func (r *repository) UpdatePatient(p domain.Patient, id int) (domain.Patient, error) {

	if !r.storage.ExistsPatient(id) {
		return domain.Patient{}, errors.New("code value already exists")
	}

	err := r.storage.UpdatePatient(p, id)
	if err != nil {
		return domain.Patient{}, errors.New("error updating product")
	}
	return p, nil
}

func (r *repository) DeletePatient(id int) error {
	err := r.storage.DeletePatient(id)
	if err != nil {
		return err
	}
	return nil
}