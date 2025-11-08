package sdk

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// HTTPClient is a client for interacting with the service HTTP API
type HTTPClient struct {
	baseURL    string
	httpClient *http.Client
	logger     logger
}

// Option is a functional option for configuring the HTTPClient
type Option func(*HTTPClient)

// WithHTTPClient allows setting a custom http.Client
func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *HTTPClient) {
		c.httpClient = httpClient
	}
}

// NewHTTPClient creates a new HTTP API client
func NewHTTPClient(baseURL string, logger logger, opts ...Option) *HTTPClient {
	c := &HTTPClient{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		logger: logger,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

// Health checks the health of the mercury API
func (c *HTTPClient) Health(ctx context.Context) (*HealthResponse, error) {
	endpoint := fmt.Sprintf("%s/healthz", c.baseURL)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	var health HealthResponse
	if err := c.doRequest(req, &health); err != nil {
		return nil, err
	}

	return &health, nil
}

// Add your HTTP client methods here
// Example:
// func (c *HTTPClient) GetUser(ctx context.Context, id string) (*User, error) {
//     var user User
//     err := c.doRequest(ctx, http.MethodGet, "/v1/users/"+id, nil, &user)
//     return &user, err
// }

// doRequest performs an HTTP request with JSON encoding/decoding
func (c *HTTPClient) doRequest(req *http.Request, result any) error {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer func() {
		// Drain and close to allow connection reuse
		_, _ = io.Copy(io.Discard, resp.Body)
		if err := resp.Body.Close(); err != nil {
			c.logger.Error("failed to close response body", "error", err)
		}
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for error responses
	if resp.StatusCode >= 400 {
		var errResp map[string]string
		if err := json.Unmarshal(body, &errResp); err != nil {
			return fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(body))
		}
		return fmt.Errorf("API error (%d): %s", resp.StatusCode, errResp["error"])
	}

	// Decode success response
	if err := json.Unmarshal(body, result); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	return nil
}
