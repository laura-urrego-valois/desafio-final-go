package store

import (
	"proyecto_final_go/internal/domain"
)

type AppointmentStoreInterface interface {
	Read(id int) (domain.Appointment, error)
	ReadByPatientDNI(patientDNI string) ([]domain.Appointment, error)
	Create(appointment domain.Appointment) error
	CreateByPatientDNIAndDentistLicense(patientDNI string, license string, date string, hour string, description string) ([]domain.Appointment, error)
	Update(appointment domain.Appointment) error
	Delete(id int) error
	GetAll() ([]domain.Appointment, error)
	Exists(id int) (bool, error)
	PatchDescription(id int, description string) error
}
