package noteservice

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/golang-jwt/jwt"
	"github.com/paniccaaa/notes-kode-edu/internal/domain/models"
)

type Storage interface {
	GetNotes(ctx context.Context, userID int64) ([]models.Note, error)
	CreateNote(ctx context.Context, note models.Note) (models.Note, error)
}

type NoteService struct {
	storage Storage
	log     *slog.Logger
}

func NewNoteService(storage Storage, log *slog.Logger) *NoteService {
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

func (s *NoteService) CreateNote(ctx context.Context, note models.Note) (models.Note, error) {
	const op = "services.note-service.CreateNote"

	log := s.log.With(slog.String("op", op))

	var userClaims jwt.MapClaims
	userClaims, ok := ctx.Value("userClaims").(jwt.MapClaims)
	if !ok {
		log.Error("failed to retrieve from ctx")

		return models.Note{}, fmt.Errorf("%s: failed to retrieve from ctx", op)
	}

	userIDFloat := userClaims["uid"].(float64)
	userID := int64(userIDFloat)

	note.UserID = userID

	note, err := s.storage.CreateNote(ctx, note)
	if err != nil {
		log.Error("failed to create note", slog.String("err", err.Error()))

		return models.Note{}, fmt.Errorf("%s: %w", op, err)
	}

	return note, nil
}
