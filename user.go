package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id          interface{} `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string      `json:"name" bson:"name"`
	Avatar      string      `json:"avatar,omitempty" bson:"avatar,omitempty"`
	Email       string      `json:"email" bson:"email"`
	Password    string      `json:"password" bson:"password"`
	CreatedDate time.Time   `json:"created_date" bson:"created_date"`
	LastUpdate  time.Time   `json:"last_update" bson:"last_update"`
}

func getUsers(w http.ResponseWriter, req *http.Request) {

}

func getCurrentUser(w http.ResponseWriter, req *http.Request) {

}

func updateUser(w http.ResponseWriter, req *http.Request) {

}

func deleteUser(w http.ResponseWriter, req *http.Request) {

}

func signin(w http.ResponseWriter, req *http.Request) {

}

func signup(w http.ResponseWriter, req *http.Request) {
	e := req.FormValue("email")
	p := req.FormValue("password")
	n := req.FormValue("name")
	bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	handleResponseError(err, w, http.StatusBadRequest)
	t := time.Now()
	user := User{Name: n, Email: e, Password: string(bs), CreatedDate: t, LastUpdate: t}
	coll := client.Database(database).Collection("users")
	err = coll.FindOne(context.TODO(), bson.D{{Key: "email", Value: e}}).Err()
	if err == nil { // existed user in database
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{err.Error(), http.StatusBadRequest})
		return
	}
	result, err := coll.InsertOne(context.TODO(), &user)
	handleResponseError(err, w, http.StatusInternalServerError)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
	// create jwt token
}

func signout(w http.ResponseWriter, req *http.Request) {

}
