package service

import (
	"errors"
	"proyecto_final_go/internal/domain"
	"proyecto_final_go/internal/repository"
)

type AppointmentService interface {
	Create(appointment domain.Appointment) error
	//TODO CreateByPatientDNIAndDentistLicense(patientDNI string, license string) ([]domain.Appointment, error)
	GetByID(id int) (domain.Appointment, error)
	GetByPatientDNI(patientDNI string) ([]domain.Appointment, error)
	GetAll() ([]domain.Appointment, error)
	Update(appointment domain.Appointment) error
	PatchDescription(id int, description string) error
	Delete(id int) error
}

// -------------------------------------------
type appointmentService struct {
	appointmentRepo repository.AppointmentRepository
}

func NewAppointmentService(appointmentRepo repository.AppointmentRepository) AppointmentService {
	return &appointmentService{appointmentRepo}
}

// -------------------------------------------
func (s *appointmentService) Create(appointment domain.Appointment) error {
	existingAppointments, err := s.appointmentRepo.GetAll()
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
	return s.appointmentRepo.Create(appointment)
}

func (s *appointmentService) GetByID(id int) (domain.Appointment, error) {
	appointment, err := s.appointmentRepo.GetByID(id)
	if err != nil {
		return domain.Appointment{}, err
	}
	return appointment, nil
}

func (s *appointmentService) GetByPatientDNI(patientDNI string) ([]domain.Appointment, error) {
	appointments, err := s.appointmentRepo.GetByPatientDNI(patientDNI)
	if err != nil {
		return nil, err
	}
	return appointments, nil
}

func (s *appointmentService) GetAll() ([]domain.Appointment, error) {
	appointments, err := s.appointmentRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return appointments, nil
}

func (s *appointmentService) Update(appointment domain.Appointment) error {
	existingAppointment, err := s.appointmentRepo.GetByID(appointment.Id)
	if err != nil {
		return err
	}
	existingAppointments, err := s.appointmentRepo.GetAll()
	if err != nil {
		return err
	}
	for _, existing := range existingAppointments {
		if existing.Id != appointment.Id {
			if existing.Patient.Id == appointment.Patient.Id &&
				existing.Date == appointment.Date &&
				existing.Hour == appointment.Hour {
				return errors.New("Patient already has an appointment at the same date and time")
			}
			if existing.Dentist.Id == appointment.Dentist.Id &&
				existing.Date == appointment.Date &&
				existing.Hour == appointment.Hour {
				return errors.New("Dentist already has an appointment at the same date and time")
			}
		}
	}

	if appointment.Date != "" {
		existingAppointment.Date = appointment.Date
	}
	if appointment.Hour != "" {
		existingAppointment.Hour = appointment.Hour
	}
	if appointment.Description != "" {
		existingAppointment.Description = appointment.Description
	}
	err = s.appointmentRepo.Update(existingAppointment)
	if err != nil {
		return err
	}

	return nil
}

func (s *appointmentService) PatchDescription(id int, description string) error {
	err := s.appointmentRepo.PatchDescription(id, description)
	if err != nil {
		return err
	}
	return nil
}

func (s *appointmentService) Delete(id int) error {
	err := s.appointmentRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
