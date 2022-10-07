package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Level string

type Status string

const (
	LevelEssential Level = "essential"
	LevelImportant Level = "important"
	LevelSemi      Level = "semi-important"
	LevelLess      Level = "less-important"
	LevelMinor     Level = "minor"
)

const (
	StatusSuccess Status = "success"
	StatusUntouch Status = "untouch"
	StatusFailure Status = "failure"
	StatusSkipped Status = "skipped"
)

type Repetition struct {
	Time   time.Time `json:"time" bson:"time"`
	Level  Level     `json:"level" bson:"level"`
	Status Status    `json:"status" bson:"status"`
}

type Topic struct {
	Id          interface{}  `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string       `json:"title" bson:"title"`
	Repetition  []Repetition `json:"repetition" bson:"repetition"`
	Description string       `json:"description" bson:"description"`
	CreatedDate time.Time    `json:"created_date" bson:"created_date"`
	LastUpdate  time.Time    `json:"last_update" bson:"last_update"`
}

func getTopics(w http.ResponseWriter, req *http.Request) {
	coll := client.Database(database).Collection("topics")
	filter := bson.D{}
	cursor, err := coll.Find(req.Context(), filter)
	if handleResponseError(err, w, http.StatusInternalServerError) {
		return
	}
	var topics []Topic
	err = cursor.All(req.Context(), &topics)
	if handleResponseError(err, w, http.StatusInternalServerError) {
		return
	}
	handleResponseSuccess(topics, w, http.StatusOK)
}

func getTopic(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if handleResponseError(err, w, http.StatusBadRequest) {
		return
	}
	coll := client.Database(database).Collection("topics")
	filter := bson.D{{Key: "_id", Value: id}}
	var topic Topic
	err = coll.FindOne(req.Context(), filter).Decode(&topic)
	if handleResponseError(err, w, http.StatusInternalServerError) {
		return
	}
	handleResponseSuccess(topic, w, http.StatusOK)
}

func createTopic(w http.ResponseWriter, req *http.Request) {
	var data Topic
	err := json.NewDecoder(req.Body).Decode(&data)
	if handleResponseError(err, w, http.StatusBadRequest) {
		return
	}
	coll := client.Database(database).Collection("topics")
	result, err := coll.InsertOne(req.Context(), &data)
	if handleResponseError(err, w, http.StatusInternalServerError) {
		return
	}
	id := result.InsertedID.(primitive.ObjectID)
	handleResponseSuccess(CreatedResponse{id.Hex()}, w, http.StatusCreated)
}

func updateTopic(w http.ResponseWriter, req *http.Request) {
	var data Topic
	err := json.NewDecoder(req.Body).Decode(&data)
	if handleResponseError(err, w, http.StatusBadRequest) {
		return
	}
	coll := client.Database(database).Collection("topics")
	params := mux.Vars(req)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if handleResponseError(err, w, http.StatusBadRequest) {
		return
	}
	update := bson.D{{Key: "$set", Value: &data}}
	filter := bson.D{{Key: "_id", Value: id}}
	result, err := coll.UpdateOne(req.Context(), filter, update)
	if handleResponseError(err, w, http.StatusBadRequest) {
		return
	}
	handleResponseSuccess(result, w, http.StatusOK)
}

func deleteTopic(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if handleResponseError(err, w, http.StatusBadRequest) {
		return
	}
	coll := client.Database(database).Collection("topics")
	filter := bson.D{{Key: "_id", Value: id}}
	result, err := coll.DeleteOne(req.Context(), filter)
	if handleResponseError(err, w, http.StatusInternalServerError) {
		return
	}
	handleResponseSuccess(result, w, http.StatusOK)
}
