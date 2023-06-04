package handler

import (
	"authGo/config"
	"authGo/repository"
	"authGo/request"
	"authGo/response"
	"authGo/service"
	"encoding/json"
	"net/http"
)

type RegisterHandler struct {
	cfg *config.Config
}

func NewRegisterHandler(cfg *config.Config) *AuthHandler {
	return &AuthHandler{
		cfg: cfg,
	}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		req := new(request.RegisterRequest)
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// check if email already exist
		user, err := repository.NewUserRepository().RegisterUser(req.Name, req.Email, req.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		//fmt.Println(user)
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
