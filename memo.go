package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Memo struct {
	Id          interface{} `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string      `json:"title"`
	Author      string      `json:"author"`
	Content     string      `json:"content"`
	CreatedDate time.Time   `json:"createdDate"`
	LastUpdate  time.Time   `json:"lastUpdate"`
}

func getMemos(w http.ResponseWriter, req *http.Request) {
	coll := client.Database(database).Collection("memos")
	filter := bson.D{}
	cursor, err := coll.Find(context.TODO(), filter)
	handlePanicError(err)
	var results []bson.M
	err = cursor.All(context.TODO(), &results)
	handlePanicError(err)
	json.NewEncoder(w).Encode(results)
}

func getMemo(w http.ResponseWriter, req *http.Request) {

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
	result, err := coll.InsertOne(context.TODO(), doc)
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
