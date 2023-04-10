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

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetDentistById(id int) (domain.Dentist, error) {
	p, err := s.repository.GetDentistById(id)
	if err != nil {
		return domain.Dentist{}, err
	}
	return p, nil
}


func (s *service) CreateDentist(d domain.Dentist) (domain.Dentist, error) {
	d, err := s.repository.CreateDentist(d)
	if err != nil {
		return domain.Dentist{}, err
	}
	return d, nil
}


func (s *service) UpdateDentist(id int, d domain.Dentist) (domain.Dentist, error) {
	p, err := s.repository.GetDentistById(id)
	if err != nil {
		return domain.Dentist{}, err
	}
	if d.Name != "" {
		p.Name = d.Name
	}
	if d.Surname != "" {
		p.Surname = d.Surname
	}
	if d.License != "" {
		p.License = d.License 
	}
	
	p, err = s.repository.UpdateDentist(id, p)
	if err != nil {
		return domain.Dentist{}, err
	}
	return p, nil
}

func (s *service) DeleteDentist(id int) error {
	err := s.repository.DeleteDentist(id)
	if err != nil {
		return err
	}
	return nil
}
