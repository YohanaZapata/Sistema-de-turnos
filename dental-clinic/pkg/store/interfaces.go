package store

import "github.com/bootcamp-go/consignas-go-db.git/internal/domain"

type StoreInterface interface {
    // Read devuelve un dentista por su id
    ReadDentista(id string) (domain.Dentista, error)
    // Create agrega un nuevo dentista
    CreateDentista(dentista domain.Dentista) error
    // Update actualiza un dentista
    UpdateDentista(dentista domain.Dentista) error
    // Delete elimina un dentista
    DeleteDentista(id string) error
    // Exists verifica si un dentista existe
    ExistsDentista(codeValue string) bool
}