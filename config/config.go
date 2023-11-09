package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Port     string
	Password string
	DB       int
	BaseURL  string
}

func Credentials() (Config, error) {
	password := os.Getenv("DB_PASSWORD")
	// if password == "" {
	// 	return Config{}, errors.New("error loading password from environment")
	// }

	port := os.Getenv("DB_PORT")
	if port == "" {
		return Config{}, errors.New("error loading port from environment")
	}

	db, err := strconv.Atoi(os.Getenv("DB_NAME"))
	if err != nil {
		return Config{}, errors.New("error loading port from environment")
	}

	baseurl := os.Getenv("BASE_URL")
	if port == "" {
		return Config{}, errors.New("error loading base url from environment")
	}
	fmt.Println("base url is ", baseurl)

	return Config{Password: password, Port: port, DB: db, BaseURL: baseurl}, nil
}
