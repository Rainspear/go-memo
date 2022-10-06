package main

type ErrorResponse struct {
	Error      string `json:"error" bson:"error"`
	StatusCode int    `json:"status_code" bson:"status_code"`
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
