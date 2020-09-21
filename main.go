package main

import (
	"log"

	"github.com/firmanJS/fiber-with-mongo/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() { // entry point to our program

	app := fiber.New() // call the New() method - used to instantiate a new Fiber App
	app.Use(cors.New())
	app.Use(logger.New())
	router.SetupRoutes(app)
	log.Fatal(app.Listen(":2000"))
}
