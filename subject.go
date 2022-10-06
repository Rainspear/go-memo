package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
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
	case StatusSuccess, StatusFailure, StatusSkipped:
		return true
	default:
		return false
	}
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
