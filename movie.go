package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Movie struct {
	Id        string   `json:"id"`
	Title     string   `json:"title"`
	Genres    []string `json:"genre"`
	Cast      []string `json:"cast"`
	Countries []string `json:"countries"`
}

func getMovie(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	fmt.Printf("getMovie params %+v \n", params)
	// fmt.Printf("getMovie client %+v \n", client)
}

func getMovies(w http.ResponseWriter, req *http.Request) {
	// start finding
	coll := client.Database("sample_mflix").Collection("movies")
	filter := bson.D{}
	// filter := bson.D{{Key: "runtime", Value: 1}}

	cursor, err := coll.Find(context.TODO(), filter, options.Find().SetLimit(20))
	handlePanicError(err)
	// var results []bson.M
	var results []Movie
	err = cursor.All(context.TODO(), &results)
	handlePanicError(err)
	// end finding
	// response to client request
	json.NewEncoder(w).Encode(results)
}

func createMovie(w http.ResponseWriter, req *http.Request) {

}

func updateMovie(w http.ResponseWriter, req *http.Request) {

}

func deleteMovie(w http.ResponseWriter, req *http.Request) {

}
