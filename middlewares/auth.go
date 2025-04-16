package middlewares

import (
	"context"
	"movie-watchlist-api/handlers"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		tokenString := strings.Split(authHeader, "Bearer ")[1]
		claims := &handlers.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("your_secret_key"), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "username", claims.Username)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
