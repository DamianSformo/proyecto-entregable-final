package patient

import (
	"time"
	"github.com/DamianSformo/proyecto-entregable-final/internal/domain"
)

type Service interface {
	GetPatientByID(id int) (domain.Patient, error)
	GetPatientByDni(dni int) (domain.Patient, error)
	CreatePatient(p domain.Patient) (domain.Patient, error) 
	DeletePatient(id int) error
	UpdatePatient(p domain.Patient, id int) (domain.Patient, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetPatientByID(id int) (domain.Patient, error) {
	p, err := s.repository.GetPatientByID(id) 
	if err != nil {
		return domain.Patient{}, err
	}
	return p, nil
}

func (s *service) GetPatientByDni(dni int) (domain.Patient, error) {
	p, err := s.repository.GetPatientByDni(dni) 
	if err != nil {
		return domain.Patient{}, err
	}
	return p, nil
}

func (s *service) CreatePatient(p domain.Patient) (domain.Patient, error) {
	t := time.Now()
	p.Date = t.Format("2006-01-02")

	p, err := s.repository.CreatePatient(p)
	if err != nil {
		return domain.Patient{}, err
	}

	return p, nil
}

func (s *service) UpdatePatient(p domain.Patient, id int) (domain.Patient, error) {
	pat, err := s.repository.GetPatientByID(id)
	if err != nil {
		return domain.Patient{}, err
	}

	if pat.Date != "" {
		p.Date = pat.Date 
	}
	
	pat, err = s.repository.UpdatePatient(p, id)
	if err != nil {
		return domain.Patient{}, err
	}

	pat.Id = id
	
	return pat, nil
}

func (s *service) DeletePatient(id int) error  {
	err := s.repository.DeletePatient(id)
	if err != nil {
		return err
	}
	return nil
} 
