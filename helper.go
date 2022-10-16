package main

import (
	"context"
	"encoding/json"
	"fmt"
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

func structToMap(obj interface{}, newMap map[string]interface{}) bool {
	bs, err := json.Marshal(obj) // Convert to a json string
	if err != nil {
		fmt.Println("Can not parse from struct to json: ", err)
		return false
	}
	err = json.Unmarshal(bs, &newMap) // Convert to a map
	if err != nil {
		fmt.Println("Can not parse from json to map: ", err)
		return false
	}
	return true
}
