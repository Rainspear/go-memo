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
	TopicId     string      `json:"topic_id" bson:"topic_id"`
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
	memos := []Memo{}
	err = cursor.All(context.TODO(), &memos)
	handlePanicError(err)
	handleResponseSuccess(&memos, w, http.StatusOK)
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
	handleResponseSuccess(&memo, w, http.StatusOK)
}

func createMemo(w http.ResponseWriter, req *http.Request) {
	coll := client.Database(database).Collection("memos")
	t := time.Now()
	var data Memo
	err := json.NewDecoder(req.Body).Decode(&data)
	handleResponseError(err, w, http.StatusBadRequest)
	data.LastUpdate = t
	data.CreatedDate = t
	result, err := coll.InsertOne(context.TODO(), &data)
	handlePanicError(err)
	id := result.InsertedID.(primitive.ObjectID)
	handleResponseSuccess(CreatedResponse{id.Hex()}, w, http.StatusCreated)
}

func updateMemo(w http.ResponseWriter, req *http.Request) {
	// parse params
	params := mux.Vars(req)
	// get id from params
	id, err := primitive.ObjectIDFromHex(params["id"])
	handleResponseError(err, w, http.StatusBadRequest)
	// create filter
	filter := bson.D{{Key: "_id", Value: id}}
	// create data from body
	var data Memo
	err = json.NewDecoder(req.Body).Decode(&data)
	handleResponseError(err, w, http.StatusBadRequest)
	data.LastUpdate = time.Now()
	// create update data with operator $set
	update := bson.D{{Key: "$set", Value: &data}}
	// find and update
	coll := client.Database(database).Collection("memos")
	result, err := coll.UpdateOne(context.TODO(), filter, update)
	handleResponseError(err, w, http.StatusBadRequest)
	// return response
	handleResponseSuccess(result, w, http.StatusOK)
}

func deleteMemo(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, err := primitive.ObjectIDFromHex(params["id"])
	handleResponseError(err, w, http.StatusBadRequest)
	filter := bson.D{{Key: "_id", Value: id}}
	coll := client.Database(database).Collection("memos")
	result, err := coll.DeleteOne(context.TODO(), filter)
	handleResponseError(err, w, http.StatusBadRequest)
	handleResponseSuccess(result, w, http.StatusOK)
}
