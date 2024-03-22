package store

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/bootcamp-go/consignas-go-db.git/internal/domain"
)

type jsonStore struct {
	pathToFile string
}

// loadDentistas carga los dentistas desde un archivo json
func (s *jsonStore) loadDentistas() ([]domain.Dentista, error) {
	var dentistas []domain.Dentista
	file, err := os.ReadFile(s.pathToFile)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(file), &dentistas)
	if err != nil {
		return nil, err
	}
	return dentistas, nil
}

// saveDentistas guarda los dentistas en un archivo json
func (s *jsonStore) saveDentistas(dentistas []domain.Dentista) error {
	bytes, err := json.Marshal(dentistas)
	if err != nil {
		return err
	}
	return os.WriteFile(s.pathToFile, bytes, 0644)
}

func (s *jsonStore) CreateDentista(dentista domain.Dentista) error {
	dentistas, err := s.loadDentistas()
	if err != nil {
		return err
	}
	dentista.ID = len(dentistas) + 1
	dentistas = append(dentistas, dentista)
	return s.saveDentistas(dentistas)
}

// NewJsonStore crea un nuevo store de dentistas
func NewJsonStore(path string) StoreInterface {
	_, err := os.Stat(path)
	if err != nil {
		panic(err)
	}
	return &jsonStore{
		pathToFile: path,
	}
}

func (s *jsonStore) Read(id int) (domain.Dentista, error) {
	dentistas, err := s.loadDentistas()
	if err != nil {
		return domain.Dentista{}, err
	}
	for _, dentista := range dentistas {
		if dentista.ID == id {
			return dentista, nil
		}
	}
	return domain.Dentista{}, errors.New("dentista not found")
}

func (s *jsonStore) Create(dentista domain.Dentista) error {
	dentistas, err := s.loadDentistas()
	if err != nil {
		return err
	}
	dentista.ID = len(dentistas) + 1
	dentistas = append(dentistas, dentista)
	return s.saveDentistas(dentistas)
}

func (s *jsonStore) Update(dentista domain.Dentista) error {
	dentistas, err := s.loadDentistas()
	if err != nil {
		return err
	}
	for i, d := range dentistas {
		if d.ID == dentista.ID {
			dentistas[i] = dentista
			return s.saveDentistas(dentistas)
		}
	}
	return errors.New("dentista not found")
}

func (s *jsonStore) Delete(id int) error {
	dentistas, err := s.loadDentistas()
	if err != nil {
		return err
	}
	for i, d := range dentistas {
		if d.ID == id {
			dentistas = append(dentistas[:i], dentistas[i+1:]...)
			return s.saveDentistas(dentistas)
		}
	}
	return errors.New("dentista not found")}

func (s *jsonStore) Exists(codeValue string) bool {
	dentistas, err := s.loadDentistas()
	if err != nil {
		return false
	}
	for _, d := range dentistas {
		if d.CodeValue == codeValue {
			return true
		}
	}
	return false
}

