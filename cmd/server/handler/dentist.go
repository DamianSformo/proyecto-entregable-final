package handler

import (
	"errors"
	"os"
	"strconv"

	"github.com/DamianSformo/proyecto-entregable-final/internal/dentist"
	"github.com/DamianSformo/proyecto-entregable-final/internal/domain"
	"github.com/DamianSformo/proyecto-entregable-final/pkg/web"
	"github.com/gin-gonic/gin"
)

type dentistHandler struct {
	service dentist.Service
}

func NewDentistHandler(service dentist.Service) *dentistHandler {
	return &dentistHandler{
		service: service,
	}
}

func (handler *dentistHandler) GetDentistById() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("Invalid format id"))
			return
		}
		dentist, err := handler.service.GetDentistById(id)
		if err != nil {
			web.Failure(c, 404, errors.New(err.Error()))
			return
		}
		web.Success(c, 200, dentist)
	}
}

func (handler *dentistHandler) PostDentist() gin.HandlerFunc {
	return func(c *gin.Context) {

		var dentist domain.Dentist
		c.BindJSON(&dentist)

		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("Token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("Invalid token"))
			return
		}

		valid, err := validateDentistEmpty(&dentist)
		if !valid {
			web.Failure(c, 400, err)
			return
		}

		d, err := handler.service.CreateDentist(dentist)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, d)
	}
}

func (handler *dentistHandler) PutDentist() gin.HandlerFunc {
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

		_, err = handler.service.GetDentistById(id)
		if err != nil {
			web.Failure(c, 404, errors.New(err.Error()))
			return
		}

		if err != nil {
			web.Failure(c, 409, err)
			return
		}

		var dentist domain.Dentist
		err = c.ShouldBindJSON(&dentist)

		valid, err := validateDentistEmpty(&dentist)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		
		d, err := handler.service.UpdateDentist(id, dentist)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, d)
	}
}

func (h *dentistHandler) PatchDentist() gin.HandlerFunc {
	type Request struct {
		Name        string  `json:"name,omitempty"`
		Surname    	string  `json:"surname,omitempty"`
		License   	string  `json:"license,omitempty"`
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

		dentist, err := h.service.GetDentistById(id)
		if err != nil {
			web.Failure(c, 404, errors.New(err.Error()))
			return
		}
		
		if r.Name != ""{
			dentist.Name = r.Name
		}

		if r.Surname != ""{
			dentist.Surname = r.Surname
		}

		if r.License != ""{
			dentist.License = r.License
		}

		d, err := h.service.UpdateDentist(id, dentist)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, d)
	}
}

func (handler *dentistHandler) DeleteDentist() gin.HandlerFunc {
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
		
		err = handler.service.DeleteDentist(id)
		if err != nil {
			web.Failure(c, 404, err)
			return
		}
		web.Success(c, 204, nil)
	}
}

func validateDentistEmpty(dentist *domain.Dentist) (bool, error) {
	switch {
	case dentist.Name == "" || dentist.Surname == "" || dentist.License == "" :
		return false, errors.New("Fields can't be empty")
	}
	return true, nil
}