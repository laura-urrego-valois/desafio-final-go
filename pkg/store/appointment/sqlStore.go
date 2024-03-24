package store

import (
	"database/sql"
	"errors"
	"proyecto_final_go/internal/domain"
)

type sqlAppointmentStore struct {
	db *sql.DB
}

func NewSqlAppointmentStore(db *sql.DB) AppointmentStoreInterface {
	return &sqlAppointmentStore{
		db: db,
	}
}

//-----------------------------------

func (s *sqlAppointmentStore) Read(id int) (domain.Appointment, error) {
	var appointment domain.Appointment
	query := `
		SELECT 
			a.id, a.date, a.hour, a.description,
			p.id AS patient_id, p.first_name AS patient_first_name, p.last_name AS patient_last_name, p.address AS patient_address, p.dni AS patient_dni, p.release_date AS patient_release_date,
			d.id AS dentist_id, d.first_name AS dentist_first_name, d.last_name AS dentist_last_name, d.license AS dentist_license
		FROM 
			appointments AS a
		INNER JOIN 
			patients AS p ON a.patient_id = p.id
		INNER JOIN 
			dentists AS d ON a.dentist_id = d.id
		WHERE 
			a.id = ?;
	`
	row := s.db.QueryRow(query, id)
	err := row.Scan(
		&appointment.Id, &appointment.Date, &appointment.Hour, &appointment.Description,
		&appointment.Patient.Id, &appointment.Patient.FirstName, &appointment.Patient.LastName, &appointment.Patient.Address, &appointment.Patient.DNI, &appointment.Patient.ReleaseDate,
		&appointment.Dentist.Id, &appointment.Dentist.FirstName, &appointment.Dentist.LastName, &appointment.Dentist.License,
	)
	if err != nil {
		return domain.Appointment{}, err
	}
	return appointment, nil
}

func (s *sqlAppointmentStore) ReadByPatientDNI(patientDNI string) ([]domain.Appointment, error) {
	var appointments []domain.Appointment
	query := `
		SELECT 
			a.id, a.date, a.hour, a.description,
			p.id AS patient_id, p.first_name AS patient_first_name, p.last_name AS patient_last_name, p.address AS patient_address, p.dni AS patient_dni, p.release_date AS patient_release_date,
			d.id AS dentist_id, d.first_name AS dentist_first_name, d.last_name AS dentist_last_name, d.license AS dentist_license
		FROM 
			appointments AS a
		INNER JOIN 
			patients AS p ON a.patient_id = p.id
		INNER JOIN 
			dentists AS d ON a.dentist_id = d.id
		WHERE 
			p.dni = ?;
	`
	rows, err := s.db.Query(query, patientDNI)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var appointment domain.Appointment
		err := rows.Scan(
			&appointment.Id, &appointment.Date, &appointment.Hour, &appointment.Description,
			&appointment.Patient.Id, &appointment.Patient.FirstName, &appointment.Patient.LastName, &appointment.Patient.Address, &appointment.Patient.DNI, &appointment.Patient.ReleaseDate,
			&appointment.Dentist.Id, &appointment.Dentist.FirstName, &appointment.Dentist.LastName, &appointment.Dentist.License,
		)
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

func (s *sqlAppointmentStore) Create(appointment domain.Appointment) error {
	query := `
		INSERT INTO appointments (date, hour, description, patient_id, dentist_id) 
		VALUES (?, ?, ?, ?, ?);
	`
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(appointment.Date, appointment.Hour, appointment.Description, appointment.Patient.Id, appointment.Dentist.Id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlAppointmentStore) CreateByPatientDNIAndDentistLicense(patientDNI string, license string, date string, hour string, description string) ([]domain.Appointment, error) {
	var appointments []domain.Appointment

	var patientID int
	patientQuery := "SELECT id FROM patients WHERE dni = ?"
	err := s.db.QueryRow(patientQuery, patientDNI).Scan(&patientID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Patient not found")
		}
		return nil, err
	}

	var dentistID int
	dentistQuery := "SELECT id FROM dentists WHERE license = ?"
	err = s.db.QueryRow(dentistQuery, license).Scan(&dentistID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Dentist not found")
		}
		return nil, err
	}

	createQuery := "INSERT INTO appointments (patient_id, dentist_id, date, hour, description) VALUES (?, ?, ?, ?, ?)"
	_, err = s.db.Exec(createQuery, patientID, dentistID, date, hour, description)
	if err != nil {
		return nil, err
	}

	appointmentsQuery := `
		SELECT 
			a.id, a.date, a.hour, a.description,
			p.id AS patient_id, p.first_name AS patient_first_name, p.last_name AS patient_last_name, p.address AS patient_address, p.dni AS patient_dni, p.release_date AS patient_release_date,
			d.id AS dentist_id, d.first_name AS dentist_first_name, d.last_name AS dentist_last_name, d.license AS dentist_license
		FROM 
			appointments AS a
		INNER JOIN 
			patients AS p ON a.patient_id = p.id
		INNER JOIN 
			dentists AS d ON a.dentist_id = d.id
		WHERE 
			a.patient_id = ? AND a.dentist_id = ?
	`
	rows, err := s.db.Query(appointmentsQuery, patientID, dentistID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var appointment domain.Appointment
		err := rows.Scan(
			&appointment.Id, &appointment.Date, &appointment.Hour, &appointment.Description,
			&appointment.Patient.Id, &appointment.Patient.FirstName, &appointment.Patient.LastName, &appointment.Patient.Address, &appointment.Patient.DNI, &appointment.Patient.ReleaseDate,
			&appointment.Dentist.Id, &appointment.Dentist.FirstName, &appointment.Dentist.LastName, &appointment.Dentist.License,
		)
		if err != nil {
			return nil, err
		}
		appointments = append(appointments, appointment)
	}

	return appointments, nil
}

func (s *sqlAppointmentStore) Update(appointment domain.Appointment) error {
	query := `
		UPDATE appointments 
		SET date = ?, hour = ?, description = ?, patient_id = ?, dentist_id = ?
		WHERE id = ?;
	`
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(appointment.Date, appointment.Hour, appointment.Description, appointment.Patient.Id, appointment.Dentist.Id, appointment.Id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("Appointment not found")
	}
	return nil
}

func (s *sqlAppointmentStore) Delete(id int) error {
	query := `
		DELETE FROM appointments 
		WHERE id = ?;
	`
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("Appointment not found")
	}
	return nil
}

func (s *sqlAppointmentStore) GetAll() ([]domain.Appointment, error) {
	var appointments []domain.Appointment

	query := `
		SELECT 
			a.id, a.date, a.hour, a.description,
			p.id AS patient_id, p.first_name AS patient_first_name, p.last_name AS patient_last_name, p.address AS patient_address, p.dni AS patient_dni, p.release_date AS patient_release_date,
			d.id AS dentist_id, d.first_name AS dentist_first_name, d.last_name AS dentist_last_name, d.license AS dentist_license
		FROM 
			appointments AS a
		INNER JOIN 
			patients AS p ON a.patient_id = p.id
		INNER JOIN 
			dentists AS d ON a.dentist_id = d.id;
	`

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var appointment domain.Appointment
		if err := rows.Scan(
			&appointment.Id, &appointment.Date, &appointment.Hour, &appointment.Description,
			&appointment.Patient.Id, &appointment.Patient.FirstName, &appointment.Patient.LastName, &appointment.Patient.Address, &appointment.Patient.DNI, &appointment.Patient.ReleaseDate,
			&appointment.Dentist.Id, &appointment.Dentist.FirstName, &appointment.Dentist.LastName, &appointment.Dentist.License,
		); err != nil {
			return nil, err
		}
		appointments = append(appointments, appointment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return appointments, nil
}

func (s *sqlAppointmentStore) Exists(id int) bool {
	var exists bool
	var appointmentId int
	query := "SELECT id FROM appointments WHERE id = ?;"
	row := s.db.QueryRow(query, id)
	err := row.Scan(&appointmentId)
	if err != nil {
		return false
	}
	if appointmentId > 0 {
		exists = true
	}
	return exists
}

func (s *sqlAppointmentStore) PatchDescription(id int, description string) error {
	query := "UPDATE appointments SET description = ? WHERE id = ?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(description, id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("Appointment not found")
	}
	return nil
}
