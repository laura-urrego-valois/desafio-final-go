package repository

import (
	"errors"
	"proyecto_final_go/internal/domain"
	store "proyecto_final_go/pkg/store/patient"
)

// ----------------------------------
type PRepository interface {
	Create(patient domain.Patient) error
	GetByID(id int) (domain.Patient, error)
	GetAll() ([]domain.Patient, error)
	Update(patient domain.Patient) error
	PatchAddress(id int, address string) error
	Delete(id int) error
}

// ----------------------------------
type patientRepository struct {
	storage store.PatientStoreInterface
}

func NewRepository(storage store.PatientStoreInterface) PRepository {
	return &patientRepository{storage}
}

//------------------------------------

func (r *patientRepository) GetByID(id int) (domain.Patient, error) {
	patient, err := r.storage.Read(id)
	if err != nil {
		return domain.Patient{}, errors.New("patient not found")
	}
	return patient, nil

}

func (r *patientRepository) Create(p domain.Patient) error {
	if !r.storage.Exists(p.DNI) {
		return errors.New("DNI already exists")
	}
	err := r.storage.Create(p)
	if err != nil {
		return err
	}
	return nil
}

func (r *patientRepository) Delete(id int) error {
	err := r.storage.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (r *patientRepository) Update(p domain.Patient) error {
	if !r.storage.Exists(p.DNI) {
		return errors.New("DNI already exists")
	}
	err := r.storage.Update(p)
	if err != nil {
		return errors.New("error updating product")
	}
	return nil
}

func (r *patientRepository) GetAll() ([]domain.Patient, error) {
	patients, err := r.storage.GetAll()
	if err != nil {
		return nil, err
	}
	return patients, nil
}

func (r *patientRepository) PatchAddress(id int, address string) error {
	patient, err := r.storage.Read(id)
	if err != nil {
		return errors.New("patient not found")
	}
	patient.Address = address
	err = r.storage.PatchAddress(patient.Id, patient.Address)
	if err != nil {
		return err
	}
	return nil
}
