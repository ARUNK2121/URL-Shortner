package database

import (
	"errors"

	"github.com/ARUNK2121/url_shortner/config"
	"github.com/go-redis/redis"
)

func InitDatabase(c config.Config) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     c.Port,
		Password: c.Password, // no password set ("")
		DB:       c.DB,       // use default DB (0)
	})
	if client == nil {
		return &redis.Client{}, errors.New("error in creating redis client")
	}

	return client, nil
}
