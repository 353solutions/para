package main

import (
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAPI_Health(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/health", nil)

	api := API{
		log: slog.New(slog.DiscardHandler),
	}
	mux := buildMux(&api)
	mux.ServeHTTP(w, r)
	/* Bypass mux
	api.Health(w, mux)
	 api.Health(w, r)
	*/

	resp := w.Result()
	require.Equal(t, http.StatusOK, resp.StatusCode)
}
