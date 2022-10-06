package main

import (
	"net/http"
	"time"
)

type Repetition struct {
	TimeRepetition []time.Time `json:"time_repetition" bson:"time_repetition"`
	Level          string      `json:"level" bson:"level"`
}

type Subject struct {
	Id         interface{}  `json:"_id,omitempty" bson:"_id,omitempty"`
	Title      string       `json:"title" bson:"title"`
	Repetition []Repetition `json:"repetition" bson:"repetition"`
}

func getSubjects(w http.ResponseWriter, req *http.Request) {
}

func getSubject(w http.ResponseWriter, req *http.Request) {
}

func createSubject(w http.ResponseWriter, req *http.Request) {
}

func updateSubject(w http.ResponseWriter, req *http.Request) {
}

func deleteSubject(w http.ResponseWriter, req *http.Request) {
}
