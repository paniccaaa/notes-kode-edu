package noteservice

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/golang-jwt/jwt"
	"github.com/paniccaaa/notes-kode-edu/internal/domain/models"
	"github.com/paniccaaa/notes-kode-edu/internal/storage/postgres"
)

type Storage interface {
	GetNotes(ctx context.Context, userID int64) ([]models.Note, error)
	CreateNote(ctx context.Context, note models.Note) (int, error)
}

type NoteService struct {
	storage Storage
	log     *slog.Logger
}

func NewNoteService(storage *postgres.Storage, log *slog.Logger) *NoteService {
	return &NoteService{storage: storage, log: log}
}

func (s *NoteService) GetNotes(ctx context.Context) ([]models.Note, error) {
	const op = "services.note-service.GetNotes"

	log := s.log.With(slog.String("op", op))

	var userClaims jwt.MapClaims

	userClaims, ok := ctx.Value("userClaims").(jwt.MapClaims)
	if !ok {
		log.Error("failed to retrieve from ctx")

		return nil, fmt.Errorf("%s: failed to retrieve from ctx", op)
	}

	userIDFloat := userClaims["uid"].(float64)
	userID := int64(userIDFloat)

	notes, err := s.storage.GetNotes(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return notes, nil
}

func (s *NoteService) CreateNote(ctx context.Context, note models.Note) (int, error) {
	return s.storage.CreateNote(ctx, note)
}
