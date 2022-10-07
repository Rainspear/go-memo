package main

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

func closeMongoClient(client *mongo.Client) {
	if err := client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}

func getEnvVariable(key string) string {
	err := godotenv.Load(".env")
	handlePanicError(err)
	return os.Getenv(key)
}
