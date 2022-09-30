package main

type ErrorResponse struct {
	Error      string `json:"error" bson:"error"`
	StatusCode int    `json:"status_code" bson:"status_code"`
}
