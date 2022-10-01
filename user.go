package main

import "net/http"

type User struct {
	Id       interface{} `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string      `json:"name" bson:"name"`
	Avatar   string      `json:"avatar" bson:"avatar"`
	Password string      `json:"password" bson:"password"`
}

func getUsers(w http.ResponseWriter, req *http.Request) {

}

func getCurrentUser(w http.ResponseWriter, req *http.Request) {

}

func createUser(w http.ResponseWriter, req *http.Request) {

}

func updateUser(w http.ResponseWriter, req *http.Request) {

}

func deleteUser(w http.ResponseWriter, req *http.Request) {

}

func signin(w http.ResponseWriter, req *http.Request) {

}

func signup(w http.ResponseWriter, req *http.Request) {

}

func signout(w http.ResponseWriter, req *http.Request) {

}
