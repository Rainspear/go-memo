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

type Subject struct {
	Id         interface{}  `json:"_id,omitempty" bson:"_id,omitempty"`
	Title      string       `json:"title" bson:"title"`
	Repetition []Repetition `json:"repetition" bson:"repetition"`
}

func (l Level) IsValid() bool {
	switch l {
	case LevelEssential, LevelImportant, LevelSemi, LevelLess, LevelMinor:
		return true
	default:
		return false
	}
}

func (s Status) IsValid() bool {
	switch s {
	case StatusSuccess, StatusFailure, StatusSkipped, StatusUntouch:
		return true
	default:
		return false
	}
}

func getSubjects(w http.ResponseWriter, req *http.Request) {
	coll := client.Database(database).Collection("subjects")
	filter := bson.D{}
	cursor, err := coll.Find(req.Context(), filter)
	handleResponseError(err, w, http.StatusInternalServerError)
	var results []Subject
	err = cursor.All(req.Context(), &results)
	handleResponseError(err, w, http.StatusInternalServerError)
	json.NewEncoder(w).Encode(results)
}

func getSubject(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, err := primitive.ObjectIDFromHex(params["id"])
	handleResponseError(err, w, http.StatusBadRequest)
	coll := client.Database(database).Collection("subjects")
	filter := bson.D{{Key: "_id", Value: id}}
	coll.FindOne(req.Context(), filter)
}

func createSubject(w http.ResponseWriter, req *http.Request) {
	var data Subject
	err := json.NewDecoder(req.Body).Decode(&data)
	handleResponseError(err, w, http.StatusBadRequest)
	coll := client.Database(database).Collection("subjects")
	result, err := coll.InsertOne(req.Context(), &data)
	handleResponseError(err, w, http.StatusInternalServerError)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}

func updateSubject(w http.ResponseWriter, req *http.Request) {
}

func deleteSubject(w http.ResponseWriter, req *http.Request) {
}
