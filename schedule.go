package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Status string

type Level string

const (
	LevelEssential Level = "essential"
	LevelImportant Level = "important"
	LevelSemi      Level = "semi-important"
	LevelLess      Level = "less-important"
	LevelMinor     Level = "minor"
)

const (
	StatusSuccess Status = "success"
	StatusUntouch Status = "untouch"
	StatusFailure Status = "failure"
	StatusSkipped Status = "skipped"
)

type ScheduleParams struct {
	Id       interface{} `json:"id,omitempty" bson:"_id,omitempty"`
	FromDate int64       `json:"from_date,omitempty" bson:"from_date"`
	ToDate   int64       `json:"to_date,omitempty" bson:"to_date"`
}

type Schedule struct {
	Id          interface{}   `json:"id,omitempty" bson:"_id,omitempty"`
	TopicId     interface{}   `json:"topic_id" bson:"topic_id"`
	AuthorId    interface{}   `json:"author_id" bson:"author_id"`
	Topic       TopicResponse `json:"topic" bson:"topic"`
	Author      UserResponse  `json:"author" bson:"author"`
	Level       Level         `json:"level" bson:"level"`
	Status      Status        `json:"status" bson:"status"`
	Time        int64         `json:"time" bson:"time"`
	CreatedDate int64         `json:"created_date" bson:"created_date"`
	LastUpdate  int64         `json:"last_update" bson:"last_update"`
}

type ScheduleResponse struct {
	Level       Level  `json:"level" bson:"level"`
	Status      Status `json:"status" bson:"status"`
	Time        int64  `json:"time" bson:"time"`
	CreatedDate int64  `json:"created_date" bson:"created_date"`
	LastUpdate  int64  `json:"last_update" bson:"last_update"`
}

func getSchedulesByFilter(ctx context.Context, pipeline mongo.Pipeline) ([]Schedule, error) {
	var schedules []Schedule
	coll := client.Database(database).Collection(SCHEDULE_COLLECTION)
	cursor, err := coll.Aggregate(ctx, pipeline)
	err = cursor.All(ctx, &schedules)
	if err != nil {
		return nil, err
	}
	return schedules, nil
}

func getSchedules(w http.ResponseWriter, req *http.Request) {
	loggedUser := (req.Context().Value(USER_CONTEXT_KEY)).(User)
	from_date, err := strconv.ParseInt(req.FormValue("from_date"), 10, 64)
	if err != nil {
		from_date = 0
	}
	to_date, err := strconv.ParseInt(req.FormValue("to_date"), 10, 64)
	if err != nil {
		to_date = time.Date(100000, 12, 31, 11, 59, 59, 999, time.UTC).Unix()
	}
	matchDateAndAuthorStage := bson.D{{Key: "$match", Value: bson.D{{Key: "author_id", Value: loggedUser.Id.(primitive.ObjectID)}, {Key: "$and",
		Value: bson.A{
			bson.D{{Key: "time", Value: bson.D{{Key: "$gte", Value: from_date}}}},
			bson.D{{Key: "time", Value: bson.D{{Key: "$lte", Value: to_date}}}},
		},
	}}}}
	pipe := mongo.Pipeline{matchDateAndAuthorStage}
	topicId, err := primitive.ObjectIDFromHex(req.FormValue("topic_id"))
	if err == nil {
		topic, err := getTopicById(req.Context(), topicId)
		if handleResponseError(err, w, http.StatusBadRequest) {
			return
		}
		if topic.AuthorId != loggedUser.Id {
			handleResponseError(fmt.Errorf("topic was not belong to this user"), w, http.StatusUnauthorized)
			return
		}
		fmt.Println("matchStage reach here")
		matchTopicStage := bson.D{{Key: "$match", Value: bson.D{{Key: "topic_id", Value: topicId}}}}
		pipe = mongo.Pipeline{matchDateAndAuthorStage, matchTopicStage}
	}
	s, err := getSchedulesByFilter(req.Context(), pipe)
	if handleResponseError(err, w, http.StatusInternalServerError) {
		return
	}
	handleResponseSuccess(s, w, http.StatusOK)
}

func createSchedule(w http.ResponseWriter, req *http.Request) {
	loggedUser := (req.Context().Value(USER_CONTEXT_KEY)).(User)
	var data Schedule
	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		handleResponseError(err, w, http.StatusBadRequest)
	}
	// verifying topic's author same as schedule author for authorization
	if data.TopicId == "" {
		handleResponseError(fmt.Errorf("this schedule does not belong to any topic"), w, http.StatusBadRequest)
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
	// filter field before response
	var topicResponse TopicResponse
	err = copier.Copy(&topicResponse, topic)
	var ur UserResponse
	err = copier.Copy(&ur, loggedUser)
	t := time.Now().Unix()
	if err != nil {
		handleResponseError(err, w, http.StatusInternalServerError)
		return
	}
	// insert to db
	data.CreatedDate = t
	data.LastUpdate = t
	data.AuthorId = loggedUser.Id
	data.Author = ur
	data.TopicId = topicId
	data.Topic = topicResponse
	collSchedule := client.Database(database).Collection(SCHEDULE_COLLECTION)
	result, err := collSchedule.InsertOne(req.Context(), &data)
	if handleResponseError(err, w, http.StatusInternalServerError) {
		return
	}
	id := result.InsertedID.(primitive.ObjectID)
	handleResponseSuccess(CreatedResponse{id.Hex()}, w, http.StatusCreated)
}

func updateSchedule(w http.ResponseWriter, req *http.Request) {
	var data Schedule
	err := json.NewDecoder(req.Body).Decode(&data)
	if handleResponseError(err, w, http.StatusBadRequest) {
		return
	}
	coll := client.Database(database).Collection(SCHEDULE_COLLECTION)
	params := mux.Vars(req)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if handleResponseError(err, w, http.StatusBadRequest) {
		return
	}
	update := bson.D{{Key: "$set", Value: &data}}
	filter := bson.D{{Key: "_id", Value: id}}
	result, err := coll.UpdateOne(req.Context(), filter, update)
	if handleResponseError(err, w, http.StatusBadRequest) {
		return
	}
	handleResponseSuccess(result, w, http.StatusOK)
}

func deleteSchedule(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if handleResponseError(err, w, http.StatusBadRequest) {
		return
	}
	filter := bson.D{{Key: "_id", Value: id}}
	coll := client.Database(database).Collection(SCHEDULE_COLLECTION)
	result, err := coll.DeleteOne(req.Context(), filter)
	if handleResponseError(err, w, http.StatusInternalServerError) {
		return
	}
	handleResponseSuccess(result, w, http.StatusOK)
}
