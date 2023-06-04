package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	Port                        string
	AccessTokenSecret           string
	AccessTokenLifetimeMinutes  int
	RefreshTokenSecret          string
	RefreshTokenLifetimeMinutes int
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	AccessTokenLifetimeMinutes, _ := strconv.Atoi(os.Getenv("ACCESS_TOKEN_LIFETIME_MINUTES"))
	RefreshTokenLifetimeMinutes, _ := strconv.Atoi(os.Getenv("REFRESH_TOKEN_LIFETIME_MINUTES"))

	return &Config{
		Port:                        os.Getenv("PORT"),
		AccessTokenSecret:           os.Getenv("ACCESS_TOKEN_SECRET"),
		AccessTokenLifetimeMinutes:  AccessTokenLifetimeMinutes,
		RefreshTokenSecret:          os.Getenv("REFRESH_TOKEN_SECRET"),
		RefreshTokenLifetimeMinutes: RefreshTokenLifetimeMinutes,
	}
}
