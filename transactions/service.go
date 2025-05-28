package transactions

import (
	"fmt"
	"net/http"

	"github.com/diogenes-moreira/propaga-sdk/client"
	"github.com/diogenes-moreira/propaga-sdk/models"
)

// Service provides methods for interacting with transactions in the Propaga API
type Service struct {
	client *client.Client
}

// NewService creates a new instance of the transactions service
func NewService(client *client.Client) *Service {
	return &Service{
		client: client,
	}
}

// List retrieves a list of transactions based on the specified parameters
func (s *Service) List(params *models.TransactionListParams) (*models.TransactionListResponse, error) {
	result := &models.TransactionListResponse{}

	// Endpoint placeholder - should be updated when documentation is available
	err := s.client.DoRequest(http.MethodGet, "/v1/transaction", params, result)
	if err != nil {
		return nil, fmt.Errorf("error listing transactions: %w", err)
	}

	return result, nil
}

// Get retrieves a specific transaction by its ID
func (s *Service) Get(id string) (*models.Transaction, error) {
	result := &models.Transaction{}

	// Endpoint placeholder - should be updated when documentation is available
	path := fmt.Sprintf("/v1/transaction/%s", id)
	err := s.client.DoRequest(http.MethodGet, path, nil, result)
	if err != nil {
		return nil, fmt.Errorf("error getting transaction %s: %w", id, err)
	}

	return result, nil
}

// GetByExternalID retrieves a transaction by its external ID
func (s *Service) GetByExternalID(externalID string) (*models.Transaction, error) {
	result := &models.Transaction{}

	// Endpoint placeholder - should be updated when documentation is available
	path := fmt.Sprintf("/v1/transaction/external/%s", externalID)
	err := s.client.DoRequest(http.MethodGet, path, nil, result)
	if err != nil {
		return nil, fmt.Errorf("error getting transaction by external ID %s: %w", externalID, err)
	}

	return result, nil
}

// Create creates a new transaction
func (s *Service) Create(params *models.TransactionCreateParams) (*models.Transaction, error) {
	result := &models.Transaction{}

	// Endpoint placeholder - should be updated when documentation is available
	err := s.client.DoRequest(http.MethodPost, "/v1/transaction", params, result)
	if err != nil {
		return nil, fmt.Errorf("error creating transaction: %w", err)
	}

	return result, nil
}

// Update updates an existing transaction
func (s *Service) Update(id string, params *models.TransactionUpdateParams) (*models.Transaction, error) {
	result := &models.Transaction{}

	// Endpoint placeholder - should be updated when documentation is available
	path := fmt.Sprintf("/v1/transaction/%s", id)
	err := s.client.DoRequest(http.MethodPut, path, params, result)
	if err != nil {
		return nil, fmt.Errorf("error updating transaction %s: %w", id, err)
	}

	return result, nil
}

// Cancel cancels an existing transaction
func (s *Service) Cancel(id string) (*models.Transaction, error) {
	result := &models.Transaction{}

	// Endpoint placeholder - should be updated when documentation is available
	path := fmt.Sprintf("/v1/transaction/%s/cancel", id)
	err := s.client.DoRequest(http.MethodPost, path, nil, result)
	if err != nil {
		return nil, fmt.Errorf("error canceling transaction %s: %w", id, err)
	}

	return result, nil
}
