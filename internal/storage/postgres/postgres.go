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

func (s *Storage) CreateNote(ctx context.Context, note models.Note) (int, error) {
	return 0, nil
}

func (s *Storage) SaveUser(ctx context.Context, passHash []byte, email string) (int64, error) {
	const op = "storage.postgres.SaveUser"

	var id int64
	query := "INSERT INTO users (email, pass_hash) VALUES ($1, $2) RETURNING id"

	if err := s.Db.QueryRowContext(ctx, query, email, passHash).Scan(&id); err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

func (s *Storage) User(ctx context.Context, email string) (models.User, error) {
	const op = "storage.postgres.User"

	var user models.User

	row := s.Db.QueryRowContext(ctx, "SELECT * FROM users WHERE email = $1;", email)
	if err := row.Scan(&user.ID, &user.Email, &user.PassHash); err != nil {
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}
	return user, nil
}
