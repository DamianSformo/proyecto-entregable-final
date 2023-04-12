package patient

import (
	"time"
	"github.com/DamianSformo/proyecto-entregable-final/internal/domain"
)

type Service interface {
	GetPatientByID(id int) (domain.Patient, error)
	GetPatientByDni(dni int) (domain.Patient, error)
	CreatePatient(p domain.Patient) (domain.Patient, error) 
	UpdatePatient(p domain.Patient, id int) (domain.Patient, error)
	DeletePatient(id int) error

}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (service *service) GetPatientByID(id int) (domain.Patient, error) {
	patient, err := service.repository.GetPatientByID(id) 
	if err != nil {
		return domain.Patient{}, err
	}
	return patient, nil
}

func (service *service) GetPatientByDni(dni int) (domain.Patient, error) {
	patient, err := service.repository.GetPatientByDni(dni) 
	if err != nil {
		return domain.Patient{}, err
	}
	return patient, nil
}

func (service *service) CreatePatient(p domain.Patient) (domain.Patient, error) {
	t := time.Now()
	p.Date = t.Format("2006-01-02")

	p, err := service.repository.CreatePatient(p)
	if err != nil {
		return domain.Patient{}, err
	}

	return p, nil
}

func (service *service) UpdatePatient(p domain.Patient, id int) (domain.Patient, error) {

	patient, err := service.repository.GetPatientByID(id)

	if err != nil {
		return domain.Patient{}, err
	}

	if patient.Date != "" {
		p.Date = patient.Date 
	}
	
	patient, err = service.repository.UpdatePatient(id, p)
	if err != nil {
		return domain.Patient{}, err
	}
	
	return patient, nil
}

func (service *service) DeletePatient(id int) error  {
	err := service.repository.DeletePatient(id)
	if err != nil {
		return err
	}
	return nil
} 
