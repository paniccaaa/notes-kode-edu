package router

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/paniccaaa/notes-kode-edu/internal/http/handlers/note"
	"github.com/paniccaaa/notes-kode-edu/internal/storage/postgres"
)

func InitRouter(log *slog.Logger, storage *postgres.Storage) http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	//	router.Use(authMiddleware)

	router.Route("/notes", func(r chi.Router) {
		r.Get("/", note.HandleGetNotes(log, storage))

		r.Post("/", note.HandleCreateNote(log, storage))
	})

	return router
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//w.Write([]byte("AUTH MIDDLEWARE"))
		slog.Info("INFO INFO")
	})
}
