package router

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/paniccaaa/notes-kode-edu/internal/http/handlers/auth"
	"github.com/paniccaaa/notes-kode-edu/internal/http/handlers/note"
	"github.com/paniccaaa/notes-kode-edu/internal/http/middlewares/authorization"
)

func InitRouter(log *slog.Logger, noteService note.NoteService, authService auth.AuthService) http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Route("/notes", func(r chi.Router) {
		r.Use(authorization.AuthMiddleware)

		r.Get("/", note.HandleGetNotes(log, noteService))

		r.Post("/", note.HandleCreateNote(log, noteService))
	})

	router.Group(func(r chi.Router) {
		r.Post("/login", auth.HandleLogin(log, authService))

		r.Post("/register", auth.HandleRegister(log, authService))
	})

	return router
}
