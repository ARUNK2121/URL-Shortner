package config

import (
	"errors"
	"os"
	"strconv"
)

type Config struct {
	Port     string
	Password string
	DB       int
}

func Credentials() (Config, error) {
	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		return Config{}, errors.New("error loading password from environment")
	}

	port := os.Getenv("DB_PORT")
	if port == "" {
		return Config{}, errors.New("error loading port from environment")
	}

	db, err := strconv.Atoi(os.Getenv("DB_NAME"))
	if err != nil {
		return Config{}, errors.New("error loading port from environment")
	}

	return Config{Password: password, Port: port, DB: db}, nil
}
