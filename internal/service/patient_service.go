package service

import (
	"proyecto_final_go/internal/domain"
	"proyecto_final_go/internal/repository"
)

type PatientService interface {
	Create(patient domain.Patient) error
	GetByID(id int) (domain.Patient, error)
	GetAll() ([]domain.Patient, error)
	Update(patient domain.Patient) error
	PatchAddress(id int, address string) error
	Delete(id int) error
}

// -------------------------------------------
type patientService struct {
	r repository.PatientRepository
}

func NewPatientService(r repository.PatientRepository) PatientService {
	return &patientService{r}
}

//-------------------------------------------

func (s *patientService) GetByID(id int) (domain.Patient, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Patient{}, err
	}
	return p, nil
}

func (s *patientService) Create(p domain.Patient) error {
	err := s.r.Create(p)
	if err != nil {
		return err
	}
	return nil
}

func (s *patientService) Update(pa domain.Patient) error {
	p, err := s.r.GetByID(pa.Id)
	if err != nil {
		return err
	}
	if pa.FirstName != "" {
		p.FirstName = pa.FirstName
	}
	if pa.LastName != "" {
		p.LastName = pa.LastName
	}
	if pa.Address != "" {
		p.Address = pa.Address
	}
	if pa.DNI != "" {
		p.DNI = pa.DNI
	}
	if pa.ReleaseDate != "" {
		p.ReleaseDate = pa.ReleaseDate
	}
	err = s.r.Update(p)
	if err != nil {
		return err
	}
	return nil
}

func (s *patientService) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *patientService) GetAll() ([]domain.Patient, error) {
	patients, err := s.r.GetAll()
	if err != nil {
		return nil, err
	}
	return patients, nil
}

func (s *patientService) PatchAddress(id int, address string) error {
	err := s.r.PatchAddress(id, address)
	if err != nil {
		return err
	}
	return nil
}
