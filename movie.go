package main

type Movie struct {
	ID        string   `json:"id"`
	Title     string   `json:"title"`
	Genres    []string `json:"genre"`
	Cast      []string `json:"cast"`
	Countries []string `json:"country"`
}
