package note

import (
	"log/slog"
	"net/http"

	"github.com/paniccaaa/notes-kode-edu/internal/storage/postgres"
)

func HandleCreateNote(log *slog.Logger, storage *postgres.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("POST NOTES METHOD HERE"))
	}
}
