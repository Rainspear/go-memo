package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id          interface{} `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string      `json:"name" bson:"name"`
	Avatar      string      `json:"avatar," bson:"avatar,"`
	Email       string      `json:"email" bson:"email"`
	Password    string      `json:"password" bson:"password"`
	CreatedDate time.Time   `json:"created_date" bson:"created_date"`
	LastUpdate  time.Time   `json:"last_update" bson:"last_update"`
	Tokens      []Token     `json:"tokens" bson:"tokens"`
}

type UserResponse struct {
	Name   string `json:"name" bson:"name"`
	Avatar string `json:"avatar," bson:"avatar,"`
	Email  string `json:"email" bson:"email"`
}

type Token struct {
	Token       string    `json:"token" bson:"token"`
	CreatedDate time.Time `json:"created_date" bson:"created_date"`
}

func getUsers(w http.ResponseWriter, req *http.Request) {
	coll := client.Database(database).Collection(USER_COLLECTION)
	matchStage := bson.D{{Key: "$match", Value: bson.D{}}}
	addFieldStage := bson.D{{Key: "$project", Value: bson.D{{Key: "password", Value: 0}}}}
	cursor, err := coll.Aggregate(req.Context(), mongo.Pipeline{matchStage, addFieldStage})
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
	loggedUser := (req.Context().Value(USER_CONTEXT_KEY)).(UserClaims)
	coll := client.Database(database).Collection(USER_COLLECTION)
	var u UserResponse
	err := coll.FindOne(req.Context(), bson.D{{Key: "email", Value: loggedUser.Email}}).Decode(&u)
	if handleResponseError(err, w, http.StatusInternalServerError) {
		return
	}
	handleResponseSuccess(u, w, http.StatusOK)
}

func updateUser(w http.ResponseWriter, req *http.Request) {

}

func deleteUser(w http.ResponseWriter, req *http.Request) {

}

func signin(w http.ResponseWriter, req *http.Request) {
	var user User
	err := json.NewDecoder(req.Body).Decode(&user)
	if handleResponseError(err, w, http.StatusBadRequest) {
		return
	}
	// validate input
	if user.Email == "" || user.Password == "" {
		handleResponseError(fmt.Errorf("you must specify email and password"), w, http.StatusBadRequest)
		return
	}
	// normailize data
	coll := client.Database(database).Collection(USER_COLLECTION)
	filter := bson.D{{Key: "email", Value: user.Email}}
	var u User
	err = coll.FindOne(req.Context(), filter).Decode(&u)
	if handleResponseError(err, w, http.StatusBadRequest) {
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(user.Password))
	if handleResponseError(err, w, http.StatusBadRequest) {
		return
	}
	// create token
	userClaims := UserClaims{uuid.New().String(), user.Email, jwt.StandardClaims{}}
	token, err := createToken(&userClaims)
	if handleResponseError(err, w, http.StatusInternalServerError) {
		return
	}
	tokenObject := Token{token, time.Now()}
	u.Tokens = append(u.Tokens, tokenObject)
	_, err = coll.UpdateOne(req.Context(), bson.D{{Key: "email", Value: u.Email}},
		bson.D{
			{Key: "$set", Value: bson.D{
				{Key: "tokens", Value: u.Tokens}}}})
	if handleResponseError(err, w, http.StatusInternalServerError) {
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:  SESSION_COOKIE_KEY,
		Value: token,
	})
	handleResponseSuccess(UserResponse{Name: u.Name, Avatar: u.Avatar, Email: u.Email}, w, http.StatusOK)
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
	coll := client.Database(database).Collection(string(USER_COLLECTION))
	var existedUser User
	err = coll.FindOne(req.Context(), bson.D{{Key: "email", Value: user.Email}}).Decode(&existedUser)
	if existedUser.Email != "" && err == nil { // existed user in database
		handleResponseError(fmt.Errorf("email already existed"), w, http.StatusBadRequest)
		return
	}
	// create data to save
	t := time.Now()
	u := UserClaims{uuid.New().String(), user.Email, jwt.StandardClaims{}}
	token, err := createToken(&u)
	if handleResponseError(err, w, http.StatusInternalServerError) {
		return
	}
	userTokenObject := Token{token, t}
	user.Password = string(bs)
	user.CreatedDate = t
	user.LastUpdate = t
	user.Tokens = append(user.Tokens, userTokenObject)
	_, err = coll.InsertOne(req.Context(), &user)
	if handleResponseError(err, w, http.StatusInternalServerError) {
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:  SESSION_COOKIE_KEY,
		Value: token,
	})
	handleResponseSuccess(UserResponse{Name: user.Name, Avatar: user.Avatar, Email: user.Email}, w, http.StatusCreated)
}

func signout(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie(SESSION_COOKIE_KEY)
	if err != nil {
		handleResponseError(err, w, http.StatusBadRequest)
		return
	}
	c = &http.Cookie{Name: SESSION_COOKIE_KEY, Value: "", MaxAge: -1}
	http.SetCookie(w, c)
	handleResponseSuccess("logout successfully", w, http.StatusOK)
}
