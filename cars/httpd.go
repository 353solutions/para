package main

import (
	// _ "expvar" // for side effect of registering in default mux

	"context"
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
$ ~/go/bin/go1.24.3 version
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
	// mux.HandleFunc("POST /rides", api.insertHandler)
	// Wrap single handler
	h := topMiddleware(log, http.HandlerFunc(api.insertHandler))
	// Example: Wrap everything
	// h := topMiddleware(log, mux)
	// server.Handler = h
	mux.Handle("POST /rides", h)
	mux.Handle("GET /debug/vars", expvar.Handler())
	mux.HandleFunc("GET /rides/{id}", api.getHandler)

	// Built-in router
	// End with / - prefix match (/users/: /users/a, /users/b ...)
	// Otherwise - exact match

	srv := http.Server{
		Addr:    ":8080",
		Handler: mux,

		ReadTimeout: time.Second,
		// TODO: More timeouts
	}

	//mapDemo()

	log.Info("server starting", "address", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}

func (a *API) getHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	a.log.Info("get", "id", id)

	fmt.Fprintln(w, "OK")
}

func mapDemo() {
	m := map[any]any{}

	type keyType string
	var key keyType = "a"
	m["a"] = 1
	m[key] = 2
	fmt.Println(m) // [a:2 a:1]
}

type ctxVars struct {
	Login string
	// TODO: More info
}

// Unexported so only our code can use it
type ctxKeyType string

var ctxKey ctxKeyType = "params"

func getCtxVars(ctx context.Context) *ctxVars {
	val := ctx.Value(ctxKey)
	if val == nil {
		return &ctxVars{}
	}

	v, ok := val.(*ctxVars)
	if !ok {
		// TODO: Log? panic?
		return &ctxVars{}
	}
	return v
}

// Middleware: function that gets http.Handler and return http.Handler
func topMiddleware(log *slog.Logger, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// pre handler

		log.Info("request", "method", r.Method, "path", r.URL.Path)

		user, passwd, ok := r.BasicAuth()
		if !ok || !(user == "joe" && passwd == "baz00ka") {
			log.Error("bad auth", "user", user, "remote", r.RemoteAddr)
			http.Error(w, "bad auth", http.StatusUnauthorized)
			return
		}

		// Passed down to handler
		v := ctxVars{
			Login: user,
		}
		ctx := context.WithValue(r.Context(), ctxKey, &v)
		r = r.WithContext(ctx)

		h.ServeHTTP(w, r) // next handler

		// post handler
		// Wrap to get things (such as status code)
	}

	return http.HandlerFunc(fn)
}

func (a *API) insertHandler(w http.ResponseWriter, r *http.Request) {
	v := getCtxVars(r.Context())
	a.log.Info("insert", "login", v.Login)
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

/* Software layers

API
Business
Data

*/
