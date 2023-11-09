package controller

import (
	"fmt"

	"github.com/ARUNK2121/url_shortner/config"
	"github.com/ARUNK2121/url_shortner/models"
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

	//take the url from a form value
	url := c.FormValue("url", "invalid")
	if url == "invalid" {
		return nil
	}

	fmt.Println("url:", url)

	//generate a random unique code for url
	var Unique string
	var Error error
	for {
		uid := uuid.NewV4()
		Unique = uid.String()
		fmt.Println("unique:", Unique)
		_, Error = h.DB.Get(Unique).Result()
		if Error == redis.Nil {
			break
		}
	}

	//store the unique id as key and the original url as value in redis
	if err := h.DB.Set(Unique, url, 0).Err(); err != nil {
		return c.Status(500).SendString("error in storing to redis")
	}

	fmt.Println("db:", h.DB)

	//append the random generated code along with the base url
	baseURL := h.Config.BaseURL
	response := models.ShortResponse{}
	response.URL = baseURL + "/" + Unique

	//return the new url

	return c.JSON(response)
}

func (h *controller) AccessTheURL(c *fiber.Ctx) error {

	//get unique id from url
	result := c.Params("short")
	if result == "" {
		return fiber.ErrBadRequest
	}

	//find URL corresponding to the unique code
	val, err := h.DB.Get(result).Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("val", val)

	//redirect to the original url

	return nil
}
