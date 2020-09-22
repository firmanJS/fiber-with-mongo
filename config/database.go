package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
	"github.com/firmanJS/fiber-with-mongo/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDb() {
	MONGO_HOST := config.Config(MONGO_HOST)

	opt := options.Credential{
		Username:   config.Config(MONGO_USERNAME),
		Password:   config.Config(MONGO_PASSWORD),
		AuthSource: config.Config(MONGO_SOURCE),
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("mongodb://%v:27017", MONGO_HOST)).SetAuth(opt))

	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database(config.Config(MONGO_DB_NAME))

	defer cancel()
}
