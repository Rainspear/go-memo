package main

import (
	"log"
	"net/http"
	"time"
)

func addConfigMiddleware(f http.HandlerFunc) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		t := time.Now()
		f(w, r)
		log.Printf("%s %dms\n", r.RequestURI, time.Since(t).Milliseconds())
	})
}
