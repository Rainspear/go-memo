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

type Topic struct {
	Id          interface{}  `json:"id,omitempty" bson:"_id,omitempty"`
	AuthorId    interface{}  `json:"author_id" bson:"author_id"`
	Author      UserResponse `json:"author" bson:"author"`
	Title       string       `json:"title" bson:"title"`
	Description string       `json:"description" bson:"description"`
	CreatedDate int64        `json:"created_date" bson:"created_date"`
	LastUpdate  int64        `json:"last_update" bson:"last_update"`
}

type TopicResponse struct {
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	CreatedDate int64  `json:"created_date" bson:"created_date"`
	LastUpdate  int64  `json:"last_update" bson:"last_update"`
}

func getTopicById(ctx context.Context, topicId primitive.ObjectID) (*Topic, error) {
	if topicId.IsZero() {
		return nil, fmt.Errorf("topic_id is required")
	}
	var topic Topic
	filter := bson.D{{Key: "_id", Value: topicId}}
	coll := client.Database(database).Collection(TOPIC_COLLECTION)
	err := coll.FindOne(ctx, filter).Decode(&topic)
	if err != nil {
		return nil, err
	}
	return &topic, nil
}

func getTopics(w http.ResponseWriter, req *http.Request) {
	// coll := client.Database(database).Collection(TOPIC_COLLECTION)
	// from_date, err := strconv.ParseInt(req.FormValue("from_date"), 10, 64)
	// if err != nil {
	// 	from_date = 0
	// }
	// to_date, err := strconv.ParseInt(req.FormValue("to_date"), 10, 64)
	// if err != nil {
	// 	to_date = time.Now().Unix()
	// }
	// filterRepetitionStage := bson.D{
	// 	{Key: "$addFields",
	// 		Value: bson.D{
	// 			{Key: "repetition",
	// 				Value: bson.D{
	// 					{Key: "$filter",
	// 						Value: bson.D{
	// 							{Key: "input", Value: "$repetition"},
	// 							{Key: "as", Value: "repetition"},
	// 							{Key: "cond",
	// 								Value: bson.D{
	// 									{Key: "$and",
	// 										Value: bson.A{
	// 											bson.D{
	// 												{Key: "$gte",
	// 													Value: bson.A{
	// 														"$$repetition.time",
	// 														from_date,
	// 													},
	// 												},
	// 											},
	// 											bson.D{
	// 												{Key: "$lte",
	// 													Value: bson.A{
	// 														"$$repetition.time",
	// 														to_date,
	// 													},
	// 												},
	// 											},
	// 										},
	// 									},
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
	// addFieldStage := bson.D{
	// 	{Key: "$addFields",
	// 		Value: bson.D{
	// 			{Key: "repetition",
	// 				Value: bson.D{
	// 					{Key: "$sortArray",
	// 						Value: bson.D{
	// 							{Key: "input", Value: "$repetition"},
	// 							{Key: "sortBy", Value: bson.D{{Key: "time", Value: 1}}},
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
	loggedUser := req.Context().Value(USER_CONTEXT_KEY).(User)
	coll := client.Database(database).Collection(TOPIC_COLLECTION)
	matchStage := bson.D{{Key: "$match", Value: bson.D{{Key: "author_id", Value: loggedUser.Id}}}}
	cursor, err := coll.Aggregate(req.Context(), mongo.Pipeline{matchStage})
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
	id, err := primitive.ObjectIDFromHex(params["id"])
	topic, err := getTopicById(req.Context(), id)
	if handleResponseError(err, w, http.StatusInternalServerError) {
		return
	}
	handleResponseSuccess(topic, w, http.StatusOK)
	// from_date, err := strconv.ParseInt(req.FormValue("from_date"), 10, 64)
	// if err != nil {
	// 	from_date = 0
	// }
	// to_date, err := strconv.ParseInt(req.FormValue("to_date"), 10, 64)
	// if err != nil {
	// 	to_date = time.Now().Unix()
	// }
	// id, err := primitive.ObjectIDFromHex(params["id"])
	// if handleResponseError(err, w, http.StatusBadRequest) {
	// 	return
	// }
	// coll := client.Database(database).Collection(TOPIC_COLLECTION)
	// matchStage := bson.D{{Key: "$match", Value: bson.D{{Key: "_id", Value: id}}}}
	// filterRepetitionStage := bson.D{
	// 	{Key: "$addFields",
	// 		Value: bson.D{
	// 			{Key: "repetition",
	// 				Value: bson.D{
	// 					{Key: "$filter",
	// 						Value: bson.D{
	// 							{Key: "input", Value: "$repetition"},
	// 							{Key: "as", Value: "repetition"},
	// 							{Key: "cond",
	// 								Value: bson.D{
	// 									{Key: "$and",
	// 										Value: bson.A{
	// 											bson.D{
	// 												{Key: "$gte",
	// 													Value: bson.A{
	// 														"$$repetition.time",
	// 														from_date},
	// 												},
	// 											},
	// 											bson.D{
	// 												{Key: "$lte",
	// 													Value: bson.A{
	// 														"$$repetition.time",
	// 														to_date,
	// 													},
	// 												},
	// 											},
	// 										},
	// 									},
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
	// addFieldStage := bson.D{
	// 	{Key: "$addFields",
	// 		Value: bson.D{
	// 			{Key: "repetition",
	// 				Value: bson.D{
	// 					{Key: "$sortArray",
	// 						Value: bson.D{
	// 							{Key: "input", Value: "$repetition"},
	// 							{Key: "sortBy", Value: bson.D{{Key: "time", Value: 1}}},
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
	// limitStage := bson.D{{Key: "$limit", Value: 1}}
	// ----------------------------------------------------------------
	// cursor, err := coll.Aggregate(req.Context(), mongo.Pipeline{matchStage, limitStage})
	// if handleResponseError(err, w, http.StatusInternalServerError) {
	// 	return
	// }
	// var topics []Topic
	// err = cursor.All(req.Context(), &topics)
	// if handleResponseError(err, w, http.StatusInternalServerError) {
	// 	return
	// }
	// if len(topics) > 0 {
	// 	handleResponseSuccess(topics[0], w, http.StatusOK)
	// 	return
	// }
	// handleResponseSuccess(topics, w, http.StatusOK)

}

func createTopic(w http.ResponseWriter, req *http.Request) {
	loggedUser := (req.Context().Value(USER_CONTEXT_KEY)).(User)
	var data Topic
	err := json.NewDecoder(req.Body).Decode(&data)
	if handleResponseError(err, w, http.StatusBadRequest) {
		return
	}
	t := time.Now().Unix()
	var ur UserResponse
	err = copier.Copy(&ur, loggedUser)
	if err != nil {
		handleResponseError(err, w, http.StatusInternalServerError)
	}
	data.CreatedDate = t
	data.LastUpdate = t
	data.AuthorId = loggedUser.Id
	data.Author = ur
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
