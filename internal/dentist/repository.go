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

func (repository *repository) GetDentistById(id int) (domain.Dentist, error) {

	if !repository.storage.ExistsDentist(id) {
		return domain.Dentist{}, errors.New("There is no dentist with this Id")
	}

	dentist, err := repository.storage.GetDentistById(id)
	if err != nil {
		return domain.Dentist{}, errors.New("Not found")
	}

	return dentist, nil

}

func (repository *repository) GetDentistByLicense(license string) (domain.Dentist, error) {

	dentist, err := repository.storage.GetDentistByLicense(license)
	if err != nil {
		return domain.Dentist{}, err
	}
	return dentist, nil
}

func (repository *repository) CreateDentist(d domain.Dentist) (domain.Dentist, error) {
	
	_, err := repository.storage.GetDentistByLicense(d.License) 

	if err == nil {
		return domain.Dentist{}, errors.New("Exist dentist with this license")
	} 
	
	id, err := repository.storage.CreateDentist(d)
	if err != nil {
		return domain.Dentist{}, errors.New("Error creating product")
	}

	d.Id = int(id) 

	return d, nil
}

func (repository *repository) UpdateDentist(id int, d domain.Dentist) (domain.Dentist, error) {

	if !repository.storage.ExistsDentist(id) {
		return domain.Dentist{}, errors.New("There is no dentist with this Id") 
	}

	dres, err := repository.storage.GetDentistByLicense(d.License) 

	if dres.Id != 0 && dres.Id != d.Id  {
		return domain.Dentist{}, errors.New("Exist dentist with this license")
	}

	err = repository.storage.Update(d, id)
	if err != nil {
		return domain.Dentist{}, errors.New("Error updating product")
	}
	return d, nil
}

func (repository *repository) DeleteDentist(id int) error {

	if !repository.storage.ExistsDentist(id) {
		return errors.New("There is no dentist with this Id")
	}

	err := repository.storage.DeleteDentist(id)
	if err != nil {
		return err
	}
	return nil
}