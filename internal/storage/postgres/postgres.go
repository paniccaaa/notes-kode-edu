package postgres

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/paniccaaa/notes-kode-edu/internal/domain/models"
)

type Storage struct {
	Db *sql.DB
}

func NewPostgres(dbURI string) (*Storage, error) {
	const op = "storage.postgres.NewPostgres"

	db, err := sql.Open("postgres", dbURI)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{Db: db}, nil
}

func (s *Storage) GetNotes(ctx context.Context) ([]models.Note, error) {
	return nil, nil
}

// Логика получения заметок из базы данных

func (s *Storage) CreateNote(ctx context.Context, note models.Note) (int, error) {
	return 0, nil
}

// Логика создания новой заметки в базе данных
