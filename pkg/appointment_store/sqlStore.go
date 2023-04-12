package appointment_store

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

func (s *sqlStore) GetAppointmentById(id int)(domain.Appointment, error){
	var appointment domain.Appointment
	query := "SELECT * FROM appointments WHERE id = ?"
	row := s.db.QueryRow(query, id)
	err := row.Scan(&appointment.Id, &appointment.Date, &appointment.Description, &appointment.Patient.Id, &appointment.Dentist.Id)
	if err != nil{
		return domain.Appointment{}, err
	}

	return appointment, nil
}

func (s *sqlStore) GetAppointmentByDni(dni int)([]domain.Appointment, error){
	var appointments []domain.Appointment
	query := "SELECT * FROM appointments WHERE patient = ?"

	rows, err := s.db.Query(query, dni)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var appointment domain.Appointment
		err := rows.Scan(&appointment.Id, &appointment.Date, &appointment.Description, &appointment.Patient.Id, &appointment.Dentist.Id)
		if err != nil {
			return nil, err
		}
		appointments = append(appointments, appointment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return appointments, nil
}

func (s *sqlStore)CreateAppointment(appointment domain.Appointment) (int64, error){
	
	query := "insert into appointments (patient, dentist, date, description) VALUES (?, ?, ?, ?)"
	stm, err := s.db.Prepare(query)
	if err != nil{
		return 0, err
	} 

	res, err := stm.Exec(appointment.Patient.Id, appointment.Dentist.Id, appointment.Date, appointment.Description)
	
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

func (s * sqlStore)UpdateAppointment(appointment domain.Appointment, id int) error{
	query := "UPDATE appointments SET patient=?, dentist=?, date=?, description=? WHERE id=?"
	stm, err := s.db.Prepare(query)
	if err != nil{
		return err
	}

	res, err := stm.Exec(appointment.Patient.Id, appointment.Dentist.Id, appointment.Date, appointment.Description, id)
	if err != nil{
		return err
	}

	if _, err := res.RowsAffected(); err != nil{
		return err
	}

	return nil
}


func (s * sqlStore)DeleteAppointment(id int) error{
	query := "DELETE FROM appointments WHERE id = ?"

	stm, err := s.db.Prepare(query)
	if err != nil{
		return err
	}

	res, err := stm.Exec(id)
	if err != nil{
		return err
	}
	fmt.Println(res)
	
	return nil
}


func (s * sqlStore)ExistsAppointment(id int) bool{
	row := s.db.QueryRow("SELECT id FROM appointments WHERE id = ?", id)
	var i int
	if err := row.Scan(&i); err != nil{
		return false
	}
	if i > 0 {
		return true
	}

	return false
}

