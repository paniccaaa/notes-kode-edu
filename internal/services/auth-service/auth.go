package authservice

import (
	"context"
	"log/slog"
	"time"

	"github.com/paniccaaa/notes-kode-edu/internal/domain/models"
	"github.com/paniccaaa/notes-kode-edu/internal/storage/postgres"
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

func NewAuthService(storage *postgres.Storage, log *slog.Logger, tokenTTL time.Duration) *AuthService {
	return &AuthService{storage: storage, log: log, tokenTTL: tokenTTL}
}

func (s *AuthService) Login(ctx context.Context, email, password string) (models.User, error) {
	return s.storage.User(ctx, email)
}

func (s *AuthService) Register(ctx context.Context, email, password string) (int64, error) {
	return s.storage.SaveUser(ctx, []byte(password), email)
}
