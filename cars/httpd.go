package main

import (
	// _ "expvar" // for side effect of registering in default mux
	"encoding/json"
	"expvar"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

/* Exercise:
Create a database and add it to API
Write a handler that gets an event and inserts it to the database

$ curl -d@_extra/event.json http://localhost:8080/rides

CRUD -> REST HTTP Verbs
- Create: POST
- Read: GET
- Update: PUT, PATCH
- Delete: DELETE
*/

/* Install go 1.24
$ ./_extra/update-go 1.24.3
$ ~/go/bin/go1.24.3 verion
$ alias go=~/go/bin/go1.24.3
*/

func main() {
	log := slog.Default().With("app", "cars")
	db, err := NewDB("cars.db")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: create db - %s\n", err)
		os.Exit(1)
	}
	defer db.Close()

	api := API{
		log: log, // dependency injection
		db:  db,
	}
	// Routing
	// Will use default router (mux)
	// http.HandleFunc("GET /health", healthHandler)
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", api.healthHandler)
	mux.HandleFunc("POST /rides", api.insertHandler)
	mux.Handle("GET /debug/vars", expvar.Handler())

	srv := http.Server{
		Addr:    ":8080",
		Handler: mux,

		ReadTimeout: time.Second,
		// TODO: More timeouts
	}

	log.Info("server starting", "address", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}

func (a *API) insertHandler(w http.ResponseWriter, r *http.Request) {
	// Parse + Validate data
	var e Event
	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		a.log.Error("bad request", "error", err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	if err := e.Validate(); err != nil {
		a.log.Error("invalid event", "error", err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	// Work
	if err := a.db.Insert(e); err != nil {
		a.log.Error("insert", "error", err)
		http.Error(w, "can't insert", http.StatusInternalServerError)
		return
	}

	// Output
	resp := map[string]any{
		"id": e.ID,
	}
	if err := writeJSON(w, resp); err != nil {
		// Can't send to client
		a.log.Error("write JSON", "error", err)
	}
}

func writeJSON(w http.ResponseWriter, resp any) error {
	w.Header().Set("content-type", "application/json")
	return json.NewEncoder(w).Encode(resp)
}

func (a *API) healthHandler(w http.ResponseWriter, r *http.Request) {
	a.log.Info("health", "path", r.URL.Path)
	// TODO: Real health check
	fmt.Fprintln(w, "OK")
}

type API struct {
	log *slog.Logger
	db  *DB
}
