package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var uri string
var database string
var jwtKey []byte

func init() {
	var err error
	database = getEnvVariable("MONGODB_DATABASE")
	uri = getEnvVariable("MONGODB_URI")
	jwtKey := []byte(getEnvVariable("JWT_KEY"))
	if database == "" {
		log.Fatal("You must set your 'MONGODB_DATABASE' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	if string(jwtKey) == "" {
		log.Fatal("You must set your 'JWT_KEY' environmental variable")
	}
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	handlePanicError(err)
}

func main() {
	r := mux.NewRouter()
	// r.Use(mux.CORSMethodMiddleware(r)) // cors middleware
	// movies
	r.HandleFunc("/movies", addConfigMiddleware(getMovies)).Methods("GET")
	r.HandleFunc("/movies/{id}", addConfigMiddleware(getMovie)).Methods("GET")
	r.HandleFunc("/movies", addConfigMiddleware(createMovie)).Methods("POST")
	r.HandleFunc("/movies/{id}", addConfigMiddleware(updateMovie)).Methods("PUT")
	r.HandleFunc("/movies/{id}", addConfigMiddleware(deleteMovie)).Methods("DELETE")
	// topic
	r.HandleFunc("/topics", addConfigMiddleware(getTopics)).Methods("GET")
	r.HandleFunc("/topics/{id}", addConfigMiddleware(getTopic)).Methods("GET")
	r.HandleFunc("/topics", addConfigMiddleware(createTopic)).Methods("POST")
	r.HandleFunc("/topics/{id}", addConfigMiddleware(updateTopic)).Methods("PUT")
	r.HandleFunc("/topics/{id}", addConfigMiddleware(deleteTopic)).Methods("DELETE")
	// memos
	r.HandleFunc("/memos", addConfigMiddleware(getMemos)).Methods("GET")
	r.HandleFunc("/memos/{id}", addConfigMiddleware(getMemo)).Methods("GET")
	r.HandleFunc("/memos", addConfigMiddleware(createMemo)).Methods("POST")
	r.HandleFunc("/memos/{id}", addConfigMiddleware(updateMemo)).Methods("PUT")
	r.HandleFunc("/memos/{id}", addConfigMiddleware(deleteMemo)).Methods("DELETE")
	// user
	r.HandleFunc("/users", addConfigMiddleware(authorizeUser(getUsers))).Methods("GET")
	r.HandleFunc("/current-user", addConfigMiddleware(authorizeUser(getCurrentUser))).Methods("GET")
	r.HandleFunc("/users/{id}", addConfigMiddleware(updateUser)).Methods("PUT")
	r.HandleFunc("/users/{id}", addConfigMiddleware(deleteUser)).Methods("DELETE")
	r.HandleFunc("/signin", addConfigMiddleware(signin)).Methods("POST")
	r.HandleFunc("/signup", addConfigMiddleware(signup)).Methods("POST")
	r.HandleFunc("/signout", addConfigMiddleware(authorizeUser(signout))).Methods("POST")
	// default
	r.Handle("/favicon.ico", http.NotFoundHandler())
	fmt.Printf("Starting server at port 8089 \r\n")
	// cors config
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{getEnvVariable("ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})
	log.Fatal((http.ListenAndServe(":8089", handlers.CORS(originsOk, headersOk, methodsOk)(r))))
	defer closeMongoClient(client)
}
