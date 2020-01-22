package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetHomepage(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetWelcome)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "Welcome to Homepage!"

	if rr.Body.String() != expected {
		t.Errorf("handler returned wrong body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestPostPing(t *testing.T) {
	pingJSON := []byte(`{"body":"ping"}`)

	req, err := http.NewRequest("POST", "/ping", bytes.NewBuffer(pingJSON))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PostPing)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "pong"
	if rr.Body.String() != expected {
		t.Errorf("handler returned wrong body: got %v want %v",
			rr.Body.String(), expected)
	}

}
