package service

import (
	"proyecto_final_go/internal/domain"
	"proyecto_final_go/internal/repository"
)

type DentistService interface {
	Create(dentist *domain.Dentist) (*domain.Dentist, error)
	GetByID(id int) (*domain.Dentist, error)
	GetAll() ([]*domain.Dentist, error)
	Update(dentist *domain.Dentist) error
	UpdateFields(id int, updates map[string]interface{}) error
	Delete(id int) error
}

type dentistService struct {
	dentistRepo repository.DentistRepository
}

func NewDentistService(dentistRepo repository.DentistRepository) DentistService {
	return &dentistService{dentistRepo}
}

func (s *dentistService) Create(dentist *domain.Dentist) (*domain.Dentist, error) {
	dentist, err := s.dentistRepo.Create(dentist)
	if err != nil {
		return &domain.Dentist{}, err
	}
	return dentist, nil
}

func (s *dentistService) GetByID(id int) (*domain.Dentist, error) {
	dentist, err := s.dentistRepo.GetByID(id)
	if err != nil {
		return &domain.Dentist{}, err
	}
	return dentist, nil
}

func (s *dentistService) GetAll() ([]*domain.Dentist, error) {
	dentists, err := s.dentistRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return dentists, nil
}

func (s *dentistService) Update(dentist *domain.Dentist) error {
	existingDentist, err := s.dentistRepo.GetByID(dentist.ID)
	if err != nil {
		return err
	}
	if dentist.FirstName != "" {
		existingDentist.FirstName = dentist.FirstName
	}
	if dentist.LastName != "" {
		existingDentist.LastName = dentist.LastName
	}
	if dentist.License != "" {
		existingDentist.License = dentist.License
	}
	err = s.dentistRepo.Update(existingDentist)
	if err != nil {

		return err
	}

	return nil
}

func (s *dentistService) UpdateFields(id int, updates map[string]interface{}) error {
	existingDentist, err := s.dentistRepo.GetByID(id)
	if err != nil {
		return err
	}
	for key, value := range updates {
		switch key {
		case "first_name":
			if firstName, ok := value.(string); ok {
				existingDentist.FirstName = firstName
			}
		case "last_name":
			if lastName, ok := value.(string); ok {
				existingDentist.LastName = lastName
			}
		case "license":
			if license, ok := value.(string); ok {
				existingDentist.License = license
			}
		}
	}
	err = s.dentistRepo.Update(existingDentist)
	if err != nil {
		return err
	}
	return nil
}

func (s *dentistService) Delete(id int) error {
	err := s.dentistRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
