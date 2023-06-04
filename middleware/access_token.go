package middleware

import (
	"authGo/config"
	"authGo/service"
	"context"
	"net/http"
)

type Func func(handler http.Handler) http.Handler

func ValidAccessToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		AuthHeader := r.Header.Get("Authorization")
		accessTokenString := service.GetTokenFromBearerString(AuthHeader)

		claims, err := service.ValidateToken(accessTokenString, config.NewConfig().AccessTokenSecret)
		if err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		c := context.WithValue(r.Context(), config.NewConfig().AccessTokenSecret, claims)
		req := r.WithContext(c)

		next.ServeHTTP(w, req)
	})
}
