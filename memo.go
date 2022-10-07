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
	coll := client.Database(database).Collection(MEMO_COLLECTION)
	filter := bson.D{}
	cursor, err := coll.Find(context.TODO(), filter)
	if handleResponseError(err, w, http.StatusInternalServerError) {
		return
	}
	memos := []Memo{}
	err = cursor.All(context.TODO(), &memos)
	if handleResponseError(err, w, http.StatusInternalServerError) {
		return
	}
	handleResponseSuccess(&memos, w, http.StatusOK)
}

func getMemo(w http.ResponseWriter, req *http.Request) {
	coll := client.Database(database).Collection(MEMO_COLLECTION)
	params := mux.Vars(req)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if handleResponseError(err, w, http.StatusBadRequest) {
		return
	}
	filter := bson.D{{Key: "_id", Value: id}}
	var memo Memo
	err = coll.FindOne(context.TODO(), filter).Decode(&memo)
	if handleResponseError(err, w, http.StatusNotFound) {
		return
	}
	handleResponseSuccess(&memo, w, http.StatusOK)
}

func createMemo(w http.ResponseWriter, req *http.Request) {
	coll := client.Database(database).Collection(MEMO_COLLECTION)
	t := time.Now()
	var data Memo
	err := json.NewDecoder(req.Body).Decode(&data)
	if handleResponseError(err, w, http.StatusBadRequest) {
		return
	}
	data.LastUpdate = t
	data.CreatedDate = t
	result, err := coll.InsertOne(context.TODO(), &data)
	if handleResponseError(err, w, http.StatusInternalServerError) {
		return
	}
	id := result.InsertedID.(primitive.ObjectID)
	handleResponseSuccess(CreatedResponse{id.Hex()}, w, http.StatusCreated)
}

func updateMemo(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if handleResponseError(err, w, http.StatusBadRequest) {
		return
	}
	filter := bson.D{{Key: "_id", Value: id}}
	var data Memo
	err = json.NewDecoder(req.Body).Decode(&data)
	if handleResponseError(err, w, http.StatusBadRequest) {
		return
	}
	data.LastUpdate = time.Now()
	update := bson.D{{Key: "$set", Value: &data}}
	coll := client.Database(database).Collection(MEMO_COLLECTION)
	result, err := coll.UpdateOne(context.TODO(), filter, update)
	if handleResponseError(err, w, http.StatusBadRequest) {
		return
	}
	handleResponseSuccess(result, w, http.StatusOK)
}

func deleteMemo(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if handleResponseError(err, w, http.StatusBadRequest) {
		return
	}
	filter := bson.D{{Key: "_id", Value: id}}
	coll := client.Database(database).Collection(MEMO_COLLECTION)
	result, err := coll.DeleteOne(context.TODO(), filter)
	if handleResponseError(err, w, http.StatusInternalServerError) {
		return
	}
	handleResponseSuccess(result, w, http.StatusOK)
}
