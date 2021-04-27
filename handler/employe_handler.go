package handler

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/firmanJS/fiber-with-mongo/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

func ConnectDb() error {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://fimz:fimz@localhost:27017/dbboillerplate?authSource=admin"))

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	db := client.Database("dbboillerplate")

	if err != nil {
		return err
	}

	mg = MongoInstance{
		Client: client,
		Db:     db,
	}

	return nil
}

var mg MongoInstance

func GetEmploye(c *fiber.Ctx) error {

	query := bson.D{{}}
	cursor, err := mg.Db.Collection("employees").Find(c.Context(), query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	var employees []models.Employee = make([]models.Employee, 0)

	// iterate the cursor and decode each item into an Employee
	if err := cursor.All(c.Context(), &employees); err != nil {
		return c.Status(500).SendString(err.Error())

	}
	// return employees list in JSON format
	return c.JSON(employees)
}
