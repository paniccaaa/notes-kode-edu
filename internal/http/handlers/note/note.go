package note

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/paniccaaa/notes-kode-edu/internal/domain/models"
)

type NoteService interface {
	GetNotes(ctx context.Context) ([]models.Note, error)
	CreateNote(ctx context.Context, note models.Note) (int, error)
}

func HandleGetNotes(log *slog.Logger, note NoteService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		notes, err := note.GetNotes(r.Context())
		if err != nil {
			http.Error(w, "failed to get notes", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&notes)
	}
}

func HandleCreateNote(log *slog.Logger, note NoteService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("POST NOTES METHOD HERE"))
	}
}
