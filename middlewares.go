package main

import (
	"context"
	"log"
	"net/http"
	"time"
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
		c, err := r.Cookie(SESSION_COOKIE_KEY)
		if err != nil {
			handleResponseError(err, w, http.StatusUnauthorized)
			return
		}
		userClaim, err := parseToken(c.Value)
		if err != nil {
			handleResponseError(err, w, http.StatusInternalServerError)
			return
		}
		ctx := context.WithValue(r.Context(), USER_CONTEXT_KEY, *userClaim)
		f(w, r.WithContext(ctx))
	})
}
