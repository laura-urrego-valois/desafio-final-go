package repository

import (
	"errors"
	"proyecto_final_go/internal/domain"
)

type DentistRepository interface {
	Create(dentist *domain.Dentist) (*domain.Dentist, error)
	GetByID(id int) (*domain.Dentist, error)
	GetAll() ([]*domain.Dentist, error)
	Update(dentist *domain.Dentist) error
	UpdateFields(id int, updates map[string]interface{}) error
	Delete(id int) error
}

type dentistRepository struct {
	list []*domain.Dentist
}

// ! NewDentistRepository crea un nuevo repositorio de dentistas
func NewDentistRepository(list []*domain.Dentist) DentistRepository {
	return &dentistRepository{list}
}

// ! Create agrega un nuevo dentista
func (r *dentistRepository) Create(dentist *domain.Dentist) (*domain.Dentist, error) {
	for _, existingDentist := range r.list {
		if existingDentist.License == dentist.License {
			return nil, errors.New("Dentist with the same license already exists")
		}
	}
	dentist.ID = len(r.list) + 1
	r.list = append(r.list, dentist)
	return dentist, nil
}

// ! GetByID busca un dentista por su ID
func (r *dentistRepository) GetByID(id int) (*domain.Dentist, error) {
	for _, d := range r.list {
		if d.ID == id {
			return d, nil
		}
	}
	return nil, errors.New("Dentist not found")
}

// ! GetAll devuelve todos los dentistas
func (r *dentistRepository) GetAll() ([]*domain.Dentist, error) {
	return r.list, nil
}

// ! Update actualiza un dentista
func (r *dentistRepository) Update(dentist *domain.Dentist) error {
	for i, d := range r.list {
		if d.ID == dentist.ID {
			r.list[i] = dentist
			return nil
		}
	}
	return errors.New("Dentist not found")
}

// ! UpdateFields actualiza uno o más campos de un dentista
func (r *dentistRepository) UpdateFields(id int, updates map[string]interface{}) error {
	for _, d := range r.list {
		if d.ID == id {
			for key, value := range updates {
				switch key {
				case "first_name":
					d.FirstName = value.(string)
				case "last_name":
					d.LastName = value.(string)
				case "license":
					d.License = value.(string)
				}
			}
			return nil
		}
	}
	return errors.New("Dentist not found")
}

// ! Delete elimina un dentista por su ID
func (r *dentistRepository) Delete(id int) error {
	for i, d := range r.list {
		if d.ID == id {
			r.list = append(r.list[:i], r.list[i+1:]...)
			return nil
		}
	}
	return errors.New("Dentist not found")
}
