package store

import (
	"database/sql"
	"proyecto-final-go/internal/domain"
)

type sqlStore struct {
	db *sql.DB
}

func NewSqlStore(db *sql.DB) DentistStoreInterface {
	return &sqlStore{
		db: db,
	}
}

//-----------------------------------

func (s *sqlStore) Read(id int) (domain.Dentist, error) {
	var dentist domain.Dentist
	query := "SELECT * FROM dentists WHERE id = ?;"
	row := s.db.QueryRow(query, id)
	err := row.Scan(&dentist.Id, &dentist.FirstName, &dentist.LastName, &dentist.License)
	if err != nil {
		return domain.Dentist{}, err
	}
	return dentist, nil
}

func (s *sqlStore) Create(dentist domain.Dentist) error {
	query := "INSERT INTO dentists (FirstName, LastName, License) VALUES (?, ?, ?);"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(dentist.FirstName, dentist.LastName, dentist.License)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlStore) Update(dentist domain.Dentist) error {
	query := "UPDATE dentists SET FirstName = ?, LastName = ?, License = ? WHERE id = ?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(dentist.FirstName, dentist.LastName, dentist.License, dentist.Id)
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
	query := "DELETE FROM dentists WHERE id = ?;"
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

func (s *sqlStore) Exists(license string) bool {
	var exists bool
	var id int
	query := "SELECT id FROM dentists WHERE License = ?;"
	row := s.db.QueryRow(query, license)
	err := row.Scan(&id)
	if err != nil {
		return false
	}
	if id > 0 {
		exists = true
	}
	return exists
}

func (s *sqlStore) PatchLicense(id int, license string) error {
	query := "UPDATE dentists SET License = ? WHERE id = ?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(license, id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlStore) GetAll() ([]domain.Dentist, error) {
	var dentists []domain.Dentist
	query := "SELECT * FROM dentists;"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var dentist domain.Dentist
		if err := rows.Scan(&dentist.Id, &dentist.FirstName, &dentist.LastName, &dentist.License); err != nil {
			return nil, err
		}
		dentists = append(dentists, dentist)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return dentists, nil
}
