package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Post struct {
	Body string `json:"body"`
}

var posts []Post

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Homepage!"))
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", welcome).Methods("GET")
	router.HandleFunc("/ping", ping).Methods("POST")
	http.ListenAndServe(":8000", router)
}
