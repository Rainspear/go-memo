package main

import (
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
	var user User
	err := json.NewDecoder(req.Body).Decode(&user)
	if handleResponseError(err, w, http.StatusBadRequest) {
		return
	}
	bs, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if handleResponseError(err, w, http.StatusBadRequest) {
		return
	}
	coll := client.Database(database).Collection(USER_COLLECTION)
	err = coll.FindOne(req.Context(), bson.D{{Key: "email", Value: user.Email}}).Err()
	if err == nil { // existed user in database
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{err.Error(), http.StatusBadRequest})
		return
	}
	// create data to save
	t := time.Now()
	user.Password = string(bs)
	user.CreatedDate = t
	user.LastUpdate = t
	result, err := coll.InsertOne(req.Context(), &user)
	if handleResponseError(err, w, http.StatusInternalServerError) {
		return
	}
	w.WriteHeader(http.StatusCreated)
	// create jwt token and return response
	json.NewEncoder(w).Encode(result)
}

func signout(w http.ResponseWriter, req *http.Request) {

}
