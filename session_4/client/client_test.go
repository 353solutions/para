package client

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestClient_HealthError(t *testing.T) {
	c := Client{
		BaseURL: "https://example.com",
	}
	c.c.Transport = mockTripper{}

	err := c.Health(context.Background())
	t.Logf("err: %v\n", err)
	if err == nil {
		t.Fatal("expected error")
	}
}

func (mockTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	return nil, fmt.Errorf("can't connect")
}

type mockTripper struct{}
