package dentist

import (
	"github.com/DamianSformo/proyecto-entregable-final/internal/domain"
)

type Service interface {
	GetDentistById(id int) (domain.Dentist, error)
	CreateDentist(p domain.Dentist) (domain.Dentist, error)
	UpdateDentist(id int, p domain.Dentist) (domain.Dentist, error)
	DeleteDentist(id int) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (service *service) GetDentistById(id int) (domain.Dentist, error) {
	dentist, err := service.repository.GetDentistById(id)
	if err != nil {
		return domain.Dentist{}, err
	}
	return dentist, nil
}


func (service *service) CreateDentist(d domain.Dentist) (domain.Dentist, error) {
	d, err := service.repository.CreateDentist(d)
	if err != nil {
		return domain.Dentist{}, err
	}
	return d, nil
}


func (service *service) UpdateDentist(id int, d domain.Dentist) (domain.Dentist, error) {

	dentist, err := service.repository.GetDentistById(id)

	if err != nil {
		return domain.Dentist{}, err
	}
	if d.Name != "" {
		dentist.Name = d.Name
	}
	if d.Surname != "" {
		dentist.Surname = d.Surname
	}
	if d.License != "" {
		dentist.License = d.License 
	}
	
	dentist, err = service.repository.UpdateDentist(id, dentist)
	if err != nil {
		return domain.Dentist{}, err
	}

	return dentist, nil
}

func (service *service) DeleteDentist(id int) error {
	err := service.repository.DeleteDentist(id)
	if err != nil {
		return err
	}
	return nil
}
