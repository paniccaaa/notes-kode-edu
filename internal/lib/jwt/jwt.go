package jwt

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/paniccaaa/notes-kode-edu/internal/domain/models"
)

func NewToken(user models.User, duration time.Duration) (string, error) {
	const op = "lib.jwt.NewToken"

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["uid"] = user.ID
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(duration).Unix()

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return tokenString, nil
}

func VerifyToken(tokenStr string) (jwt.MapClaims, error) {
	const op = "lib.jwt.VerifyToken"
	// Parse the token
	parsedToken, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil || !parsedToken.Valid {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	// Extract claims
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok {
		return claims, nil
	}

	return nil, fmt.Errorf("could not extract claims from token")
}
