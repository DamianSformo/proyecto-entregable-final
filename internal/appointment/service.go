package appointment

import (
	"github.com/DamianSformo/proyecto-entregable-final/internal/dentist"
	"github.com/DamianSformo/proyecto-entregable-final/internal/domain"
	"github.com/DamianSformo/proyecto-entregable-final/internal/patient"
)

type Service interface {
	GetAppointmentByDni(dni int) ([]domain.Appointment, error)
	CreateAppointment(a domain.Appointment) (domain.Appointment, error)
	CreateAppointmentByDniAndLicense(a domain.Appointment, dniPatient int, licenseDentist string) (domain.Appointment, error)
}

type service struct {
	repository Repository
	repositoryPatient patient.Repository
	repositoryDentist dentist.Repository
}

func NewService(r Repository, rp patient.Repository, rd dentist.Repository) Service {
	return &service{r, rp, rd}
}

func (s *service) CreateAppointment(a domain.Appointment) (domain.Appointment, error) {
	
	p, err := s.repositoryPatient.GetPatientByID(a.Patient.Id) 
	if err != nil {
		return domain.Appointment{}, err
	}

	d, err := s.repositoryDentist.GetDentistById(a.Dentist.Id) 
	if err != nil {
		return domain.Appointment{}, err
	}

	a.Patient = p
	a.Dentist = d

	a, err = s.repository.CreateAppointment(a)
	if err != nil {
		return domain.Appointment{}, err
	}

	return a, nil
}

func (s *service) CreateAppointmentByDniAndLicense(a domain.Appointment, dniPatient int, licenseDentist string) (domain.Appointment, error) {
	
	p, err := s.repositoryPatient.GetPatientByDni(dniPatient) 
	if err != nil {
		return domain.Appointment{}, err
	}

	d, err := s.repositoryDentist.GetDentistByLicense(licenseDentist) 
	if err != nil {
		return domain.Appointment{}, err
	}

	a.Patient = p
	a.Dentist = d

	a, err = s.repository.CreateAppointment(a)
	if err != nil {
		return domain.Appointment{}, err
	}

	return a, nil
}

func (s *service) GetAppointmentByDni(dni int) ([]domain.Appointment, error) {

	p, err := s.repositoryPatient.GetPatientByDni(dni) 
	if err != nil {
		return nil, err
	}

	appointments, err := s.repository.GetAppointmentByDni(p.Id) 
	if err != nil {
		return nil, err
	}

	for i := range appointments {
		d, err := s.repositoryDentist.GetDentistById(appointments[i].Dentist.Id) 
		if err != nil {
			return nil, err
		}
		appointments[i].Dentist = d
		appointments[i].Patient = p
	}

	return appointments, nil
}
