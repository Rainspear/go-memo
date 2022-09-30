package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Memo struct {
	Id          interface{} `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string      `json:"title" bson:"title"`
	Author      string      `json:"author" bson:"author"`
	Content     string      `json:"content" bson:"content"`
	CreatedDate time.Time   `json:"created_date" bson:"created_date"`
	LastUpdate  time.Time   `json:"last_update" bson:"last_update"`
}

func getMemos(w http.ResponseWriter, req *http.Request) {
	coll := client.Database(database).Collection("memos")
	filter := bson.D{}
	cursor, err := coll.Find(context.TODO(), filter)
	handlePanicError(err)
	// var results []bson.M
	var memos []Memo
	err = cursor.All(context.TODO(), &memos)
	handlePanicError(err)
	json.NewEncoder(w).Encode(memos)
}

func getMemo(w http.ResponseWriter, req *http.Request) {
	coll := client.Database(database).Collection("memos")
	params := mux.Vars(req)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResponse{err.Error(), http.StatusBadRequest})
		return
	}
	filter := bson.D{{Key: "_id", Value: id}}
	var memo Memo
	err = coll.FindOne(context.TODO(), filter).Decode(&memo)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResponse{err.Error(), http.StatusNotFound})
		return
	}
	json.NewEncoder(w).Encode(memo)

}

func createMemo(w http.ResponseWriter, req *http.Request) {
	coll := client.Database(database).Collection("memos")
	title := req.FormValue("title")
	author := req.FormValue("author")
	content := req.FormValue("content")
	t := time.Now().UTC()
	doc := bson.D{
		{Key: "title", Value: title},
		{Key: "author", Value: author},
		{Key: "content", Value: content},
		{Key: "createdDate", Value: t},
		{Key: "lastUpdate", Value: t}}
	result, err := coll.InsertOne(context.TODO(), &doc)
	handlePanicError(err)
	id := result.InsertedID.(primitive.ObjectID)
	json.NewEncoder(w).Encode(struct {
		Id string `json:"id"`
	}{
		id.Hex()})
}

func updateMemo(w http.ResponseWriter, req *http.Request) {

}

func deleteMemo(w http.ResponseWriter, req *http.Request) {

}
