package handler

import (
	"context"
	"log"

	"github.com/firmanJS/fiber-with-mongo/database"
	"github.com/firmanJS/fiber-with-mongo/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func Index(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Welcome",
	})
}

func GetUsers(c *fiber.Ctx) error {
	var ctx = context.Background()
	db, err := database.Connect()

	if err != nil {
		log.Fatal(err.Error())
	}

	csr, err := db.Collection("users").Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err.Error())
	}
	defer csr.Close(ctx)

	result := make([]models.Users, 0)
	for csr.Next(ctx) {
		var row models.Users
		err := csr.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}

		result = append(result, row)
	}
	return c.Status(200).JSON(fiber.Map{"success": true, "users": result})

}
