package handler

import (
	"errors"
	"os"
	"strconv"
	"github.com/DamianSformo/proyecto-entregable-final/internal/appointment"
	"github.com/DamianSformo/proyecto-entregable-final/internal/domain"
	"github.com/DamianSformo/proyecto-entregable-final/pkg/web"
	"github.com/gin-gonic/gin"
)

type appointmentHandler struct {
	service appointment.Service
}

func NewAppointmentHandler(service appointment.Service) *appointmentHandler {
	return &appointmentHandler{
		service : service,
	}
}


func (handler *appointmentHandler) GetAppointmentById() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("Invalid format id"))
			return
		}
		appointment, err := handler.service.GetAppointmentById(id)
		if err != nil {
			web.Failure(c, 404, errors.New(err.Error()))
			return
		}
		web.Success(c, 200, appointment)
	}
}


func (handler *appointmentHandler) GetAppointmentByDni() gin.HandlerFunc {
	return func(c *gin.Context) {
		dniParam := c.Param("dni")
		dni, err := strconv.Atoi(dniParam)
		if err != nil {
			web.Failure(c, 400, errors.New("Invalid format dni"))
			return
		}
		appointment, err := handler.service.GetAppointmentByDni(dni)
		if err != nil {
			web.Failure(c, 404, errors.New(err.Error()))
			return
		}
		web.Success(c, 200, appointment)
	}
}


func (handler *appointmentHandler) PostAppointment() gin.HandlerFunc {

	return func(c *gin.Context) {

		var a domain.Appointment
		c.BindJSON(&a)

		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("Token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("Invalid token"))
			return
		}

		valid, err := validateAppointmentEmpty(&a)
		if !valid {
			web.Failure(c, 400, err)
			return
		}

		appointment, err := handler.service.CreateAppointment(a)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, appointment)
	}
}


func (handler *appointmentHandler) PostAppointmentByDniAndLicense() gin.HandlerFunc {

	return func(c *gin.Context) {

		dniPatientParam := c.Param("dni")
		dniPatient, err := strconv.Atoi(dniPatientParam)
		if err != nil {
			web.Failure(c, 400, errors.New("Invalid format dni"))
			return
		}

		licenseDentist := c.Param("license")

		var a domain.Appointment
		c.BindJSON(&a)

		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("Token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("Invalid token"))
			return
		}

		appointment, err := handler.service.CreateAppointmentByDniAndLicense(a, dniPatient, licenseDentist)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, appointment)
	}
}


func (handler *appointmentHandler) DeleteAppointment() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("Token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("Invalid token"))
			return
		}
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("Invalid id"))
			return
		}
		err = handler.service.DeleteAppointment(id)
		if err != nil {
			web.Failure(c, 404, err)
			return
		}
		web.Success(c, 204, nil)
	}
}


func validateAppointmentEmpty(appointment *domain.Appointment) (bool, error) {
	switch {
	case appointment.Date == "" || appointment.Dentist.Id == 0 || appointment.Patient.Id == 0 :
		return false, errors.New("Fields can't be empty")
	}
	return true, nil
}