package account

import (
	"fmt"
	"net/http"

	"github.com/diogenes-moreira/propaga-sdk/client"
	"github.com/diogenes-moreira/propaga-sdk/models"
)

// Service provides methods for interacting with accounts in the Propaga API
type Service struct {
	client *client.Client
}

// NewService creates a new instance of the account service
func NewService(client *client.Client) *Service {
	return &Service{
		client: client,
	}
}

// List retrieves a list of accounts based on the specified parameters
func (s *Service) List(params *models.AccountListParams) (*models.AccountListResponse, error) {
	result := &models.AccountListResponse{}

	// Endpoint placeholder - should be updated when documentation is available
	err := s.client.DoRequest(http.MethodGet, "/v1/accounts", params, result)
	if err != nil {
		return nil, fmt.Errorf("error listing accounts: %w", err)
	}

	return result, nil
}

// Get retrieves a specific account by its ID
func (s *Service) Get(id string) (*models.Account, error) {
	result := &models.Account{}

	// Endpoint placeholder - should be updated when documentation is available
	path := fmt.Sprintf("/v1/accounts/%s", id)
	err := s.client.DoRequest(http.MethodGet, path, nil, result)
	if err != nil {
		return nil, fmt.Errorf("error getting account %s: %w", id, err)
	}

	return result, nil
}

// Create creates a new account
func (s *Service) Create(params *models.AccountCreateParams) (*models.Account, error) {
	result := &models.Account{}

	// Endpoint placeholder - should be updated when documentation is available
	err := s.client.DoRequest(http.MethodPost, "/v1/accounts", params, result)
	if err != nil {
		return nil, fmt.Errorf("error creating account: %w", err)
	}

	return result, nil
}

// Update updates an existing account
func (s *Service) Update(id string, params *models.AccountUpdateParams) (*models.Account, error) {
	result := &models.Account{}

	// Endpoint placeholder - should be updated when documentation is available
	path := fmt.Sprintf("/v1/accounts/%s", id)
	err := s.client.DoRequest(http.MethodPut, path, params, result)
	if err != nil {
		return nil, fmt.Errorf("error updating account %s: %w", id, err)
	}

	return result, nil
}

// Suspend suspends an account
func (s *Service) Suspend(id string) (*models.Account, error) {
	result := &models.Account{}

	// Endpoint placeholder - should be updated when documentation is available
	path := fmt.Sprintf("/v1/accounts/%s/suspend", id)
	err := s.client.DoRequest(http.MethodPost, path, nil, result)
	if err != nil {
		return nil, fmt.Errorf("error suspending account %s: %w", id, err)
	}

	return result, nil
}

// Activate activates a suspended account
func (s *Service) Activate(id string) (*models.Account, error) {
	result := &models.Account{}

	// Endpoint placeholder - should be updated when documentation is available
	path := fmt.Sprintf("/v1/accounts/%s/activate", id)
	err := s.client.DoRequest(http.MethodPost, path, nil, result)
	if err != nil {
		return nil, fmt.Errorf("error activating account %s: %w", id, err)
	}

	return result, nil
}
