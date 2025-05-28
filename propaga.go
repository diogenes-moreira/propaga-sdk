package propaga

import (
	"time"

	"github.com/diogenes-moreira/propaga-sdk/auth"
	"github.com/diogenes-moreira/propaga-sdk/client"
	"github.com/diogenes-moreira/propaga-sdk/transactions"
)

// Client is the main client for interacting with the Propaga API
type Client struct {
	// Underlying HTTP client
	httpClient *client.Client

	// Available services
	Auth         *auth.Service
	Transactions *transactions.Service
}

// NewClient creates a new instance of the Propaga client with default configuration
func NewClient(apiKey string, staging bool) *Client {
	httpClient := client.NewClient(apiKey, staging)
	return newClientWithHTTPClient(httpClient)
}

// NewClientWithOptions creates a new instance of the client with custom options
func NewClientWithOptions(apiKey, baseURL string, timeout time.Duration) *Client {
	httpClient := client.NewClientWithOptions(apiKey, baseURL, timeout)
	return newClientWithHTTPClient(httpClient)
}

// Helper function to create a client with an existing HTTP client
func newClientWithHTTPClient(httpClient *client.Client) *Client {
	c := &Client{
		httpClient: httpClient,
	}

	// Initialize services
	c.Auth = auth.NewService(httpClient)
	c.Transactions = transactions.NewService(httpClient)

	return c
}
