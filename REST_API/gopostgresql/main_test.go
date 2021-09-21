package main

import (
	
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPosts(t *testing.T) {
	req, err := http.NewRequest("GET", "/posts", nil)
	
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	
	handler := http.HandlerFunc(getPosts)

	handler.ServeHTTP(rr, req)
	
	
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	
	// Check the response body is what we expect.
	expected := `[
		{
			"id": "2",
			"title": "second"
		},
		{
			"id": "4",
			"title": "second"
		},
		{
			"id": "6",

			"title": "t"
		}
	]`

	
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
