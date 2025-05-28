package models

// KYC represents a Know Your Customer verification in the Propaga system
type KYC struct {
	ID           string                 `json:"id"`
	CustomerID   string                 `json:"customer_id"`
	Status       string                 `json:"status"`
	DocumentType string                 `json:"document_type"`
	DocumentID   string                 `json:"document_id"`
	FullName     string                 `json:"full_name"`
	DateOfBirth  string                 `json:"date_of_birth,omitempty"`
	Address      string                 `json:"address,omitempty"`
	CreatedAt    string                 `json:"created_at"`
	UpdatedAt    string                 `json:"updated_at"`
	VerifiedAt   string                 `json:"verified_at,omitempty"`
	RejectedAt   string                 `json:"rejected_at,omitempty"`
	RejectReason string                 `json:"reject_reason,omitempty"`
	Metadata     map[string]interface{} `json:"metadata,omitempty"`
}

// KYCStatus represents the possible states of a KYC verification
const (
	KYCStatusPending  = "pending"
	KYCStatusVerified = "verified"
	KYCStatusRejected = "rejected"
	KYCStatusExpired  = "expired"
)

// KYCDocumentType represents the possible document types for KYC verification
const (
	KYCDocumentTypeID        = "id_card"
	KYCDocumentTypePassport  = "passport"
	KYCDocumentTypeDriverLic = "drivers_license"
)

// KYCListParams represents the parameters for listing KYC verifications
type KYCListParams struct {
	Limit      int    `json:"limit,omitempty"`
	Offset     int    `json:"offset,omitempty"`
	CustomerID string `json:"customer_id,omitempty"`
	Status     string `json:"status,omitempty"`
	StartDate  string `json:"start_date,omitempty"`
	EndDate    string `json:"end_date,omitempty"`
}

// KYCCreateParams represents the parameters for creating a KYC verification
type KYCCreateParams struct {
	CustomerID   string                 `json:"customer_id"`
	DocumentType string                 `json:"document_type"`
	DocumentID   string                 `json:"document_id"`
	FullName     string                 `json:"full_name"`
	DateOfBirth  string                 `json:"date_of_birth,omitempty"`
	Address      string                 `json:"address,omitempty"`
	Metadata     map[string]interface{} `json:"metadata,omitempty"`
}

// KYCUpdateParams represents the parameters for updating a KYC verification
type KYCUpdateParams struct {
	Status       string                 `json:"status,omitempty"`
	DocumentType string                 `json:"document_type,omitempty"`
	DocumentID   string                 `json:"document_id,omitempty"`
	FullName     string                 `json:"full_name,omitempty"`
	DateOfBirth  string                 `json:"date_of_birth,omitempty"`
	Address      string                 `json:"address,omitempty"`
	RejectReason string                 `json:"reject_reason,omitempty"`
	Metadata     map[string]interface{} `json:"metadata,omitempty"`
}

// KYCListResponse represents the response when listing KYC verifications
type KYCListResponse struct {
	Data       []KYC `json:"data"`
	TotalCount int   `json:"total_count"`
	Limit      int   `json:"limit"`
	Offset     int   `json:"offset"`
}
