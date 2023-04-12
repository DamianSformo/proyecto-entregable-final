package handler

import (
	"errors"
	"os"
	"strconv"
	"github.com/DamianSformo/proyecto-entregable-final/internal/domain"
	"github.com/DamianSformo/proyecto-entregable-final/internal/patient"
	"github.com/DamianSformo/proyecto-entregable-final/pkg/web"
	"github.com/gin-gonic/gin"
)

type patientHandler struct {
	service patient.Service
}

func NewPatientHandler(service patient.Service) *patientHandler {
	return &patientHandler{
		service : service,
	}
}

func (handler *patientHandler) GetPatientByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("Invalid format id"))
			return
		}
		patient, err := handler.service.GetPatientByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New(err.Error()))
			return
		}
		web.Success(c, 200, patient)
	}
}


func (handler *patientHandler) GetPatientByDni() gin.HandlerFunc {
	return func(c *gin.Context) {
		dniParam := c.Param("dni")
		dni, err := strconv.Atoi(dniParam)
		if err != nil {
			web.Failure(c, 400, errors.New("Invalid format dni"))
			return
		}
		patient, err := handler.service.GetPatientByDni(dni)
		if err != nil {
			web.Failure(c, 404, errors.New(err.Error()))
			return
		}
		web.Success(c, 200, patient)
	}
}


func (handler *patientHandler) PostPatient() gin.HandlerFunc {

	return func(c *gin.Context) {

		var patient domain.Patient
		c.BindJSON(&patient)

		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("Token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("Invalid token"))
			return
		}

		valid, err := validatePatientEmpty(&patient)
		if !valid {
			web.Failure(c, 400, err)
			return
		}

		validDate, err := validatePatientDate(&patient)
		if !validDate {
			web.Failure(c, 400, err)
			return
		}

		if patient.DNI == 0 {
			web.Failure(c, 400, errors.New("Invalid format dni"))
			return
		}

		p, err := handler.service.CreatePatient(patient)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, p)
	}
}


func (handler *patientHandler) PutPatient() gin.HandlerFunc {
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

		_, err = handler.service.GetPatientByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New(err.Error()))
			return
		}

		var patient domain.Patient
		c.BindJSON(&patient)

		valid, err := validatePatientEmpty(&patient)
		if !valid {
			web.Failure(c, 400, err)
			return
		}

		validDate, err := validatePatientDate(&patient)
		if !validDate {
			web.Failure(c, 400, err)
			return
		}

		p, err := handler.service.UpdatePatient(patient, id)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p)
	}
}


func (handler *patientHandler) PatchPatient() gin.HandlerFunc {
	type Request struct {
		Name        string  `json:"name,omitempty"`
		Surname		string	`json:"surname,omitempty"`
		DNI 		int		`json:"dni,omitempty"`
		Address     string	`json:"address,omitempty"`
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
		if err != nil {
			web.Failure(c, 400, errors.New("Invalid request body"))
			return
		}

		patient, err := handler.service.GetPatientByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New(err.Error()))
			return
		}
		
		if r.Name != ""{
			patient.Name = r.Name
		}

		if r.Surname != ""{
			patient.Surname = r.Surname
		}

		if r.DNI > 0{
			patient.DNI = r.DNI
		}

		if r.Address != ""{
			patient.Address = r.Address
		}

		p, err := handler.service.UpdatePatient(patient, id)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p)
	}
}


func (handler *patientHandler) DeletePatient() gin.HandlerFunc {
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
		err = handler.service.DeletePatient(id)
		if err != nil {
			web.Failure(c, 404, err)
			return
		}
		web.Success(c, 204, nil)
	}
}


func validatePatientEmpty(patient *domain.Patient) (bool, error) {
	switch {
	case patient.Name == "" || patient.Surname == "" || patient.Address == "" :
		return false, errors.New("Fields can't be empty")
	}
	return true, nil
}


func validatePatientDate(patient *domain.Patient) (bool, error) {
	if patient.Date != ""{
		return false, errors.New("Date cannot be edited")
	}
	return true, nil
}