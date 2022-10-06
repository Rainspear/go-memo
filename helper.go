package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

type Enum interface {
	IsValid() bool
}

func handlePanicError(err error) {
	if err != nil {
		panic(err)
	}
}

func handleResponseError(err error, w http.ResponseWriter, statusCode int) {
	if err != nil {
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(ErrorResponse{err.Error(), statusCode})
		return
	}
}

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
