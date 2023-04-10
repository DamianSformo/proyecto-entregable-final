package dentist

import (
	"errors"
	"github.com/DamianSformo/proyecto-entregable-final/internal/domain"
	"github.com/DamianSformo/proyecto-entregable-final/pkg/dentist_store"
)

type Repository interface {
	GetDentistById(id int) (domain.Dentist, error)
	GetDentistByLicense(license string) (domain.Dentist, error)
	CreateDentist(p domain.Dentist) (domain.Dentist, error)
	UpdateDentist(id int, p domain.Dentist) (domain.Dentist, error)
	DeleteDentist(id int) error
}

type repository struct {
	storage dentist_store.StoreInterface
}

func NewRepository(storage dentist_store.StoreInterface) Repository {
	return &repository{storage}
}

func (r *repository) GetDentistById(id int) (domain.Dentist, error) {

	if !r.storage.ExistsDentist(id) {
		return domain.Dentist{}, errors.New("There is no dentist with this Id")
	}

	dentist, err := r.storage.GetDentistById(id)
	if err != nil {
		return domain.Dentist{}, errors.New("Not found")
	}
	return dentist, nil

}

func (r *repository) GetDentistByLicense(license string) (domain.Dentist, error) {

	dentist, err := r.storage.GetDentistByLicense(license)
	if err != nil {
		return domain.Dentist{}, err
	}
	return dentist, nil
}

func (r *repository) CreateDentist(d domain.Dentist) (domain.Dentist, error) {
	
	_, err := r.storage.GetDentistByLicense(d.License) 
	if err == nil {
		return domain.Dentist{}, errors.New("exist dentist with this license")
	} 
	
	id, err := r.storage.CreateDentist(d)
	if err != nil {
		return domain.Dentist{}, errors.New("error creating product")
	}

	d.Id = int(id) 

	return d, nil
}

func (r *repository) UpdateDentist(id int, d domain.Dentist) (domain.Dentist, error) {

	if !r.storage.ExistsDentist(id) {
		return domain.Dentist{}, errors.New("There is no dentist with this Id")
	}

	dres, err := r.storage.GetDentistByLicense(d.License) 
	if dres.Id != d.Id {
		return domain.Dentist{}, errors.New("exist dentist with this license")
	}

	err = r.storage.Update(d, id)
	if err != nil {
		return domain.Dentist{}, errors.New("error updating product")
	}
	return d, nil
}

func (r *repository) DeleteDentist(id int) error {
	err := r.storage.DeleteDentist(id)
	if err != nil {
		return err
	}
	return nil
}