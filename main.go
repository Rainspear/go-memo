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
	// if err := godotenv.Load(); err != nil {
	// 	log.Println("No .env file found")
	// }
	// uri := os.Getenv("URL_MONGODB")
	// if uri == "" {
	// 	log.Fatal("You must set your 'URL_MONGODB' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	// }
	// client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	// if err != nil {
	// 	panic(err)
	// }
	// defer func() {
	// 	if err := client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()
	// coll := client.Database("sample_mflix").Collection("movies")
	// title := "Back to the Future"
	// var result bson.M
	// err = coll.FindOne(context.TODO(), bson.D{{Key: "title", Value: title}}).Decode(&result)
	// if err == mongo.ErrNoDocuments {
	// 	fmt.Printf("No document was found with the title %s\n", title)
	// 	return
	// }
	// if err != nil {
	// 	panic(err)
	// }
	// jsonData, err := json.MarshalIndent(result, "", "    ")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%s\n", jsonData)

	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", postMovies).Methods("POST")
	r.HandleFunc("/movies/{id}", putMovies).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovies).Methods("DELETE")
	fmt.Printf("Starting server at port 8089\n")
	log.Fatal((http.ListenAndServe(":8089", r)))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func getEnvVariable(key string) string {
	err := godotenv.Load(".env")
	check(err)
	return os.Getenv(key)
}

func getMovie(w http.ResponseWriter, req *http.Request) {
	uri := getEnvVariable("URL_MONGODB")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	check(err)
	fmt.Printf("%+v\n", client)
	defer func() { // if main was exit this function was called to disconnect db
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	coll := client.Database("sample_mflix").Collection("movies")
	title := "Back to the Future"

	var result bson.M
	err = coll.FindOne(context.TODO(), bson.D{{Key: "title", Value: title}}).Decode(&result)
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
	fmt.Fprintf(w, "%s\n", string(jsonData))
}

func getMovies(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "%+v\n", client)
}

func postMovies(w http.ResponseWriter, req *http.Request) {

}

func putMovies(w http.ResponseWriter, req *http.Request) {

}

func deleteMovies(w http.ResponseWriter, req *http.Request) {

}
