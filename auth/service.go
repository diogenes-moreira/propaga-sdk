package auth

import (
	"github.com/diogenes-moreira/propaga-sdk/client"
)

// Service provides methods for authentication with the Propaga API
type Service struct {
	client *client.Client
}

// NewService creates a new instance of the authentication service
func NewService(client *client.Client) *Service {
	return &Service{
		client: client,
	}
}

// ValidateToken validates that the API token is valid
// Note: According to the documentation, Propaga tokens never expire
func (s *Service) ValidateToken() (bool, error) {
	// Since tokens never expire according to the documentation,
	// this function simply returns true if the client has a token configured
	return s.client.APIKey != "", nil
}
