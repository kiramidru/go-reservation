package database

import (
	"context"
	"log"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() *mongo.Client {
	url, err := godotenv.Load("MONGODB_URI")

	if err != nil {
		log.Fatal(err)
	}
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(url).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		log.Fatal(err)
	}

	return client
}
