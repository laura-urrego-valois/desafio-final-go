package store

import "proyecto_final_go/internal/domain"

type DentistStoreInterface interface {
	Read(id int) (domain.Dentist, error)
	Create(product domain.Dentist) error
	Update(product domain.Dentist) error
	Delete(id int) error
	GetAll() ([]domain.Dentist, error)
	Exists(license string) (bool, error)
	PatchLicense(id int, license string) error
}
