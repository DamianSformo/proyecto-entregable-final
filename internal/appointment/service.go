package appointment

import (
	"github.com/DamianSformo/proyecto-entregable-final/internal/dentist"
	"github.com/DamianSformo/proyecto-entregable-final/internal/domain"
	"github.com/DamianSformo/proyecto-entregable-final/internal/patient"
)

type Service interface {
	GetAppointmentById(id int) (domain.Appointment, error)
	GetAppointmentByDni(dni int) ([]domain.Appointment, error)
	CreateAppointment(a domain.Appointment) (domain.Appointment, error)
	CreateAppointmentByDniAndLicense(a domain.Appointment, dniPatient int, licenseDentist string) (domain.Appointment, error)
	UpdateAppointment(a domain.Appointment, id int) (domain.Appointment, error)
	DeleteAppointment(id int) error
}

type service struct {
	repository Repository
	repositoryPatient patient.Repository
	repositoryDentist dentist.Repository
}

func NewService(repository Repository, repositoryPatient patient.Repository, repositoryDentist dentist.Repository) Service {
	return &service{repository, repositoryPatient, repositoryDentist}
}

func (service *service) CreateAppointment(a domain.Appointment) (domain.Appointment, error) {
	
	p, err := service.repositoryPatient.GetPatientByID(a.Patient.Id) 
	if err != nil {
		return domain.Appointment{}, err
	}

	d, err := service.repositoryDentist.GetDentistById(a.Dentist.Id) 
	if err != nil {
		return domain.Appointment{}, err
	}

	a.Patient = p
	a.Dentist = d

	a, err = service.repository.CreateAppointment(a)
	if err != nil {
		return domain.Appointment{}, err
	}

	return a, nil
}

func (service *service) CreateAppointmentByDniAndLicense(a domain.Appointment, dniPatient int, licenseDentist string) (domain.Appointment, error) {
	
	p, err := service.repositoryPatient.GetPatientByDni(dniPatient) 
	if err != nil {
		return domain.Appointment{}, err
	}

	d, err := service.repositoryDentist.GetDentistByLicense(licenseDentist) 
	if err != nil {
		return domain.Appointment{}, err
	}

	a.Patient = p
	a.Dentist = d

	a, err = service.repository.CreateAppointment(a)
	if err != nil {
		return domain.Appointment{}, err
	}

	return a, nil
}

func (service *service) UpdateAppointment(a domain.Appointment, id int) (domain.Appointment, error) {

	appointment, err := service.repository.GetAppointmentById(id)
	if err != nil {
		return domain.Appointment{}, err
	}

	p, err := service.repositoryPatient.GetPatientByID(a.Patient.Id) 
	if err != nil {
		return domain.Appointment{}, err
	}

	d, err := service.repositoryDentist.GetDentistById(a.Dentist.Id) 
	if err != nil {
		return domain.Appointment{}, err
	}

	a.Patient = p
	a.Dentist = d

	appointment, err = service.repository.UpdateAppointment(a, id)
	if err != nil {
		return domain.Appointment{}, err
	}
	
	return appointment, nil
}

func (service *service) GetAppointmentById(id int) (domain.Appointment, error) {

	appointment, err := service.repository.GetAppointmentById(id)
	if err != nil { 
		return domain.Appointment{}, err
	}

	p, err := service.repositoryPatient.GetPatientByID(appointment.Patient.Id) 
	if err != nil {
		return domain.Appointment{}, err
	}
	d, err := service.repositoryDentist.GetDentistById(appointment.Dentist.Id) 
	if err != nil {
		return domain.Appointment{}, err
	}

	appointment.Patient = p
	appointment.Dentist = d

	return appointment, nil
}

func (service *service) GetAppointmentByDni(dni int) ([]domain.Appointment, error) {

	p, err := service.repositoryPatient.GetPatientByDni(dni) 
	if err != nil {
		return nil, err
	}

	appointments, err := service.repository.GetAppointmentByDni(p.Id) 
	if err != nil {
		return nil, err
	}

	for i := range appointments {
		d, err := service.repositoryDentist.GetDentistById(appointments[i].Dentist.Id) 
		if err != nil {
			return nil, err
		}
		appointments[i].Dentist = d
		appointments[i].Patient = p
	}

	return appointments, nil
}

func (service *service) DeleteAppointment(id int) error {
	err := service.repository.DeleteAppointment(id)
	if err != nil {
		return err
	}
	return nil
}

