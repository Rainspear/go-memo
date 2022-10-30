package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
	StatusUntouch Status = "untouch"
	StatusFailure Status = "failure"
	StatusSkipped Status = "skipped"
)

type Repetition struct {
	Time   int64  `json:"time" bson:"time"`
	Level  Level  `json:"level" bson:"level"`
	Status Status `json:"status" bson:"status"`
}

type Topic struct {
	Id          interface{}  `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string       `json:"title" bson:"title"`
	Repetition  []Repetition `json:"repetition" bson:"repetition"`
	Description string       `json:"description" bson:"description"`
	CreatedDate int64        `json:"created_date" bson:"created_date"`
	LastUpdate  int64        `json:"last_update" bson:"last_update"`
}

func getTopics(w http.ResponseWriter, req *http.Request) {
	coll := client.Database(database).Collection(TOPIC_COLLECTION)
	from_date, err := strconv.ParseInt(req.FormValue("from_date"), 10, 64)
	if err != nil {
		from_date = 0
	}
	to_date, err := strconv.ParseInt(req.FormValue("to_date"), 10, 64)
	if err != nil {
		to_date = time.Now().Unix()
	}
	fmt.Println("from_date: ", from_date)
	fmt.Println("to_date: ", to_date)
	filterRepetitionStage := bson.D{
		{Key: "$addFields",
			Value: bson.D{
				{Key: "repetition",
					Value: bson.D{
						{Key: "$filter",
							Value: bson.D{
								{Key: "input", Value: "$repetition"},
								{Key: "as", Value: "repetition"},
								{Key: "cond",
									Value: bson.D{
										{Key: "$and",
											Value: bson.A{
												bson.D{
													{Key: "$gte",
														Value: bson.A{
															"$$repetition.time",
															from_date,
														},
													},
												},
												bson.D{
													{Key: "$lte",
														Value: bson.A{
															"$$repetition.time",
															to_date,
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	addFieldStage := bson.D{
		{Key: "$addFields",
			Value: bson.D{
				{Key: "repetition",
					Value: bson.D{
						{Key: "$sortArray",
							Value: bson.D{
								{Key: "input", Value: "$repetition"},
								{Key: "sortBy", Value: bson.D{{Key: "time", Value: 1}}},
							},
						},
					},
				},
			},
		},
	}
	cursor, err := coll.Aggregate(req.Context(), mongo.Pipeline{filterRepetitionStage, addFieldStage})
	if handleResponseError(err, w, http.StatusInternalServerError) {
		return
	}
	var topics []Topic
	err = cursor.All(req.Context(), &topics)
	if handleResponseError(err, w, http.StatusInternalServerError) {
		return
	}
	handleResponseSuccess(topics, w, http.StatusOK)
}

func getTopic(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	from_date, err := strconv.ParseInt(req.FormValue("from_date"), 10, 64)
	if err != nil {
		from_date = 0
	}
	to_date, err := strconv.ParseInt(req.FormValue("to_date"), 10, 64)
	if err != nil {
		to_date = time.Now().Unix()
	}
	id, err := primitive.ObjectIDFromHex(params["id"])
	if handleResponseError(err, w, http.StatusBadRequest) {
		return
	}
	fmt.Println(id)
	coll := client.Database(database).Collection(TOPIC_COLLECTION)
	matchStage := bson.D{{Key: "$match", Value: bson.D{{Key: "_id", Value: id}}}}
	fmt.Println("from_date: ", from_date)
	fmt.Println("to_date: ", to_date)
	filterRepetitionStage := bson.D{
		{Key: "$addFields",
			Value: bson.D{
				{Key: "repetition",
					Value: bson.D{
						{Key: "$filter",
							Value: bson.D{
								{Key: "input", Value: "$repetition"},
								{Key: "as", Value: "repetition"},
								{Key: "cond",
									Value: bson.D{
										{Key: "$and",
											Value: bson.A{
												bson.D{
													{Key: "$gte",
														Value: bson.A{
															"$$repetition.time",
															from_date},
													},
												},
												bson.D{
													{Key: "$lte",
														Value: bson.A{
															"$$repetition.time",
															to_date,
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	addFieldStage := bson.D{
		{Key: "$addFields",
			Value: bson.D{
				{Key: "repetition",
					Value: bson.D{
						{Key: "$sortArray",
							Value: bson.D{
								{Key: "input", Value: "$repetition"},
								{Key: "sortBy", Value: bson.D{{Key: "time", Value: 1}}},
							},
						},
					},
				},
			},
		},
	}
	limitStage := bson.D{{Key: "$limit", Value: 1}}
	// ----------------------------------------------------------------
	cursor, err := coll.Aggregate(req.Context(), mongo.Pipeline{matchStage, filterRepetitionStage, addFieldStage, limitStage})
	if handleResponseError(err, w, http.StatusInternalServerError) {
		return
	}
	var topics []Topic
	err = cursor.All(req.Context(), &topics)
	if handleResponseError(err, w, http.StatusInternalServerError) {
		return
	}
	if len(topics) > 0 {
		handleResponseSuccess(topics[0], w, http.StatusOK)
		return
	}
	handleResponseSuccess(topics, w, http.StatusOK)

}

func createTopic(w http.ResponseWriter, req *http.Request) {
	var data Topic
	err := json.NewDecoder(req.Body).Decode(&data)
	if handleResponseError(err, w, http.StatusBadRequest) {
		return
	}
	t := time.Now().Unix()
	data.CreatedDate = t
	data.LastUpdate = t
	coll := client.Database(database).Collection(TOPIC_COLLECTION)
	result, err := coll.InsertOne(req.Context(), &data)
	if handleResponseError(err, w, http.StatusInternalServerError) {
		return
	}
	id := result.InsertedID.(primitive.ObjectID)
	handleResponseSuccess(CreatedResponse{id.Hex()}, w, http.StatusCreated)
}

func updateTopic(w http.ResponseWriter, req *http.Request) {
	var data Topic
	err := json.NewDecoder(req.Body).Decode(&data)
	if handleResponseError(err, w, http.StatusBadRequest) {
		return
	}
	coll := client.Database(database).Collection(TOPIC_COLLECTION)
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

func deleteTopic(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if handleResponseError(err, w, http.StatusBadRequest) {
		return
	}
	coll := client.Database(database).Collection(TOPIC_COLLECTION)
	filter := bson.D{{Key: "_id", Value: id}}
	result, err := coll.DeleteOne(req.Context(), filter)
	if handleResponseError(err, w, http.StatusInternalServerError) {
		return
	}
	handleResponseSuccess(result, w, http.StatusOK)
}
