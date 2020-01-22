package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type post struct {
	Body string
}

func GetWelcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Homepage!"))
}

func PostPing(w http.ResponseWriter, r *http.Request) {
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
	router.HandleFunc("/v0", GetWelcome).Methods("GET")
	router.HandleFunc("/v0/ping", PostPing).Methods("POST")
	http.ListenAndServe(":8000", router)
}
