package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	fmt.Println("Welcome to go-toggle for GCP new")
	router := mux.NewRouter()
	router.HandleFunc("/posts", getPost).Methods("GET")

	http.ListenAndServe(":8000", router)
}

type Post struct {
	ID    string
	Title string
	Body  string
}

func getPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post1 Post
	post1.ID = "543210"
	post1.Title = "Welcome to go-toggle for GCP new"
	post1.Body = "Welcome to go-toggle for GCP new"
	json.NewEncoder(w).Encode(post1)
}
