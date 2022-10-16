package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/internal/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id          interface{} `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string      `json:"name" bson:"name"`
	Avatar      string      `json:"avatar,omitempty" bson:"avatar,omitempty"`
	Email       string      `json:"email" bson:"email"`
	Password    string      `json:"-" bson:"password"`
	CreatedDate time.Time   `json:"created_date" bson:"created_date"`
	LastUpdate  time.Time   `json:"last_update" bson:"last_update"`
	Tokens      []string    `json:"tokens,omitempty" bson:"tokens,omitempty"`
}

func getUsers(w http.ResponseWriter, req *http.Request) {
	coll := client.Database(database).Collection(USER_COLLECTION)
	matchStage := bson.D{{Key: "$match", Value: bson.D{}}}
	cursor, err := coll.Aggregate(req.Context(), mongo.Pipeline{matchStage})
	if handleResponseError(err, w, http.StatusInternalServerError) {
		return
	}
	var users []User
	err = cursor.All(req.Context(), &users)
	if handleResponseError(err, w, http.StatusInternalServerError) {
		return
	}
	handleResponseSuccess(users, w, http.StatusOK)
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
	var existedUser User
	err = coll.FindOne(req.Context(), bson.D{{Key: "email", Value: user.Email}}).Decode(&existedUser)
	if existedUser.Email != "" && err == nil { // existed user in database
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{"user is already existed " + err.Error(), http.StatusBadRequest})
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
	u := UserClaims{uuid.New(), result.InsertedID.(primitive.ObjectID).Hex(), jwt.StandardClaims{}}
	token, err := createToken(&u)
	if err != nil {
		handleResponseError(err, w, http.StatusInternalServerError)
	}
	http.SetCookie(w, &http.Cookie{
		Name:  "SessionID",
		Value: token,
	})
	// create jwt token and return response
	fmt.Println("token: ", token)
	handleResponseToken(token, w, http.StatusCreated)
}

func signout(w http.ResponseWriter, req *http.Request) {

}
