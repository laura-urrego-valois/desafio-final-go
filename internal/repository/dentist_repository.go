package repository

import (
	"errors"
	"proyecto_final_go/internal/domain"

	store "proyecto_final_go/pkg/store/dentist"
)

// ----------------------------------
type DentistRepository interface {
	Create(dentist domain.Dentist) error
	GetByID(id int) (domain.Dentist, error)
	GetAll() ([]domain.Dentist, error)
	Update(dentist domain.Dentist) error
	PatchLicense(id int, license string) error
	Delete(id int) error
}

// ----------------------------------
type dentistRepository struct {
	storage store.DentistStoreInterface
}

func NewDentistRepository(storage store.DentistStoreInterface) DentistRepository {
	return &dentistRepository{storage}
}

// ----------------------------------

func (r *dentistRepository) Create(dentist domain.Dentist) error {
	if !r.storage.Exists(dentist.License) {
		return errors.New("License already exists")
	}
	err := r.storage.Create(dentist)
	if err != nil {
		return err
	}
	return nil
}

func (r *dentistRepository) GetByID(id int) (domain.Dentist, error) {
	dentist, err := r.storage.Read(id)
	if err != nil {
		return domain.Dentist{}, errors.New("Dentist not found")
	}
	return dentist, nil
}

func (r *dentistRepository) GetAll() ([]domain.Dentist, error) {
	dentists, err := r.storage.GetAll()
	if err != nil {
		return nil, err
	}
	return dentists, nil
}

func (r *dentistRepository) Update(dentist domain.Dentist) error {
	if !r.storage.Exists(dentist.License) {
		return errors.New("License already exists")
	}
	err := r.storage.Update(dentist)
	if err != nil {
		return errors.New("Error updating dentist")
	}
	return nil
}

func (r *dentistRepository) PatchLicense(id int, license string) error {
	dentist, err := r.storage.Read(id)
	if err != nil {
		return errors.New("Dentist not found")
	}
	dentist.License = license
	err = r.storage.PatchLicense(dentist.Id, dentist.License)
	if err != nil {
		return err
	}
	return nil

}

func (r *dentistRepository) Delete(id int) error {
	err := r.storage.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
