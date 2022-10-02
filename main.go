package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var uri string
var database string

func init() {
	var err error
	database = getEnvVariable("MONGODB_DATABASE")
	uri = getEnvVariable("MONGODB_URI")
	if database == "" {
		log.Fatal("You must set your 'MONGODB_DATABASE' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	handlePanicError(err)
}

func main() {
	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r)) // cors middleware
	// movies
	r.HandleFunc("/movies", addConfigMiddleware(getMovies)).Methods("GET")
	r.HandleFunc("/movies/{id}", addConfigMiddleware(getMovie)).Methods("GET")
	r.HandleFunc("/movies", addConfigMiddleware(createMovie)).Methods("POST")
	r.HandleFunc("/movies/{id}", addConfigMiddleware(updateMovie)).Methods("PUT")
	r.HandleFunc("/movies/{id}", addConfigMiddleware(deleteMovie)).Methods("DELETE")
	// memos
	r.HandleFunc("/memos", addConfigMiddleware(getMemos)).Methods("GET")
	r.HandleFunc("/memos/{id}", addConfigMiddleware(getMemo)).Methods("GET")
	r.HandleFunc("/memos", addConfigMiddleware(createMemo)).Methods("POST")
	r.HandleFunc("/memos/{id}", addConfigMiddleware(updateMemo)).Methods("PUT")
	r.HandleFunc("/memos/{id}", addConfigMiddleware(deleteMemo)).Methods("DELETE")
	// user
	r.HandleFunc("/user", addConfigMiddleware(getUsers)).Methods("GET")
	r.HandleFunc("/user/{id}", addConfigMiddleware(getCurrentUser)).Methods("GET")
	r.HandleFunc("/user/{id}", addConfigMiddleware(updateUser)).Methods("PUT")
	r.HandleFunc("/user/{id}", addConfigMiddleware(deleteUser)).Methods("DELETE")
	r.HandleFunc("/signin", addConfigMiddleware(signin)).Methods("POST")
	r.HandleFunc("/signup", addConfigMiddleware(signup)).Methods("POST")
	r.HandleFunc("/signout", addConfigMiddleware(signout)).Methods("POST")
	// default
	r.Handle("/favicon.ico", http.NotFoundHandler())
	fmt.Printf("Starting server at port 8089 \r\n")
	log.Fatal((http.ListenAndServe(":8089", r)))
	defer closeMongoClient(client)
}
