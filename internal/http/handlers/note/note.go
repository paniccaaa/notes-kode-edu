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
	CreateNote(ctx context.Context, note models.Note) (models.Note, error)
}

func HandleGetNotes(log *slog.Logger, note NoteService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		notes, err := note.GetNotes(r.Context())
		if err != nil {
			http.Error(w, "failed to get notes", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(notes)
	}
}

// TODO

type createNoteRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func HandleCreateNote(log *slog.Logger, note NoteService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createNoteRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "failed to decode body", http.StatusBadRequest)
			return
		}

		n := models.Note{
			Title:       req.Title,
			Description: req.Description,
		}

		newNote, err := note.CreateNote(r.Context(), n)
		if err != nil {
			http.Error(w, "failed to create note", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&newNote)
	}
}
