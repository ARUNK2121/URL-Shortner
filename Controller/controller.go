package controller

import (
	"fmt"

	"github.com/ARUNK2121/url_shortner/config"
	"github.com/ARUNK2121/url_shortner/models"
	"github.com/asaskevich/govalidator"
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	uuid "github.com/satori/go.uuid"
)

type controller struct {
	DB     *redis.Client
	Config config.Config
}

func NewController(db *redis.Client, config config.Config) *controller {
	return &controller{
		DB:     db,
		Config: config,
	}
}

func (h *controller) Register(c *fiber.Ctx) error {

	url := c.FormValue("url", "invalid")
	if url == "invalid" {
		return nil
	}

	isURL := govalidator.IsURL(url)
	if !isURL {
		response := models.ShortResponse{Error: "try again with a valid url"}
		return c.Status(fiber.ErrBadRequest.Code).JSON(response)
	}

	var Unique string
	var Error error
	for {
		uid := uuid.NewV4()
		Unique = uid.String()
		_, Error = h.DB.Get(Unique).Result()
		if Error == redis.Nil {
			break
		}

		if Error != nil {
			panic("redis connection error")
		}
	}

	if err := h.DB.Set(Unique, url, 0).Err(); err != nil {
		return c.Status(500).SendString("error in storing to redis")
	}

	fmt.Println("db:", h.DB)

	baseURL := h.Config.BaseURL
	response := models.ShortResponse{}
	response.URL = baseURL + "/" + Unique

	return c.JSON(response)
}

func (h *controller) AccessTheURL(c *fiber.Ctx) error {

	short := c.Params("short")
	if short == "" {
		return fiber.ErrBadRequest
	}

	val, err := h.DB.Get(short).Result()
	if err != nil {
		panic(err)
	}

	c.Redirect(val)

	return nil
}
