package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Level string

const (
	LevelImportant Level = "important"
	LevelInfo      Level = "info"
	LevelWarn      Level = "warn"
	LevelError     Level = "error"
)

type Repetition struct {
	Time   time.Time `json:"time" bson:"time"`
	Level  Level     `json:"level" bson:"level"`
	Status string    `json:"status" bson:"status"`
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
	fmt.Printf("Body %+v\n", req.Body)
	var data Subject
	err := json.NewDecoder(req.Body).Decode(&data)
	handlePanicError(err)
	fmt.Printf("after json decode %v+", data)
}

func updateSubject(w http.ResponseWriter, req *http.Request) {
}

func deleteSubject(w http.ResponseWriter, req *http.Request) {
}
