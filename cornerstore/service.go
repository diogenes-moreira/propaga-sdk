package cornerstore

import (
	"fmt"
	"net/http"

	"github.com/diogenes-moreira/propaga-sdk/client"
	"github.com/diogenes-moreira/propaga-sdk/models"
)

// Service provides methods for interacting with corner stores in the Propaga API
type Service struct {
	client *client.Client
}

// NewService creates a new instance of the corner store service
func NewService(client *client.Client) *Service {
	return &Service{
		client: client,
	}
}

// List retrieves a list of corner stores based on the specified parameters
func (s *Service) List(params *models.CornerStoreListParams) (*models.CornerStoreListResponse, error) {
	result := &models.CornerStoreListResponse{}

	// Endpoint placeholder - should be updated when documentation is available
	err := s.client.DoRequest(http.MethodGet, "/v1/corner-stores", params, result)
	if err != nil {
		return nil, fmt.Errorf("error listing corner stores: %w", err)
	}

	return result, nil
}

// Get retrieves a specific corner store by its ID
func (s *Service) Get(id string) (*models.CornerStore, error) {
	result := &models.CornerStore{}

	// Endpoint placeholder - should be updated when documentation is available
	path := fmt.Sprintf("/v1/corner-stores/%s", id)
	err := s.client.DoRequest(http.MethodGet, path, nil, result)
	if err != nil {
		return nil, fmt.Errorf("error getting corner store %s: %w", id, err)
	}

	return result, nil
}

// Create creates a new corner store
func (s *Service) Create(params *models.CornerStoreCreateParams) (*models.CornerStore, error) {
	result := &models.CornerStore{}

	// Endpoint placeholder - should be updated when documentation is available
	err := s.client.DoRequest(http.MethodPost, "/v1/corner-stores", params, result)
	if err != nil {
		return nil, fmt.Errorf("error creating corner store: %w", err)
	}

	return result, nil
}

// Update updates an existing corner store
func (s *Service) Update(id string, params *models.CornerStoreUpdateParams) (*models.CornerStore, error) {
	result := &models.CornerStore{}

	// Endpoint placeholder - should be updated when documentation is available
	path := fmt.Sprintf("/v1/corner-stores/%s", id)
	err := s.client.DoRequest(http.MethodPut, path, params, result)
	if err != nil {
		return nil, fmt.Errorf("error updating corner store %s: %w", id, err)
	}

	return result, nil
}

// Delete deletes a corner store
func (s *Service) Delete(id string) error {
	// Endpoint placeholder - should be updated when documentation is available
	path := fmt.Sprintf("/v1/corner-stores/%s", id)
	err := s.client.DoRequest(http.MethodDelete, path, nil, nil)
	if err != nil {
		return fmt.Errorf("error deleting corner store %s: %w", id, err)
	}

	return nil
}
