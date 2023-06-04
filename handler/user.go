package handler

import (
	"authGo/config"
	"authGo/repository"
	"authGo/response"
	"authGo/service"
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	cfg *config.Config
}

func NewUserHandler(cfg *config.Config) *UserHandler {
	return &UserHandler{
		cfg: cfg,
	}
}

func (h *UserHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		claims, ok := r.Context().Value(h.cfg.AccessTokenSecret).(*service.JwtCustomClaims)
		if !ok {
			http.Error(w, "Failed to retrieve claims", http.StatusInternalServerError)
			return
		}

		user, err := repository.NewUserRepository().GetUserById(claims.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		resp := response.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		} // separate structure

		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(resp) // encode retrn error , we should check it

	default:
		http.Error(w, "only get method allowed", http.StatusMethodNotAllowed)
	}
}
