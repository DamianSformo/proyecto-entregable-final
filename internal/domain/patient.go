package domain

type Patient struct {
	Id          int     `json:"id"`
	Name        string  `json:"name" binding:"required"`
	Surname		string	`json:"surname" binding:"required"`
	DNI 		int		`json:"dni" binding:"required"`
	Address     string	`json:"address" binding:"required"`
	Date		string	`json:"date"`
}