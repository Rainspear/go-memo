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
	Id          interface{} `json:"_id,omitempty" bson:"_id,omitempty"`
	SubjectId   string      `json:"subject_id" bson:"subject_id"`
	Author      string      `json:"author" bson:"author"`
	Content     string      `json:"content" bson:"content"`
	Question    string      `json:"question" bson:"question"`
	CreatedDate time.Time   `json:"created_date" bson:"created_date"`
	LastUpdate  time.Time   `json:"last_update" bson:"last_update"`
}

func getMemos(w http.ResponseWriter, req *http.Request) {
	coll := client.Database(database).Collection("memos")
	filter := bson.D{}
	cursor, err := coll.Find(context.TODO(), filter)
	handlePanicError(err)
	// var results []bson.M
	memos := []Memo{}
	err = cursor.All(context.TODO(), &memos)
	handlePanicError(err)
	json.NewEncoder(w).Encode(memos)
}

func getMemo(w http.ResponseWriter, req *http.Request) {
	coll := client.Database(database).Collection("memos")
	params := mux.Vars(req)
	id, err := primitive.ObjectIDFromHex(params["id"])
	handleResponseError(err, w, http.StatusBadRequest)
	filter := bson.D{{Key: "_id", Value: id}}
	var memo Memo
	err = coll.FindOne(context.TODO(), filter).Decode(&memo)
	handleResponseError(err, w, http.StatusNotFound)
	json.NewEncoder(w).Encode(memo)
}

func createMemo(w http.ResponseWriter, req *http.Request) {
	coll := client.Database(database).Collection("memos")
	subject_id := req.FormValue("subject_id")
	author := req.FormValue("author")
	content := req.FormValue("content")
	question := req.FormValue("question")
	t := time.Now()
	doc := Memo{SubjectId: subject_id, Author: author, Content: content, Question: question, CreatedDate: t, LastUpdate: t}
	result, err := coll.InsertOne(context.TODO(), &doc)
	handlePanicError(err)
	id := result.InsertedID.(primitive.ObjectID)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(struct {
		Id string `json:"id"`
	}{
		id.Hex(),
	})
}

func updateMemo(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, err := primitive.ObjectIDFromHex(params["id"])
	handleResponseError(err, w, http.StatusBadRequest)
	filter := bson.D{{Key: "_id", Value: id}}
	subject_id := req.FormValue("subject_id")
	author := req.FormValue("author")
	content := req.FormValue("content")
	question := req.FormValue("question")
	t := time.Now()
	doc := Memo{SubjectId: subject_id, Author: author, Content: content, Question: question, LastUpdate: t}
	update := bson.D{{Key: "$set", Value: &doc}}
	coll := client.Database(database).Collection("memos")
	result, err := coll.UpdateOne(context.TODO(), filter, update)
	handleResponseError(err, w, http.StatusBadRequest)
	json.NewEncoder(w).Encode(result)
}

func deleteMemo(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, err := primitive.ObjectIDFromHex(params["id"])
	handleResponseError(err, w, http.StatusBadRequest)
	filter := bson.D{{Key: "_id", Value: id}}
	coll := client.Database(database).Collection("memos")
	result, err := coll.DeleteOne(context.TODO(), filter)
	handleResponseError(err, w, http.StatusBadRequest)
	json.NewEncoder(w).Encode(result)
}
