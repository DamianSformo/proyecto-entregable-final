package dentist_store

import "github.com/DamianSformo/proyecto-entregable-final/internal/domain"

type StoreInterface interface {
	GetDentistById(id int) (domain.Dentist, error)
	GetDentistByLicense(license string) (domain.Dentist, error)
	CreateDentist(product domain.Dentist) (int64, error)
	Update(product domain.Dentist, id int) error
	DeleteDentist(id int) error
	ExistsDentist(id int) bool
}
