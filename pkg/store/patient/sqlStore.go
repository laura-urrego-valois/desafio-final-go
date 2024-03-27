package store

import (
	"database/sql"
	"errors"
	"proyecto_final_go/internal/domain"
)

type sqlStore struct {
	db *sql.DB
}

func NewSqlStore(db *sql.DB) PatientStoreInterface {
	return &sqlStore{
		db: db,
	}
}

//-----------------------------------

func (s *sqlStore) Read(id int) (domain.Patient, error) {
	var patient domain.Patient
	query := "SELECT * FROM patients WHERE Id = ?;"
	row := s.db.QueryRow(query, id)
	err := row.Scan(&patient.Id, &patient.FirstName, &patient.LastName, &patient.Address, &patient.DNI, &patient.ReleaseDate)
	if err != nil {
		return domain.Patient{}, err
	}
	return patient, nil
}

func (s *sqlStore) Create(patient domain.Patient) error {
	query := "INSERT INTO patients (FirstName, LastName, Address, DNI, ReleaseDate) VALUES (?, ?, ?, ?, ?);"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(patient.FirstName, patient.LastName, patient.Address, patient.DNI, patient.ReleaseDate)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no rows affected")
	}

	return nil
}

func (s *sqlStore) Update(patient domain.Patient) error {
	query := "UPDATE patients SET FirstName = ?, LastName = ?, Address = ?, DNI = ?, ReleaseDate = ? WHERE Id = ?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(patient.FirstName, patient.LastName, patient.Address, patient.DNI, patient.ReleaseDate, patient.Id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlStore) Delete(id int) error {
	query := "DELETE FROM patients WHERE Id = ?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlStore) Exists(dni string) (bool, error) {
	var id int
	query := "SELECT Id FROM patients WHERE DNI = ?;"
	row := s.db.QueryRow(query, dni)
	err := row.Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return id > 0, nil
}

func (s *sqlStore) PatchAddress(id int, address string) error {
	query := "UPDATE patients SET Address = ? WHERE Id = ?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(address, id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlStore) GetAll() ([]domain.Patient, error) {
	var patients []domain.Patient
	query := "SELECT * FROM patients;"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var patient domain.Patient
		if err := rows.Scan(&patient.Id, &patient.FirstName, &patient.LastName, &patient.Address, &patient.DNI, &patient.ReleaseDate); err != nil {
			return nil, err
		}
		patients = append(patients, patient)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return patients, nil
}
