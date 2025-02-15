package bootstrap

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDatabase(env *Env) mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbUser := env.DBUser
	dbPass := env.DBPass
	dbHost := env.DBHost
	dbPort := env.DBPort

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	mongodbURI := fmt.Sprintf("mongodb://%s:%s@%s:%s", dbUser, dbPass, dbHost, dbPort)

	if dbUser == "" || dbPass == "" {
		mongodbURI = fmt.Sprintf("mongodb://%s:%s", dbHost, dbPort)
	}

	opts := options.Client().ApplyURI(mongodbURI).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	return *client
}

func CloseMongoDBConnection(client mongo.Client) {
    err := client.Disconnect(context.TODO())
    if err != nil {
        log.Fatal(err)
    }
    log.Println("Connection to MongoDB closed")
}
