package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Memo struct {
	Id          interface{}   `json:"id,omitempty" bson:"_id,omitempty"`
	TopicId     interface{}   `json:"topic_id" bson:"topic_id"`
	AuthorId    interface{}   `json:"author_id" bson:"author_id"`
	Topic       TopicResponse `json:"topic" bson:"topic"`
	Author      UserResponse  `json:"author" bson:"author"`
	Content     string        `json:"content" bson:"content"`
	Question    string        `json:"question" bson:"question"`
	Answer      []string      `json:"answer" bson:"answer"`
	CreatedDate int64         `json:"created_date" bson:"created_date"`
	LastUpdate  int64         `json:"last_update" bson:"last_update"`
}

type MemoResponse struct {
	Content  string   `json:"content" bson:"content"`
	Question string   `json:"question" bson:"question"`
	Answer   []string `json:"answer" bson:"answer"`
}

func getMemoById(ctx context.Context, memoId primitive.ObjectID) (*Memo, error) {
	if memoId.IsZero() {
		return nil, fmt.Errorf("memo_id is required")
	}
	var memo Memo
	filter := bson.D{{Key: "_id", Value: memoId}}
	coll := client.Database(database).Collection(MEMO_COLLECTION)
	err := coll.FindOne(ctx, filter).Decode(&memo)
	if err != nil {
		return nil, err
	}
	return &memo, nil
}

func getMemosByFilter(ctx context.Context, pipeline mongo.Pipeline) ([]Memo, error) {
	var memos []Memo
	coll := client.Database(database).Collection(MEMO_COLLECTION)
	cursor, err := coll.Aggregate(ctx, pipeline)
	err = cursor.All(ctx, &memos)
	if err != nil {
		return nil, err
	}
	return memos, nil
}

func getMemos(w http.ResponseWriter, req *http.Request) {
	loggedUser := (req.Context().Value(USER_CONTEXT_KEY)).(User)
	topicId, err := primitive.ObjectIDFromHex(req.FormValue("topic_id"))
	matchStage := bson.D{{Key: "$match", Value: bson.D{{Key: "author_id", Value: loggedUser.Id.(primitive.ObjectID)}}}}
	if err == nil {
		topic, err := getTopicById(req.Context(), topicId)
		if handleResponseError(err, w, http.StatusBadRequest) {
			return
		}
		if topic.AuthorId != loggedUser.Id {
			handleResponseError(fmt.Errorf("topic was not belong to this user"), w, http.StatusUnauthorized)
			return
		}
		matchStage = bson.D{{Key: "$match", Value: bson.D{{Key: "author_id", Value: loggedUser.Id.(primitive.ObjectID)}, {Key: "topic_id", Value: topicId}}}}
	}
	memos, err := getMemosByFilter(req.Context(), mongo.Pipeline{matchStage})
	if handleResponseError(err, w, http.StatusInternalServerError) {
		return
	}
	handleResponseSuccess(&memos, w, http.StatusOK)
}

func getMemo(w http.ResponseWriter, req *http.Request) {
	loggedUser := (req.Context().Value(USER_CONTEXT_KEY)).(User)
	params := mux.Vars(req)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if handleResponseError(err, w, http.StatusBadRequest) {
		return
	}
	memo, err := getMemoById(req.Context(), id)
	if handleResponseError(err, w, http.StatusNotFound) {
		return
	}
	if loggedUser.Id != memo.AuthorId {
		handleResponseError(fmt.Errorf("not have authorization to access this memo"), w, http.StatusUnauthorized)
		return
	}
	handleResponseSuccess(&memo, w, http.StatusOK)
}

func createMemo(w http.ResponseWriter, req *http.Request) {
	// get authorized user
	loggedUser := (req.Context().Value(USER_CONTEXT_KEY)).(User)
	// get user value
	var data Memo
	err := json.NewDecoder(req.Body).Decode(&data)
	if handleResponseError(err, w, http.StatusBadRequest) {
		return
	}
	if data.TopicId == "" {
		handleResponseError(fmt.Errorf("this memo does not belong to any topic"), w, http.StatusBadRequest)
		return
	}
	topicId, err := primitive.ObjectIDFromHex((data.TopicId).(string))
	if handleResponseError(err, w, http.StatusBadRequest) {
		return
	}
	topic, err := getTopicById(req.Context(), topicId)
	if topic.AuthorId != loggedUser.Id {
		handleResponseError(fmt.Errorf("topic was not belong to this user"), w, http.StatusUnauthorized)
		return
	}
	// normalize data
	var ur UserResponse
	err = copier.Copy(&ur, loggedUser)
	var tr TopicResponse
	err = copier.Copy(&tr, topic)
	if err != nil {
		handleResponseError(err, w, http.StatusInternalServerError)
	}
	t := time.Now().Unix()
	data.LastUpdate = t
	data.CreatedDate = t
	data.AuthorId = loggedUser.Id
	data.Author = ur
	data.TopicId = topicId
	data.Topic = tr
	coll := client.Database(database).Collection(MEMO_COLLECTION)
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
