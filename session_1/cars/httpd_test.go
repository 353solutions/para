package main

import (
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Option: Run the server, test the API
// See also testcontainers

func TestHealth(t *testing.T) {
	api := API{
		log: slog.Default(),
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/health", nil)

	// Problem: bypass routing, middleware ...
	// Solution: Write createMux() in main
	// m := createMux()
	// m.HandleHTTP(w, r)
	api.healthHandler(w, r)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Fatal(resp.StatusCode)
	}
}
