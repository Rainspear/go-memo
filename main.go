package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// var mongodbURL string = "mongodb+srv://memoAppUser:m22password@cluster0.tirlw.mongodb.net/?retryWrites=true&w=majority"

var client *mongo.Client

func main() {
	uri := getEnvVariable("URL_MONGODB")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	check(err)
	defer func() { // if main was exit this function was called to disconnect db
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", getMovies).Methods("POST")
	r.HandleFunc("/movies/{id}", getMovies).Methods("PUT")
	r.HandleFunc("/movies/{id}", getMovies).Methods("DELETE")
	fmt.Printf("Starting server at port 8089\n")
	log.Fatal((http.ListenAndServe(":8089", r)))
}

func check(err error) {
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}
}

func getEnvVariable(key string) string {
	err := godotenv.Load(".env")
	check(err)
	return os.Getenv(key)
}

func getMovie(w http.ResponseWriter, req *http.Request) {
	coll := client.Database("sample_mflix").Collection("movies")
	title := "Back to the Future"

	var result bson.M
	err := coll.FindOne(context.TODO(), bson.D{{"title", title}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Fprintf(w, "No document was found with the title %s\n", title)
		return
	}
	if err != nil {
		panic(err)
	}
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
}

func getMovies(w http.ResponseWriter, req *http.Request) {

}
