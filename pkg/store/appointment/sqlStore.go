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
			a.Id, a.Date, a.Hour, a.Description,
			p.Id AS patient_id, p.FirstName AS patient_first_name, p.LastName AS patient_last_name, p.Address AS patient_address, p.DNI AS patient_dni, p.ReleaseDate AS patient_release_date,
			d.Id AS dentist_id, d.FirstName AS dentist_first_name, d.LastName AS dentist_last_name, d.License AS dentist_license
		FROM 
			appointments AS a
		INNER JOIN 
			patients AS p ON a.patients_Id = p.Id
		INNER JOIN 
			dentists AS d ON a.dentists_Id = d.Id
		WHERE 
			a.Id = ?;
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
		a.Id, a.Date, a.Hour, a.Description,
		p.Id AS patient_id, p.FirstName AS patient_first_name, p.LastName AS patient_last_name, p.Address AS patient_address, p.DNI AS patient_dni, p.ReleaseDate AS patient_release_date,
		d.Id AS dentist_id, d.FirstName AS dentist_first_name, d.LastName AS dentist_last_name, d.License AS dentist_license
	FROM 
		appointments AS a
	INNER JOIN 
		patients AS p ON a.patients_Id = p.Id
	INNER JOIN 
		dentists AS d ON a.dentists_Id = d.Id
	WHERE 
		p.DNI = ?;
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
		INSERT INTO appointments (Date, Hour, Description, patients_Id, dentists_Id) 
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
	patientQuery := "SELECT Id FROM patients WHERE DNI = ?"
	err := s.db.QueryRow(patientQuery, patientDNI).Scan(&patientID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Patient not found")
		}
		return nil, err
	}

	var dentistID int
	dentistQuery := "SELECT Id FROM dentists WHERE License = ?"
	err = s.db.QueryRow(dentistQuery, license).Scan(&dentistID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Dentist not found")
		}
		return nil, err
	}

	createQuery := "INSERT INTO appointments (patients_Id, dentists_Id, Date, Hour, Description) VALUES (?, ?, ?, ?, ?)"
	_, err = s.db.Exec(createQuery, patientID, dentistID, date, hour, description)
	if err != nil {
		return nil, err
	}

	appointmentsQuery := `
		SELECT 
			a.Id, a.Date, a.Hour, a.Description,
			p.Id AS patient_id, p.FirstName AS patient_first_name, p.LastName AS patient_last_name, p.Address AS patient_address, p.DNI AS patient_dni, p.ReleaseDate AS patient_release_date,
			d.Id AS dentist_id, d.FirstName AS dentist_first_name, d.LastName AS dentist_last_name, d.License AS dentist_license
		FROM 
			appointments AS a
		INNER JOIN 
			patients AS p ON a.patients_Id = p.Id
		INNER JOIN 
			dentists AS d ON a.dentists_Id = d.Id
		WHERE 
			a.patients_Id = ? AND a.dentists_Id = ?
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
		SET Date = ?, Hour = ?, Description = ?, patients_Id = ?, dentists_Id = ?
		WHERE Id = ?;
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
		WHERE Id = ?;
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
		a.Id, a.Date, a.Hour, a.Description,
		p.Id AS patient_id, p.FirstName AS patient_first_name, p.LastName AS patient_last_name, p.Address AS patient_address, p.DNI AS patient_dni, p.ReleaseDate AS patient_release_date,
		d.Id AS dentist_id, d.FirstName AS dentist_first_name, d.LastName AS dentist_last_name, d.License AS dentist_license
	FROM 
		appointments AS a
	INNER JOIN 
		patients AS p ON a.patients_Id = p.Id
	INNER JOIN 
		dentists AS d ON a.dentists_Id = d.Id
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

func (s *sqlAppointmentStore) Exists(id int) (bool, error) {
	var appointmentId int
	query := "SELECT Id FROM appointments WHERE Id = ?;"
	row := s.db.QueryRow(query, appointmentId)
	err := row.Scan(&appointmentId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return appointmentId > 0, nil
}

func (s *sqlAppointmentStore) PatchDescription(id int, description string) error {
	query := "UPDATE appointments SET Description = ? WHERE Id = ?;"
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
