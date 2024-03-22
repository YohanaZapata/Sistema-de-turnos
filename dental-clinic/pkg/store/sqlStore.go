package store

import (
	"database/sql"

	"github.com/bootcamp-go/consignas-go-db.git/internal/domain"
)

type sqlStore struct {
	db *sql.DB
}

func NewSqlStore(db *sql.DB) StoreInterface {
	func (s *sqlStore) Create(dentista domain.Dentista) error
func (s *sqlStore) Delete(id int) error
func (s *sqlStore) Exists(matricula string) bool
func (s *sqlStore) Read(id int) (domain.Dentista, error)
func (s *sqlStore) Update(dentista domain.Dentista) error
	return &sqlStore{
		db: db,
	}
}

func (s *sqlStore) Read(id int) (domain.Dentista, error) {
	var dentista domain.Dentista
	query := "SELECT * FROM dentistas WHERE id = ?;"
	row := s.db.QueryRow(query, id)
	err := row.Scan(&dentista.ID, &dentista.Nombre, &dentista.Apellido, &dentista.Matricula)
	if err != nil {
		return domain.Dentista{}, err
	}
	return dentista, nil
}

func (s *sqlStore) Create(dentista domain.Dentista) error {
	query := "INSERT INTO dentistas (nombre, apellido, matricula) VALUES (?, ?, ?);"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(dentista.Nombre, dentista.Apellido, dentista.Matricula)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlStore) Update(dentista domain.Dentista) error {
	query := "UPDATE dentistas SET nombre = ?, apellido = ?, matricula = ? WHERE id = ?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(dentista.Nombre, dentista.Apellido, dentista.Matricula, dentista.ID)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlStore) Delete(id int) error {
	query := "DELETE FROM dentistas WHERE id = ?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlStore) Exists(matricula string) bool {
	var exists bool
	var count int
	query := "SELECT COUNT(*) FROM dentistas WHERE matricula = ?;"
	row := s.db.QueryRow(query, matricula)
	err := row.Scan(&count)
	if err != nil {
		return false
	}
	exists = count > 0
	return exists
}
