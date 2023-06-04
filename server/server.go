package server

import (
	"authGo/config"
	"authGo/handler"
	"authGo/middleware"
	"log"
	"net/http"
)

func Start() {
	cfg := config.NewConfig()

	authHandler := handler.NewAuthHandler(cfg)
	registerHandler := handler.NewRegisterHandler(cfg)
	refreshHandler := handler.NewRefreshHandler(cfg)
	userHandler := handler.NewUserHandler(cfg)

	http.HandleFunc("/login", authHandler.Login)
	http.HandleFunc("/register", registerHandler.Register)
	http.Handle("/refresh", middleware.ValidRefreshToken(http.HandlerFunc(refreshHandler.GetNewAccessToken)))
	http.Handle("/profile", middleware.ValidAccessToken(http.HandlerFunc(userHandler.GetProfile)))

	log.Fatal(http.ListenAndServe(cfg.Port, nil))
}
