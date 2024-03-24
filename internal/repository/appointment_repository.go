package repository

import (
	"errors"
	"proyecto_final_go/internal/domain"
	store "proyecto_final_go/pkg/store/appointment"
)

type AppointmentRepository interface {
	Create(appointment domain.Appointment) error
	CreateByPatientDNIAndDentistLicense(patientDNI string, license string, date string, hour string, description string) ([]domain.Appointment, error)
	GetByID(id int) (domain.Appointment, error)
	GetByPatientDNI(patientDNI string) ([]domain.Appointment, error)
	GetAll() ([]domain.Appointment, error)
	Update(appointment domain.Appointment) error
	PatchDescription(id int, description string) error
	Delete(id int) error
}

// ----------------------------------
type appointmentRepository struct {
	storage store.AppointmentStoreInterface
}

func NewAppointmentRepository(storage store.AppointmentStoreInterface) AppointmentRepository {
	return &appointmentRepository{storage}
}

// ----------------------------------
func (r *appointmentRepository) Create(appointment domain.Appointment) error {
	existingAppointments, err := r.storage.GetAll()
	if err != nil {
		return err
	}
	for _, existingAppointment := range existingAppointments {
		if existingAppointment.Patient.Id == appointment.Patient.Id && existingAppointment.Date == appointment.Date && existingAppointment.Hour == appointment.Hour {
			return errors.New("Patient already has an appointment at the same date and time")
		}
		if existingAppointment.Dentist.Id == appointment.Dentist.Id && existingAppointment.Date == appointment.Date && existingAppointment.Hour == appointment.Hour {
			return errors.New("Dentist already has an appointment at the same date and time")
		}
	}
	err = r.storage.Create(appointment)
	if err != nil {
		return err
	}
	return nil
}

func (r *appointmentRepository) CreateByPatientDNIAndDentistLicense(patientDNI string, license string, date string, hour string, description string) ([]domain.Appointment, error) {
	appointments, err := r.storage.CreateByPatientDNIAndDentistLicense(patientDNI, license, date, hour, description)
	if err != nil {
		return nil, err
	}
	return appointments, nil

}

func (r *appointmentRepository) GetByID(id int) (domain.Appointment, error) {
	appointment, err := r.storage.Read(id)
	if err != nil {
		return domain.Appointment{}, errors.New("Appointment not found")
	}
	return appointment, nil
}

func (r *appointmentRepository) GetByPatientDNI(patientDNI string) ([]domain.Appointment, error) {
	allAppointments, err := r.storage.GetAll()
	if err != nil {
		return nil, err
	}
	appointmentsByPatientDNI := make([]domain.Appointment, 0)
	for _, appointment := range allAppointments {
		if appointment.Patient.DNI == patientDNI {
			appointmentsByPatientDNI = append(appointmentsByPatientDNI, appointment)
		}
	}
	return appointmentsByPatientDNI, nil
}

func (r *appointmentRepository) GetAll() ([]domain.Appointment, error) {
	appointments, err := r.storage.GetAll()
	if err != nil {
		return nil, err
	}
	return appointments, nil
}

func (r *appointmentRepository) Update(appointment domain.Appointment) error {
	err := r.storage.Update(appointment)
	if err != nil {
		return errors.New("Error updating appointment")
	}
	return nil
}

func (r *appointmentRepository) PatchDescription(id int, description string) error {
	appointment, err := r.GetByID(id)
	if err != nil {
		return err
	}
	appointment.Description = description
	err = r.storage.Update(appointment)
	if err != nil {
		return err
	}
	return nil
}

func (r *appointmentRepository) Delete(id int) error {
	err := r.storage.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
