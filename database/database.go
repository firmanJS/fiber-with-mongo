package database

import (
	"context"
	"time"

	"github.com/firmanJS/fiber-with-mongo/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var mg MongoInstance

func ConnectDb() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.Config("MONGO_HOST")))

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	db := client.Database(config.Config("MONGO_DB_NAME"))

	if err != nil {
		return err
	}

	mg = MongoInstance{
		Client: client,
		Db:     db,
	}

	return nil
}
