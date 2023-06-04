package handler

import (
	"authGo/config"
	"authGo/repository"
	"authGo/request"
	"authGo/response"
	"authGo/service"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type AuthHandler struct {
	cfg *config.Config
}

func NewAuthHandler(cfg *config.Config) *AuthHandler {
	return &AuthHandler{
		cfg: cfg,
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		req := new(request.LoginRequest)
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user, err := repository.NewUserRepository().GetUserByEmail(req.Email)
		if err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		accessString, err := service.GenerateToken(user.ID, h.cfg.AccessTokenLifetimeMinutes, h.cfg.AccessTokenSecret)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		refreshString, err := service.GenerateToken(user.ID, h.cfg.RefreshTokenLifetimeMinutes, h.cfg.RefreshTokenSecret)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := response.LoginResponse{
			AccessToken:  accessString, // secret token is one direction
			RefreshToken: refreshString,
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)

	default:
		http.Error(w, "only post is allowed", http.StatusMethodNotAllowed)
	}
}
