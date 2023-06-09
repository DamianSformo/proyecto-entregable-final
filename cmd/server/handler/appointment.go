package handler

import (
	"errors"
	"fmt"
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


func (handler *appointmentHandler) PutAppointment() gin.HandlerFunc {
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

		_, err = handler.service.GetAppointmentById(id)
		if err != nil {
			web.Failure(c, 404, errors.New(err.Error()))
			return
		}

		var appointment domain.Appointment
		c.BindJSON(&appointment)

		valid, err := validateAppointmentEmpty(&appointment)
		if !valid {
			web.Failure(c, 400, err)
			return
		}

		a, err := handler.service.UpdateAppointment(appointment, id)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, a)
	}
}


func (handler *appointmentHandler) PatchAppointment() gin.HandlerFunc {
	type RequestPatient struct {
		Id			int 	`json:"id"`
		Name        string  `json:"name,omitempty"`
		Surname		string	`json:"surname,omitempty"`
		DNI 		int		`json:"dni,omitempty"`
		Address     string	`json:"address,omitempty"`
	}

	type RequestDentist struct {
		Id			int 	`json:"id"`
		Name        string  `json:"name,omitempty"`
		Surname    	string  `json:"surname,omitempty"`
		License   	string  `json:"license,omitempty"`
	}

	type Request struct {
		Patient     	RequestPatient 	`json:"patient,omitempty"`
		Dentist			RequestDentist	`json:"dentist,omitempty"`
		Date 			string			`json:"date,omitempty"`
		Description		string			`json:"description,omitempty"`
	}
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
		
		var r Request
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("Invalid id"))
			return
		}

		err = c.ShouldBindJSON(&r)
		fmt.Println(err)
		if err != nil {
			web.Failure(c, 400, errors.New("Invalid request body"))
			return
		}

		appointment, err := handler.service.GetAppointmentById(id)
		if err != nil {
			web.Failure(c, 404, errors.New(err.Error()))
			return
		}
		
		if r.Dentist.Id > 0{
			appointment.Dentist.Id = r.Dentist.Id
		}

		if r.Patient.Id > 0{
			appointment.Patient.Id = r.Patient.Id
		}

		if r.Date != ""{
			appointment.Date = r.Date
		}

		if r.Description != ""{
			appointment.Description = r.Description
		}

		a, err := handler.service.UpdateAppointment(appointment, id)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, a)
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