package router

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/paniccaaa/notes-kode-edu/internal/http/handlers/note"
)

func InitRouter(log *slog.Logger, noteService note.NoteService) http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	//	router.Use(authMiddleware)

	router.Route("/notes", func(r chi.Router) {
		r.Get("/", note.HandleGetNotes(log, noteService))

		r.Post("/", note.HandleCreateNote(log, noteService))
	})

	return router
}

// SEE THIS: https://github.com/go-chi/jwtauth
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//w.Write([]byte("AUTH MIDDLEWARE"))
		slog.Info("INFO INFO")
	})
}
