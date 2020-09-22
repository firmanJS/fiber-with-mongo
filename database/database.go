package database

import (
	"context"

	"github.com/firmanJS/fiber-with-mongo/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() (*mongo.Database, error) {
	clientOptions := options.Client()
	clientOptions.ApplyURI(config.Config("MONGO_HOST"))
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	var ctx = context.Background()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return client.Database(config.Config("MONGO_DB_NAME")), nil
}
