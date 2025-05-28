package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	// DefaultBaseURL is the base URL for the Propaga API in production
	DefaultBaseURL = "https://api.propaga.io"

	// StagingBaseURL is the base URL for the Propaga API in staging
	StagingBaseURL = "https://staging-api.propaga.io"

	// DefaultTimeout is the default timeout for HTTP requests
	DefaultTimeout = 30 * time.Second
)

// Client is the HTTP client for interacting with the Propaga API
type Client struct {
	// BaseURL is the base URL for all API requests
	BaseURL string

	// HTTPClient is the underlying HTTP client
	HTTPClient *http.Client

	// APIKey is the API key for authentication
	APIKey string
}

// NewClient creates a new instance of the Propaga client
func NewClient(apiKey string, isStaging bool) *Client {
	url := DefaultBaseURL
	if isStaging {
		url = StagingBaseURL
	}
	return &Client{
		BaseURL: url,
		HTTPClient: &http.Client{
			Timeout: DefaultTimeout,
		},
		APIKey: apiKey,
	}
}

// NewClientWithOptions creates a new instance of the client with custom options
func NewClientWithOptions(apiKey, baseURL string, timeout time.Duration) *Client {
	return &Client{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: timeout,
		},
		APIKey: apiKey,
	}
}

// DoRequest performs an HTTP request to the Propaga API
func (c *Client) DoRequest(method, path string, body interface{}, result interface{}) error {
	// Build the full URL
	url := fmt.Sprintf("%s%s", c.BaseURL, path)

	// Prepare the request body if it exists
	var reqBody io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("error serializing request body: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonBody)
	}

	// Create the HTTP request
	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return fmt.Errorf("error creating HTTP request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", c.APIKey)

	// Perform the request
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("error performing HTTP request: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}

	// Check the status code
	if resp.StatusCode >= 400 {
		return fmt.Errorf("API error (code %d): %s", resp.StatusCode, string(respBody))
	}

	// Deserialize the response if a destination was provided
	if result != nil {
		if err := json.Unmarshal(respBody, result); err != nil {
			return fmt.Errorf("error deserializing response: %w", err)
		}
	}

	return nil
}
