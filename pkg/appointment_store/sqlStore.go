package appointment_store

import (
	"database/sql"
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

func (s *sqlStore) GetAppointmentByDni(id int)([]domain.Appointment, error){
	var appointments []domain.Appointment
	query := "SELECT * FROM appointments WHERE patient = ?"

	rows, err := s.db.Query(query, id)
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
