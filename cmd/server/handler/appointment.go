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

func NewAppointmentHandler(s appointment.Service) *appointmentHandler {
	return &appointmentHandler{
		service : s,
	}
}

func (h *appointmentHandler) GetAppointmentByDni() gin.HandlerFunc {
	return func(c *gin.Context) {
		dniParam := c.Param("dni")
		dni, err := strconv.Atoi(dniParam)
		if err != nil {
			web.Failure(c, 400, errors.New("Invalid format dni"))
			return
		}
		product, err := h.service.GetAppointmentByDni(dni)
		if err != nil {
			web.Failure(c, 404, errors.New(err.Error()))
			return
		}
		web.Success(c, 200, product)
	}
}

func (h *appointmentHandler) PostAppointment() gin.HandlerFunc {

	return func(c *gin.Context) {

		//dniPatientParam := c.Param("dniPatient")
		//dniPatient, err := strconv.Atoi(dniPatientParam)
		//if err != nil {
		//	web.Failure(c, 400, errors.New("Invalid format dni"))
		//	return
		//}

		var appointment domain.Appointment
		c.BindJSON(&appointment)

		token := c.GetHeader("TOKEN")
		//if token == "" {
		//	web.Failure(c, 401, errors.New("token not found"))
		//	return
		//}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("invalid token"))
			return
		}
		
		//if err != nil {
		//	web.Failure(c, 400, errors.New("invalid json"))
		//	return
		//}

		//valid, err := validatePatientEmpty(&appointment)
		//if !valid {
		//	web.Failure(c, 400, err)
		//	return
		//}
//
		//validDate, err := validatePatientDate(&appointment)
		//if !validDate {
		//	web.Failure(c, 400, err)
		//	return
		//}

		p, err := h.service.CreateAppointment(appointment)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, p)
	}
}

func (h *appointmentHandler) PostAppointmentByDniAndLicense() gin.HandlerFunc {

	return func(c *gin.Context) {

		dniPatientParam := c.Param("dni")
		dniPatient, err := strconv.Atoi(dniPatientParam)
		if err != nil {
			web.Failure(c, 400, errors.New("Invalid format dni"))
			return
		}

		licenseDentist := c.Param("license")

		//licenseDentist := c.Param("license")

		var appointment domain.Appointment
		c.BindJSON(&appointment)

		token := c.GetHeader("TOKEN")
		//if token == "" {
		//	web.Failure(c, 401, errors.New("token not found"))
		//	return
		//}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("invalid token"))
			return
		}
		
		//if err != nil {
		//	web.Failure(c, 400, errors.New("invalid json"))
		//	return
		//}

		//valid, err := validatePatientEmpty(&appointment)
		//if !valid {
		//	web.Failure(c, 400, err)
		//	return
		//}
//
		//validDate, err := validatePatientDate(&appointment)
		//if !validDate {
		//	web.Failure(c, 400, err)
		//	return
		//}

		p, err := h.service.CreateAppointmentByDniAndLicense(appointment, dniPatient, licenseDentist)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, p)
	}
}