package controller

import (
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
)

type controller struct {
	DB *redis.Client
}

func NewController(db *redis.Client) *controller {
	return &controller{
		DB: db,
	}
}

//func InitDatabase()redis.db

func (h *controller) Register(c *fiber.Ctx) error {

	//take the url from a form value

	//generate a random unique code for url

	//store the random code as key and original url as value

	//append the random generated code along with the base url

	//return the new url

	return nil
}

func (h *controller) AccessTheURL(c *fiber.Ctx) error {

	//get unique id from url
	result := c.Params("short")
	if result == "" {
		return fiber.ErrBadRequest
	}

	//find URL corresponding to the unique code

	//redirect to the original url

	return nil
}
