package authservice

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/paniccaaa/notes-kode-edu/internal/domain/models"
	"github.com/paniccaaa/notes-kode-edu/internal/lib/jwt"
	"github.com/paniccaaa/notes-kode-edu/internal/storage"
	"golang.org/x/crypto/bcrypt"
)

type Storage interface {
	SaveUser(ctx context.Context, passHash []byte, email string) (int64, error)
	User(ctx context.Context, email string) (models.User, error)
}

type AuthService struct {
	storage  Storage
	log      *slog.Logger
	tokenTTL time.Duration
}

func NewAuthService(storage Storage, log *slog.Logger, tokenTTL time.Duration) *AuthService {
	return &AuthService{storage: storage, log: log, tokenTTL: tokenTTL}
}

func (s *AuthService) Login(ctx context.Context, email, password string) (string, error) {
	const op = "services.auth-service.Login"

	log := s.log.With(slog.String("op", op))

	user, err := s.storage.User(ctx, email)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			log.Warn("user not found", slog.String("error", err.Error()))

			return "", fmt.Errorf("%s: %w", op, err)
		}

		log.Error("failed to get user", slog.String("error", err.Error()))

		return "", fmt.Errorf("%s: %w", op, err)
	}

	if err := bcrypt.CompareHashAndPassword(user.PassHash, []byte(password)); err != nil {
		log.Info("invalid credentials", slog.String("error", err.Error()))

		return "", fmt.Errorf("%s: %w", op, err)
	}

	token, err := jwt.NewToken(user, s.tokenTTL)
	if err != nil {
		log.Error("failed to generate token", slog.String("error", err.Error()))

		return "", fmt.Errorf("%s: %w", op, err)
	}

	return token, nil
}

func (s *AuthService) Register(ctx context.Context, email, password string) (int64, time.Duration, error) {
	const op = "services.auth-service.Register"

	log := s.log.With(slog.String("op", op))

	passHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("failed to generate password hash", slog.String("error", err.Error()))

		return 0, 0, fmt.Errorf("%s: %w", op, err)
	}

	id, err := s.storage.SaveUser(ctx, passHash, email)
	if err != nil {
		if errors.Is(err, storage.ErrUserExists) {
			log.Warn("user already exists", slog.String("error", err.Error()))
			return 0, 0, fmt.Errorf("%s: %w", op, err)
		}

		log.Error("failed to save user", slog.String("error", err.Error()))

		return 0, 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, s.tokenTTL, nil
}
