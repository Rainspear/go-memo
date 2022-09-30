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
var uri string
var movies []Movie

func init() {
	var err error
	// if err := godotenv.Load(); err != nil {
	// 	log.Println("No .env file found")
	// }
	// uri = os.Getenv("URL_MONGODB")
	uri = getEnvVariable("URL_MONGODB")
	if uri == "" {
		log.Fatal("You must set your 'URL_MONGODB' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	handlePanicError(err)
	// defer func() {
	// 	if err := client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()
	// coll := client.Database("sample_mflix").Collection("movies")
	// title := "Back to the Future"
	// var result bson.M
	// err = coll.FindOne(context.TODO(), bson.D{{"title", title}}).Decode(&result)
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
}

func test() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	var uri string
	if uri = os.Getenv("URL_MONGODB"); uri == "" {
		log.Fatal("You must set your 'URL_MONGODB' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// begin find
	coll := client.Database("sample_training").Collection("zips")
	filter := bson.D{{"pop", bson.D{{"$lte", 500}}}}

	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	// end find

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	for _, result := range results {
		output, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", output)
	}
}

func main() {
	// fmt.Printf("client %+v", client)
	// test()
	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovies).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovies).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovies).Methods("DELETE")
	fmt.Printf("Starting server at port 8089\n")
	log.Fatal((http.ListenAndServe(":8089", r)))
}

func handlePanicError(err error) {
	if err != nil {
		panic(err)
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

func getMovie(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	fmt.Printf("getMovie params %+v \n", params)
	// fmt.Printf("getMovie client %+v \n", client)
}

func getMovies(w http.ResponseWriter, req *http.Request) {
	defer closeMongoClient(client)
	w.Header().Set("Content-Type", "application/json")
	// start finding
	coll := client.Database("sample_mflix").Collection("movies")
	filter := bson.D{{Key: "runtime", Value: 1}}

	cursor, err := coll.Find(context.TODO(), filter)
	handlePanicError(err)
	var results []bson.M
	err = cursor.All(context.TODO(), &results)
	handlePanicError(err)
	// end finding
	// for _, result := range results {
	// 	output, err := json.MarshalIndent(result, "", "  ")
	// 	handlePanicError(err)
	// 	fmt.Printf("%s\n", output)
	// }
	// response to client request
	json.NewEncoder(w).Encode(results)

}

func createMovies(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

func updateMovies(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

func deleteMovies(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}
