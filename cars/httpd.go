package main

import (
	// _ "expvar" // for side effect of registering in default mux
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
*/

func main() {
	log := slog.Default().With("app", "cars")

	api := API{
		log: log, // dependency injection
	}
	// Routing
	// Will use default router (mux)
	// http.HandleFunc("GET /health", healthHandler)
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", api.healthHandler)
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

type API struct {
	log *slog.Logger
}

func (a *API) healthHandler(w http.ResponseWriter, r *http.Request) {
	a.log.Info("health", "path", r.URL.Path)
	// TODO: Real health check
	fmt.Fprintln(w, "OK")
}
