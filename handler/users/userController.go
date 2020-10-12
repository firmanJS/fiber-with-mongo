package handler

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/firmanJS/fiber-with-mongo/database"
	"github.com/firmanJS/fiber-with-mongo/helpers"
	"github.com/firmanJS/fiber-with-mongo/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func Index(c *fiber.Ctx) error {
	return helpers.ResponseMsg(c, 200, true, "Api is running", nil)
}

func UpdateUsers(c *fiber.Ctx) error {
	_id := c.Params("id")
	users := new(models.Users)
	var ctx = context.Background()
	db, err := database.Connect()

	if err != nil {
		log.Fatal(err.Error())
	}
	if err := c.BodyParser(users); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": fmt.Sprintf("Invalid post body. %s", err.Error())})
	} else {
		if docID, err := primitive.ObjectIDFromHex(_id); err != nil {
			return c.Status(400).JSON(fiber.Map{"success": false})
		} else {
			q := bson.M{"_id": docID}

			u := bson.D{
				{"$set", bson.D{
					{"username", users.Username},
					{"email", users.Email},
					{"updatedAt", time.Now()},
				},
				}}

			o := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)

			db.Collection("users").FindOneAndUpdate(ctx, q, u, o).Decode(&users)
			return c.Status(200).JSON(fiber.Map{"success": true, "users": users})
		}
	}
}

func GetByIdUsers(c *fiber.Ctx) error {
	_id := c.Params("id")
	var ctx = context.Background()
	db, err := database.Connect()

	if err != nil {
		log.Fatal(err.Error())
	}
	
	if docID, err := primitive.ObjectIDFromHex(_id); err != nil {
		return c.Status(400).JSON(fiber.Map{"success": false})
	} else {
		q := bson.M{"_id": docID}
		users := models.Users{}
		result := db.Collection("users").FindOne(ctx, q)
		result.Decode(&users)

		if result.Err() != nil {
			fmt.Println(result.Err().Error())
			return c.Status(200).JSON(fiber.Map{"success": true, "users": fmt.Sprintf("No users found for give id: %s", _id)})
		} else {
			return c.Status(200).JSON(fiber.Map{"success": true, "users": users})
		}
	}

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

	return helpers.ResponseMsg(c, 200, true, "Get Data Succesfully", result)

}

func CreateUsers(c *fiber.Ctx) error {
	users := new(models.Users)
	var ctx = context.Background()
	db, err := database.Connect()

	if err != nil {
		log.Fatal(err.Error())
	}

	users.CreatedAt = time.Now()
	users.UpdatedAt = time.Now()

	if err := c.BodyParser(users); err != nil {
		return helpers.ResponseMsg(c, 400, false, err.Error(), nil)
	} else {
		if r, err := db.Collection("users").InsertOne(ctx, users); err != nil {
			return helpers.ResponseMsg(c, 500, false, err.Error(), nil)
		} else {
			fmt.Println(r)
			return helpers.ResponseMsg(c, 500, false, "Inserted", r)
		}
	}
}

func DeleteUsers(c *fiber.Ctx) error {
	_id := c.Params("id")
	var ctx = context.Background()
	db, err := database.Connect()

	if err != nil {
		log.Fatal(err.Error())
	}

	if docID, err := primitive.ObjectIDFromHex(_id); err != nil {
		return helpers.ResponseMsg(c, 400, true, "Sucess", nil)
	} else {
		q := bson.M{"_id": docID}
		r := db.Collection("users").FindOneAndDelete(ctx, q)
		if r.Err() != nil {
			return helpers.ResponseMsg(c, 400, false, "error", r)
		} else {
			return helpers.ResponseMsg(c, 200, false, "Sucess", r)
		}
	}
}
