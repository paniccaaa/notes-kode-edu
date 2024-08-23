package note

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/paniccaaa/notes-kode-edu/internal/domain/models"
)

type NoteService interface {
	GetNotes(ctx context.Context) ([]models.Note, error)
	CreateNote(ctx context.Context, note models.Note) (int, error)
}

func HandleGetNotes(log *slog.Logger, storage NoteService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("HELLO NOTES"))
	}
}

func HandleCreateNote(log *slog.Logger, storage NoteService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("POST NOTES METHOD HERE"))
	}
}
