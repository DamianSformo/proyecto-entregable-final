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

func NewDentistHandler(s dentist.Service) *dentistHandler {
	return &dentistHandler{
		service: s,
	}
}

func (h *dentistHandler) GetDentistById() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid format id"))
			return
		}
		dentist, err := h.service.GetDentistById(id)
		if err != nil {
			web.Failure(c, 404, errors.New(err.Error()))
			return
		}
		web.Success(c, 200, dentist)
	}
}

func (h *dentistHandler) PostDentist() gin.HandlerFunc {
	return func(c *gin.Context) {

		var dentist domain.Dentist
		c.BindJSON(&dentist)

		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("invalid token"))
			return
		}

		valid, err := validateEmptys(&dentist)
		if !valid {
			web.Failure(c, 400, err)
			return
		}

		p, err := h.service.CreateDentist(dentist)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, p)
	}
}

func (h *dentistHandler) DeleteDentist() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("invalid token"))
			return
		}
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		
		err = h.service.DeleteDentist(id)
		if err != nil {
			web.Failure(c, 404, err)
			return
		}
		web.Success(c, 204, nil)
	}
}

func (h *dentistHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("token not found"))
				return
			}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("invalid token"))
			return
		}
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}

		_, err = h.service.GetDentistById(id)
		if err != nil {
			web.Failure(c, 404, errors.New(err.Error()))
			return
		}

		if err != nil {
			web.Failure(c, 409, err)
			return
		}

		var product domain.Dentist
		err = c.ShouldBindJSON(&product)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}

		valid, err := validateEmptys(&product)
		if !valid {
			web.Failure(c, 400, err)
			return
		}

		p, err := h.service.UpdateDentist(id, product)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p)
	}
}

func (h *dentistHandler) Patch() gin.HandlerFunc {
	type Request struct {
		Name        string  `json:"name,omitempty"`
		Surname    	string  `json:"surname,omitempty"`
		License   	string  `json:"license,omitempty"`
	}
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("invalid token"))
			return
		}
		
		var r Request
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}

		err = c.ShouldBindJSON(&r)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid request body"))
			return
		}

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

func validateDentistEmpty(dentist *domain.Dentist) (bool, error) {
	switch {
	case dentist.Name == "" || dentist.Surname == "" || dentist.License == "" :
		return false, errors.New("fields can't be empty")
	}
	return true, nil
}

func validateEmptys(product *domain.Dentist) (bool, error) {
	switch {
	case product.Name == "" || product.Surname == "" || product.License == "":
		return false, errors.New("fields can't be empty")
	}
	return true, nil
}