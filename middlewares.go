package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func addConfigMiddleware(f http.HandlerFunc) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		t := time.Now()
		f(w, r)
		log.Printf("%s %s %dms\n", r.RequestURI, r.Method, time.Since(t).Milliseconds())
	})
}

func authorizeUser(f http.HandlerFunc) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearer := r.Header.Get(AUTH_HEADER_KEY)
		if bearer == "" {
			handleResponseError(fmt.Errorf("header was empty"), w, http.StatusBadRequest)
			return
		}
		s := strings.Split(bearer, " ")
		if len(s) < 2 {
			handleResponseError(fmt.Errorf("not right to access"), w, http.StatusUnauthorized)
			return
		}
		token := s[1]
		if token == "" {
			handleResponseError(fmt.Errorf("invalid token"), w, http.StatusBadRequest)
			return
		}
		userClaim, err := parseToken(token)
		if err != nil {
			handleResponseError(err, w, http.StatusInternalServerError)
			return
		}
		filter := bson.D{
			{Key: "email", Value: userClaim.Email},
			{Key: "tokens", Value: bson.D{{Key: "$elemMatch", Value: bson.D{{Key: "token", Value: token}}}}},
		}
		var u User
		err = client.Database(database).Collection(USER_COLLECTION).FindOne(r.Context(), filter).Decode(&u)
		if err != nil {
			handleResponseError(fmt.Errorf("Can not authorized user"), w, http.StatusUnauthorized)
			return
		}
		if u.Email == "" {
			handleResponseError(fmt.Errorf("invalid logged user"), w, http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), USER_CONTEXT_KEY, u)
		f(w, r.WithContext(ctx))
	})
}
