package main

import (
	"context"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/golang-jwt/jwt"
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
		fmt.Println("Can not parse from struct to map: ", err)
		return false
	}
	err = json.Unmarshal(bs, &newMap) // Convert to a map
	return true
}

func generateJwtTokenAndSign(data Token) string {
	var jwtData jwt.MapClaims
	structToMap(data, jwtData) // Convert to jwt.MapClaims
	jwtKey := getEnvVariable("JWT_KEY")
	if jwtKey == "" {
		log.Fatal("You must set your 'JWT_KEY' environmental variable")
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodES256, jwtData).SignedString(jwtKey)
	if err != nil {
		fmt.Println("Can not create jwt: ", err)
		return ""
	}
	return token
}

func signMessage(msg []byte) ([]byte, error) {
	var key []byte
	for i := 1; i < 65; i++ {
		key = append(key, byte(i))
	}
	h := hmac.New(sha512.New, key)
	_, err := h.Write(msg)
	if err != nil {
		fmt.Println("Error when sigining message", err)
	}
	return h.Sum(nil), nil
}

func checkSign(msg, sig []byte) (bool, error) {
	newSig, err := signMessage(msg)

	if err != nil {
		return false, fmt.Errorf("Errot when checkSign to get signature: %w", err)
	}

	same := hmac.Equal(sig, newSig)
	return same, nil
}
