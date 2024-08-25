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

func (s *Storage) GetNotes(ctx context.Context, userID int64) ([]models.Note, error) {
	const op = "storage.postgres.GetNotes"

	notes := []models.Note{}
	query := "SELECT * FROM notes WHERE user_id = $1;"

	rows, err := s.Db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	defer rows.Close()

	for rows.Next() {
		var n models.Note

		if err := rows.Scan(&n.ID, &n.UserID, &n.Title, &n.Description); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		notes = append(notes, n)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return notes, nil
}

func (s *Storage) CreateNote(ctx context.Context, note models.Note) (models.Note, error) {
	const op = "storage.postgres.CreateNote"

	query := "INSERT INTO notes (user_id, title, description) VALUES ($1, $2, $3) returning id;"

	var id int
	if err := s.Db.QueryRowContext(ctx, query, note.UserID, note.Title, note.Description).Scan(&id); err != nil {
		return models.Note{}, fmt.Errorf("%s: %w", op, err)
	}

	note.ID = int64(id)

	return note, nil
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
