package noteservice

import (
	"context"

	"github.com/paniccaaa/notes-kode-edu/internal/domain/models"
	"github.com/paniccaaa/notes-kode-edu/internal/storage/postgres"
)

type Storage interface {
	GetNotes(ctx context.Context) ([]models.Note, error)
	CreateNote(ctx context.Context, note models.Note) (int, error)
}

type NoteService struct {
	storage Storage
}

func NewNoteService(storage *postgres.Storage) *NoteService {
	return &NoteService{storage: storage}
}

func (s *NoteService) GetNotes(ctx context.Context) ([]models.Note, error) {
	return s.storage.GetNotes(ctx)
}

func (s *NoteService) CreateNote(ctx context.Context, note models.Note) (int, error) {
	return s.storage.CreateNote(ctx, note)
}
