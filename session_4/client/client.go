package client

import (
	"context"
	"fmt"
	"net/http"
)

func (c *Client) Health(ctx context.Context) error {
	url := fmt.Sprintf("%s/health", c.BaseURL)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("Health: %q request - %w", url, err)
	}

	resp, err := c.c.Do(req)
	if err != nil {
		return fmt.Errorf("Health: GET %q - %w", url, err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Health: GET %q - %s", url, resp.Status)
	}

	return nil
}

type Client struct {
	BaseURL string

	c http.Client
}
