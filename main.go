package main

import (
	"log"

	"github.com/firmanJS/fiber-with-mongo/database"
	"github.com/firmanJS/fiber-with-mongo/helpers"
	"github.com/firmanJS/fiber-with-mongo/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() { // entry point to our program

	if err := database.ConnectDb(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New() // call the New() method - used to instantiate a new Fiber App
	app.Use(cors.New())
	app.Use(logger.New())
	router.SetupRoutes(app)

	app.Use(func(c *fiber.Ctx) error {
		return helpers.ResponseMsg(c, 404, "NotFound", nil)
	})

	log.Fatal(app.Listen(":1000"))
}
