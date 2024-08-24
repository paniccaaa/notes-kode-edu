package auth

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/paniccaaa/notes-kode-edu/internal/domain/models"
	"github.com/paniccaaa/notes-kode-edu/internal/lib/jwt"
	"github.com/paniccaaa/notes-kode-edu/internal/storage"
)

type AuthService interface {
	Login(ctx context.Context, email, password string) (string, error)
	Register(ctx context.Context, email, password string) (int64, time.Duration, error)
}

type request struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type responseRegister struct {
	Token string      `json:"token"`
	User  models.User `json:"user"`
}

func HandleRegister(log *slog.Logger, auth AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req request

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "failed to decode request body", http.StatusBadRequest)
			return
		}

		id, timeDuration, err := auth.Register(r.Context(), req.Email, req.Password)
		if err != nil {
			if errors.Is(err, storage.ErrUserExists) {
				http.Error(w, "user already exists", http.StatusConflict)
				return
			}

			http.Error(w, "failed to register", http.StatusInternalServerError)
			return
		}

		user := models.User{ID: id, Email: req.Email}
		token, err := jwt.NewToken(user, timeDuration)
		if err != nil {
			http.Error(w, "failed to generate token", http.StatusInternalServerError)

			return
		}

		res := responseRegister{Token: token, User: user}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&res)
	}
}

type responseLogin struct {
	Token   string `json:"token"`
	Message string `json:"message"`
}

func HandleLogin(log *slog.Logger, auth AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req request

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "failed to decode request body", http.StatusBadRequest)
			return
		}

		token, err := auth.Login(r.Context(), req.Email, req.Password)
		if err != nil {
			http.Error(w, "permission denied", http.StatusUnauthorized)
		}

		res := responseLogin{Token: token, Message: "User successfully logged in"}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&res)
	}
}
