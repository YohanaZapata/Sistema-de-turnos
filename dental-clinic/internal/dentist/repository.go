package dentist

import (
	"fmt"

	"github.com/bootcamp-go/consignas-go-db.git/internal/domain"
	"github.com/bootcamp-go/consignas-go-db.git/pkg/store"
)

// DentistaRepository interface
type DentistaRepository interface {
	GetDentistaByID(id string) (*domain.Dentista, error)
	GetAllDentistas() ([]domain.Dentista, error)
	CreateDentista(dentista *domain.Dentista) error
	UpdateDentista(dentista *domain.Dentista) error
	DeleteDentista(id string) error
}

// InMemoryRepository struct
type InMemoryRepository struct{
	dentistas []domain.Dentista
}

// Implementing DentistaRepository on InMemoryRepository struct
func (r *InMemoryRepository) GetDentistaByID(id string) (*domain.Dentista, error) {
	for i, dentista := range r.dentistas {
		if dentista.ID == id {
			return &r.dentistas[i], nil
		}
	}

	return nil, fmt.Errorf("Dentista with id %s not found", id)
}

func (r *InMemoryRepository) GetAllDentistas() ([]domain.Dentista, error) {
	return r.dentistas, nil
}

func (r *InMemoryRepository) CreateDentista(dentista *domain.Dentista) error {
	r.dentistas = append(r.dentistas, *dentista)
	return nil
}

func (r *InMemoryRepository) UpdateDentista(dentista *domain.Dentista) error {
	var dentistas []domain.Dentista

	for _, v := range r.dentistas {
		if v.ID == dentista.ID {
			dentistas = append(dentistas, *dentista)
		} else {
			dentistas = append(dentistas, v)
		}
	}

	r.dentistas = dentistas
	return nil
}

func (r *InMemoryRepository) DeleteDentista(id string) error {
	var dentistas []domain.Dentista

	for _, v := range r.dentistas {
		if v.ID != id {
			dentistas = append(dentistas, v)
		}
	}

	r.dentistas = dentistas
	return nil
}
