package authorization

import (
	"context"
	"net/http"
	"strings"

	"github.com/paniccaaa/notes-kode-edu/internal/lib/jwt"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "user is not authorized", http.StatusUnauthorized)
			return
		}

		const bearerPrefix = "Bearer "
		if !strings.HasPrefix(authHeader, bearerPrefix) {
			http.Error(w, "invalid token format", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, bearerPrefix)
		if tokenString == "" {
			http.Error(w, "invalid token format", http.StatusUnauthorized)
			return
		}

		userClaims, err := jwt.VerifyToken(tokenString)
		if err != nil {
			http.Error(w, "permission denied", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "userClaims", userClaims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
