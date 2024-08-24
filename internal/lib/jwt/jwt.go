package jwt

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"github.com/paniccaaa/notes-kode-edu/internal/domain/models"
)

func NewToken(user models.User, duration time.Duration) (string, error) {
	const op = "lib.jwt.NewToken"

	err := godotenv.Load()
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["uid"] = user.ID
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(duration).Unix()

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
