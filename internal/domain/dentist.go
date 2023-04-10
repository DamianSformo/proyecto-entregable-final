package domain

type Dentist struct {
	Id          int     `json:"id"`
	Name        string  `json:"name" binding:"required"`
	Surname		string	`json:"surname" binding:"required"`
	License     string	`json:"license" binding:"required"`
}