package handler

import (
	"errors"
	"fmt"
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

func NewPatientHandler(s patient.Service) *patientHandler {
	return &patientHandler{
		service : s,
	}
}



func (h *patientHandler) GetPatientByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("Invalid format id"))
			return
		}
		product, err := h.service.GetPatientByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New(err.Error()))
			return
		}
		web.Success(c, 200, product)
	}
}



func (h *patientHandler) GetPatientByDni() gin.HandlerFunc {
	return func(c *gin.Context) {
		dniParam := c.Param("dni")
		dni, err := strconv.Atoi(dniParam)
		if err != nil {
			web.Failure(c, 400, errors.New("Invalid format dni"))
			return
		}
		product, err := h.service.GetPatientByDni(dni)
		if err != nil {
			web.Failure(c, 404, errors.New(err.Error()))
			return
		}
		web.Success(c, 200, product)
	}
}


func (h *patientHandler) PostPatient() gin.HandlerFunc {

	return func(c *gin.Context) {

		var patient domain.Patient
		c.BindJSON(&patient)

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

		if patient.DNI < 1 {
			web.Failure(c, 400, errors.New("Invalid format dni"))
			return
		}

		fmt.Println(patient)

		p, err := h.service.CreatePatient(patient)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, p)
	}
}



func (h *patientHandler) PutPatient() gin.HandlerFunc {
	return func(c *gin.Context) {
		//token := c.GetHeader("TOKEN")
		//if token == "" {
		//	web.Failure(c, 401, errors.New("token not found"))
		//		return
		//	}
		//if token != os.Getenv("TOKEN") {
		//	web.Failure(c, 401, errors.New("invalid token"))
		//	return
		//}
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}

		_, err = h.service.GetPatientByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New(err.Error()))
			return
		}

		var patient domain.Patient
		c.BindJSON(&patient)

		//err = c.ShouldBindJSON(&patient)
		//if err != nil {
		//	web.Failure(c, 400, errors.New("invalid json"))
		//	return
		//}

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

		p, err := h.service.UpdatePatient(patient, id)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p)
	}
}

func (h *patientHandler) DeletePatient() gin.HandlerFunc {
	return func(c *gin.Context) {

		//token := c.GetHeader("TOKEN")
		//if token == "" {
		//	web.Failure(c, 401, errors.New("token not found"))
		//	return
		//}
		//if token != os.Getenv("TOKEN") {
		//	web.Failure(c, 401, errors.New("invalid token"))
		//	return
		//}
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		err = h.service.DeletePatient(id)
		if err != nil {
			web.Failure(c, 404, err)
			return
		}
		web.Success(c, 204, nil)
	}
}

// validateEmptys valida que los campos no esten vacios
func validatePatientEmpty(patient *domain.Patient) (bool, error) {
	switch {
	case patient.Name == "" || patient.Surname == "" || patient.DNI < 0 || patient.Address == "" :
		return false, errors.New("fields can't be empty")
	}
	return true, nil
}

// validatePatientDate valida que el campo feche no exista
func validatePatientDate(patient *domain.Patient) (bool, error) {
	if patient.Date != ""{
		return false, errors.New("la fecha no se puede editar")
	}
	return true, nil
}