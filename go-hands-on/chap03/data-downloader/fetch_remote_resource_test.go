package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func startTestHTTPServer() *httptest.Server {
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "Hello, World")
		}),
	)
	return ts
}

func TestFetchRemoteResource(t *testing.T) {
	ts := startTestHTTPServer()
	defer ts.Close()

	expected := "Hello, World"
	body, err := fetchRemoteResource(ts.URL)
	if err != nil {
		t.Fatalf("Error fetching remote resource: %v", err)
	}
	if string(body) != expected {
		t.Fatalf("Expected body to be: %v, got: %v", expected, string(body))
	}
}
