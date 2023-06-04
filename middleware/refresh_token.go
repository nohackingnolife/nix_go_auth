package middleware

import (
	"authGo/config"
	"authGo/service"
	"context"
	"net/http"
)

func ValidRefreshToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		AuthHeader := r.Header.Get("Authorization")
		refreshTokenString := service.GetTokenFromBearerString(AuthHeader)

		claims, err := service.ValidateToken(refreshTokenString, config.NewConfig().RefreshTokenSecret)
		if err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		c := context.WithValue(r.Context(), config.NewConfig().RefreshTokenSecret, claims)
		req := r.WithContext(c)

		next.ServeHTTP(w, req)
	})
}
