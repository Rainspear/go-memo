package main

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Error      string `json:"error" bson:"error"`
	StatusCode int    `json:"status_code" bson:"status_code"`
}

type SuccessResponse struct {
	Data       interface{} `json:"data" bson:"data"`
	StatusCode int         `json:"status_code" bson:"status_code"`
}

type CreatedResponse struct {
	Id string `json:"id"`
}

type UpdatedResponse struct {
	Id string `json:"id"`
}

type DeletedResponse struct {
	Id string `json:"id"`
}

func handleResponseSuccess(data interface{}, w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(SuccessResponse{
		Data:       data,
		StatusCode: statusCode,
	})
	return
}

func handlePanicError(err error) {
	if err != nil {
		panic(err)
	}
}

func handleResponseError(err error, w http.ResponseWriter, statusCode int) {
	if err != nil {
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(ErrorResponse{err.Error(), statusCode})
		return
	}
}
