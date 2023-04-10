package patient_store

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
 
func (s *sqlStore) GetPatientById(id int)(domain.Patient, error){
	var patient domain.Patient
	query := "SELECT * FROM patients WHERE id = ?"
	row := s.db.QueryRow(query, id)
	err := row.Scan(&patient.Id, &patient.Name, &patient.Surname, &patient.DNI, &patient.Address, &patient.Date)
	if err != nil{
		return domain.Patient{}, err
	}

	return patient, nil
}

func (s *sqlStore) GetPatientByDni(dni int)(domain.Patient, error){
	var patient domain.Patient
	query := "SELECT * FROM patients WHERE dni = ?"
	row := s.db.QueryRow(query, dni)
	err := row.Scan(&patient.Id, &patient.Name, &patient.Surname, &patient.DNI, &patient.Address, &patient.Date)
	if err != nil{
		return domain.Patient{}, err
	}

	return patient, nil
}

func (s *sqlStore)CreatePatient(patient domain.Patient) (int64, error){
	query := "insert into patients (name, surname, dni, address, date) VALUES (?, ?, ?, ?, ?)"
	stm, err := s.db.Prepare(query)
	if err != nil{
		return 0, err
	}

	res, err := stm.Exec(patient.Name, patient.Surname, patient.DNI, patient.Address,patient.Date)
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

func (s * sqlStore)UpdatePatient(product domain.Patient, id int) error{
	query := "UPDATE patients SET name=?, surname=?, dni=?, address=?, date=? WHERE id=?"
	stm, err := s.db.Prepare(query)
	if err != nil{
		return err
	}

	res, err := stm.Exec(product.Name, product.Surname, product.DNI, product.Address, product.Date, id)
	if err != nil{
		return err
	}

	if _, err := res.RowsAffected(); err != nil{
		return err
	}

	return nil
}

func (s * sqlStore)ExistsPatient(id int) bool{
	row := s.db.QueryRow("SELECT id FROM patients WHERE id = ?", id)
	var i int
	if err := row.Scan(&i); err != nil{
		return false
	}
	if i > 0 {
		return true
	}

	return false
}


// Delete elimina un producto
func (s * sqlStore)DeletePatient(id int) error{
	query := "DELETE FROM patients WHERE id = ?"

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
