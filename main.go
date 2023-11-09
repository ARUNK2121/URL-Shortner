package main

import (
	"fmt"
	"log"

	controller "github.com/ARUNK2121/url_shortner/Controller"
	"github.com/ARUNK2121/url_shortner/config"
	"github.com/ARUNK2121/url_shortner/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func main() {

	//load godotenv and load env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading the env file")
	}

	//load credentials into config struct
	config, err := config.Credentials()
	if err != nil {
		fmt.Println("error is:", err.Error())
		log.Fatal("error getting the credentials from env file")

	}

	//create new fiber engine
	app := fiber.New()

	//initialize database
	db, err := database.InitDatabase(config)
	if err != nil {
		log.Fatal("could not create database connection")
	}

	//get controller for handling routes
	controller := controller.NewController(db, config)

	//set the routes and start the server
	app.Use(recover.New())
	app.Post("/", controller.Register)
	app.Get("/:UID", controller.AccessTheURL)
	log.Fatal(app.Listen(":8000"))

}
