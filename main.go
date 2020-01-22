package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type post struct {
	Body string
}

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Homepage!"))
}

func ping(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var p post
	err := decoder.Decode(&p)

	if err != nil {
		panic(err)
	}

	if p.Body == "ping" {
		w.Write([]byte("pong"))
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/v0", welcome).Methods("GET")
	router.HandleFunc("/v0/ping", ping).Methods("POST")
	http.ListenAndServe(":8000", router)
}
