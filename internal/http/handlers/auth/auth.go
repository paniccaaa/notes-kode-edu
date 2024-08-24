package auth

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/paniccaaa/notes-kode-edu/internal/domain/models"
)

type AuthService interface {
	Login(ctx context.Context, email, password string) (models.User, error)
	Register(ctx context.Context, email, password string) (int64, error)
}

func HandleLogin(log *slog.Logger, storage AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("HELLO NOTES"))
	}
}

func HandleRegister(log *slog.Logger, storage AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("POST NOTES METHOD HERE"))
	}
}
