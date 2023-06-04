package handler

import (
	"authGo/config"
	"authGo/response"
	"authGo/service"
	"encoding/json"
	"net/http"
)

type RefreshHandler struct {
	cfg *config.Config
}

func NewRefreshHandler(cfg *config.Config) *RefreshHandler {
	return &RefreshHandler{cfg: cfg}
}

func (h *RefreshHandler) GetNewAccessToken(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		AuthHeader := r.Header.Get("Authorization")
		refreshTokenString := service.GetTokenFromBearerString(AuthHeader)

		claims, ok := r.Context().Value(h.cfg.RefreshTokenSecret).(*service.JwtCustomClaims)
		if !ok {
			http.Error(w, "Failed to retrieve claims", http.StatusInternalServerError)
			return
		}

		accessTokenString, err := service.GenerateToken(claims.ID, h.cfg.AccessTokenLifetimeMinutes, h.cfg.AccessTokenSecret)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := response.LoginResponse{AccessToken: accessTokenString, RefreshToken: refreshTokenString}

		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(resp)
	}
}
