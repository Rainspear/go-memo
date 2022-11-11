package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Memo struct {
	Id          interface{}  `json:"id,omitempty" bson:"_id,omitempty"`
	TopicId     interface{}  `json:"topic_id" bson:"topic_id"`
	AuthorId    interface{}  `json:"author_id" bson:"author_id"`
	Topic       Topic        `json:"topic" bson:"topic"`
	Author      UserResponse `json:"author" bson:"author"`
	Content     string       `json:"content" bson:"content"`
	Question    string       `json:"question" bson:"question"`
	Answer      []string     `json:"answer" bson:"answer"`
	CreatedDate int64        `json:"created_date" bson:"created_date"`
	LastUpdate  int64        `json:"last_update" bson:"last_update"`
}

func getMemos(w http.ResponseWriter, req *http.Request) {
	coll := client.Database(database).Collection(MEMO_COLLECTION)
	filter := bson.D{}
	cursor, err := coll.Find(req.Context(), filter)
	if handleResponseError(err, w, http.StatusInternalServerError) {
		return
	}
	memos := []Memo{}
	err = cursor.All(req.Context(), memos)
	if handleResponseError(err, w, http.StatusInternalServerError) {
		return
	}
	handleResponseSuccess(&memos, w, http.StatusOK)
}

func getMemo(w http.ResponseWriter, req *http.Request) {
	coll := client.Database(database).Collection(MEMO_COLLECTION)
	params := mux.Vars(req)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if handleResponseError(err, w, http.StatusBadRequest) {
		return
	}
	filter := bson.D{{Key: "_id", Value: id}}
	var memo Memo
	err = coll.FindOne(req.Context(), filter).Decode(&memo)
	if handleResponseError(err, w, http.StatusNotFound) {
		return
	}
	handleResponseSuccess(&memo, w, http.StatusOK)
}

func createMemo(w http.ResponseWriter, req *http.Request) {
	loggedUser := (req.Context().Value(USER_CONTEXT_KEY)).(User)
	coll := client.Database(database).Collection(MEMO_COLLECTION)
	t := time.Now().Unix()
	var data Memo
	err := json.NewDecoder(req.Body).Decode(&data)
	if handleResponseError(err, w, http.StatusBadRequest) {
		return
	}
	var ur UserResponse
	err = copier.Copy(&ur, loggedUser)
	if err != nil {
		handleResponseError(err, w, http.StatusInternalServerError)
	}
	data.LastUpdate = t
	data.CreatedDate = t
	data.AuthorId = loggedUser.Id
	data.Author = ur
	result, err := coll.InsertOne(req.Context(), data)
	if handleResponseError(err, w, http.StatusInternalServerError) {
		return
	}
	id := result.InsertedID.(primitive.ObjectID)
	handleResponseSuccess(CreatedResponse{id.Hex()}, w, http.StatusCreated)
}

func updateMemo(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if handleResponseError(err, w, http.StatusBadRequest) {
		return
	}
	filter := bson.D{{Key: "_id", Value: id}}
	var data Memo
	err = json.NewDecoder(req.Body).Decode(&data)
	if handleResponseError(err, w, http.StatusBadRequest) {
		return
	}
	data.LastUpdate = time.Now().Unix()
	update := bson.D{{Key: "$set", Value: &data}}
	coll := client.Database(database).Collection(MEMO_COLLECTION)
	result, err := coll.UpdateOne(req.Context(), filter, update)
	if handleResponseError(err, w, http.StatusBadRequest) {
		return
	}
	handleResponseSuccess(result, w, http.StatusOK)
}

func deleteMemo(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if handleResponseError(err, w, http.StatusBadRequest) {
		return
	}
	filter := bson.D{{Key: "_id", Value: id}}
	coll := client.Database(database).Collection(MEMO_COLLECTION)
	result, err := coll.DeleteOne(req.Context(), filter)
	if handleResponseError(err, w, http.StatusInternalServerError) {
		return
	}
	handleResponseSuccess(result, w, http.StatusOK)
}
