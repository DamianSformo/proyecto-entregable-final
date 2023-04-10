package dentist_store

import (
	"database/sql"
	"fmt"
	"github.com/DamianSformo/proyecto-entregable-final/internal/domain"
)

type sqlStore struct{
	db *sql.DB
}

func NewSqlStore(db * sql.DB) StoreInterface{
	return &sqlStore{
		db: db,
	}
}

func (s *sqlStore) GetDentistById(id int)(domain.Dentist, error){
	var dentist domain.Dentist
	query := "SELECT * FROM dentists WHERE id = ?;"
	row := s.db.QueryRow(query, id)

	err := row.Scan(&dentist.Id, &dentist.Name, &dentist.Surname, &dentist.License)
	if err != nil{
		return domain.Dentist{}, err
	}

	return dentist, nil
}

func (s *sqlStore) GetDentistByLicense(license string)(domain.Dentist, error){
	var dentist domain.Dentist
	query := "SELECT * FROM dentists WHERE license = ?"
	row := s.db.QueryRow(query, license)
	err := row.Scan(&dentist.Id, &dentist.Name, &dentist.Surname, &dentist.License)
	if err != nil{
		return domain.Dentist{}, err
	}

	return dentist, nil
}

func (s * sqlStore)CreateDentist(product domain.Dentist) (int64, error){
	query := "insert into dentists (name, surname, license) VALUES (?, ?, ?)"
	stm, err := s.db.Prepare(query)
	if err != nil{
		return 0, err
	}

	res, err := stm.Exec(product.Name, product.Surname, product.License)

	if err != nil{
		return 0, err
	}
 
	if _, err := res.RowsAffected(); err != nil{
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s * sqlStore)Update(product domain.Dentist, id int) error{
	query := "UPDATE dentists SET name=?, surname=?, license=? WHERE id=?"
	stm, err := s.db.Prepare(query)
	if err != nil{
		return err
	}

	res, err := stm.Exec(product.Name, product.Surname, product.License, id)
	if err != nil{
		return err
	}

	if _, err := res.RowsAffected(); err != nil{
		return err
	}

	return nil
}

func (s * sqlStore)DeleteDentist(id int) error{
	query := "DELETE FROM dentists WHERE id = ?"

	stm, err := s.db.Prepare(query)
	if err != nil{
		return err
	}

	res, err := stm.Exec(id)
	if err != nil{
		return err
	}
	fmt.Print(res)
	
	return nil
}

func (s * sqlStore)ExistsDentist(id int) bool{
	row := s.db.QueryRow("SELECT id FROM dentists WHERE id = ?", id)
	var i int
	if err := row.Scan(&i); err != nil{
		return false
	}
	if i > 0 {
		return true
	}

	return false
}
