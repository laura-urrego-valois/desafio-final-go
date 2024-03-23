package store

import "proyecto-final-go/internal/domain"

type DentistStoreInterface interface {
	Read(id int) (domain.Dentist, error)
	Create(product domain.Dentist) error
	Update(product domain.Dentist) error
	Delete(id int) error
	GetAll() ([]domain.Dentist, error)
	Exists(license string) bool
	PatchLicense(id int, license string) error
}
