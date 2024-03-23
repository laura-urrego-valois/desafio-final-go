package store

import (
	"proyecto_final_go/internal/domain"
)

type PatientStoreInterface interface {
	Read(id int) (domain.Patient, error)
	Create(product domain.Patient) error
	Update(product domain.Patient) error
	Delete(id int) error
	GetAll() ([]domain.Patient, error)
	Exists(dni string) bool
	PatchAddress(id int, address string) error
}
