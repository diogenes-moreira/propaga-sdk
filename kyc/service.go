package kyc

import (
	"fmt"
	"net/http"

	"github.com/diogenes-moreira/propaga-sdk/client"
	"github.com/diogenes-moreira/propaga-sdk/models"
)

// Service provides methods for interacting with KYC verifications in the Propaga API
type Service struct {
	client *client.Client
}

// NewService creates a new instance of the KYC service
func NewService(client *client.Client) *Service {
	return &Service{
		client: client,
	}
}

// List retrieves a list of KYC verifications based on the specified parameters
func (s *Service) List(params *models.KYCListParams) (*models.KYCListResponse, error) {
	result := &models.KYCListResponse{}

	// Endpoint placeholder - should be updated when documentation is available
	err := s.client.DoRequest(http.MethodGet, "/v1/kyc", params, result)
	if err != nil {
		return nil, fmt.Errorf("error listing KYC verifications: %w", err)
	}

	return result, nil
}

// Get retrieves a specific KYC verification by its ID
func (s *Service) Get(id string) (*models.KYC, error) {
	result := &models.KYC{}

	// Endpoint placeholder - should be updated when documentation is available
	path := fmt.Sprintf("/v1/kyc/%s", id)
	err := s.client.DoRequest(http.MethodGet, path, nil, result)
	if err != nil {
		return nil, fmt.Errorf("error getting KYC verification %s: %w", id, err)
	}

	return result, nil
}

// Create creates a new KYC verification
func (s *Service) Create(params *models.KYCCreateParams) (*models.KYC, error) {
	result := &models.KYC{}

	// Endpoint placeholder - should be updated when documentation is available
	err := s.client.DoRequest(http.MethodPost, "/v1/kyc", params, result)
	if err != nil {
		return nil, fmt.Errorf("error creating KYC verification: %w", err)
	}

	return result, nil
}

// Update updates an existing KYC verification
func (s *Service) Update(id string, params *models.KYCUpdateParams) (*models.KYC, error) {
	result := &models.KYC{}

	// Endpoint placeholder - should be updated when documentation is available
	path := fmt.Sprintf("/v1/kyc/%s", id)
	err := s.client.DoRequest(http.MethodPut, path, params, result)
	if err != nil {
		return nil, fmt.Errorf("error updating KYC verification %s: %w", id, err)
	}

	return result, nil
}

// Verify marks a KYC verification as verified
func (s *Service) Verify(id string) (*models.KYC, error) {
	result := &models.KYC{}

	// Endpoint placeholder - should be updated when documentation is available
	path := fmt.Sprintf("/v1/kyc/%s/verify", id)
	err := s.client.DoRequest(http.MethodPost, path, nil, result)
	if err != nil {
		return nil, fmt.Errorf("error verifying KYC %s: %w", id, err)
	}

	return result, nil
}

// Reject rejects a KYC verification with a reason
func (s *Service) Reject(id string, reason string) (*models.KYC, error) {
	result := &models.KYC{}

	// Endpoint placeholder - should be updated when documentation is available
	path := fmt.Sprintf("/v1/kyc/%s/reject", id)
	payload := map[string]string{"reason": reason}
	err := s.client.DoRequest(http.MethodPost, path, payload, result)
	if err != nil {
		return nil, fmt.Errorf("error rejecting KYC %s: %w", id, err)
	}

	return result, nil
}
