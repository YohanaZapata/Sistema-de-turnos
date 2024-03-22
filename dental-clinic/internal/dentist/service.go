package dentist 

import (
	"github.com/bootcamp-go/consignas-go-db.git/internal/domain"
)



// DentistaService interface
type DentistaService interface {
	GetDentista(id string) (*Dentista, error)
	GetAllDentistas() ([]Dentista, error)
	CreateDentista(dentista *Dentista) error
	UpdateDentista(dentista *Dentista) error
	DeleteDentista(id string) error
}

// Service struct
type Service struct {
	repository DentistaRepository
}

// Implementing DentistaService on Service struct
func (s *Service) GetDentista(id string) (*Dentista, error) {
	return s.repository.GetDentistaByID(id)
}

func (s *Service) GetAllDentistas() ([]Dentista, error) {
	return s.repository.GetAllDentistas()
}

func (s *Service) CreateDentista(dentista *Dentista) error {
	return s.repository.CreateDentista(dentista)
}

func (s *Service) UpdateDentista(dentista *Dentista) error {
	return s.repository.UpdateDentista(dentista)
}

func (s *Service) DeleteDentista(id string) error {
	return s.repository.DeleteDentista(id)
}
