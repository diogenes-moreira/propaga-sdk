package models

// CornerStore represents a corner store in the Propaga system
type CornerStore struct {
	ID          string                 `json:"id"`
	Name        string                 `json:"name"`
	Address     string                 `json:"address"`
	City        string                 `json:"city"`
	State       string                 `json:"state"`
	PostalCode  string                 `json:"postal_code"`
	Country     string                 `json:"country"`
	PhoneNumber string                 `json:"phone_number,omitempty"`
	Email       string                 `json:"email,omitempty"`
	Status      string                 `json:"status"`
	CreatedAt   string                 `json:"created_at"`
	UpdatedAt   string                 `json:"updated_at"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

// CornerStoreStatus represents the possible states of a corner store
const (
	CornerStoreStatusActive   = "active"
	CornerStoreStatusInactive = "inactive"
	CornerStoreStatusPending  = "pending"
)

// CornerStoreListParams represents the parameters for listing corner stores
type CornerStoreListParams struct {
	Limit     int    `json:"limit,omitempty"`
	Offset    int    `json:"offset,omitempty"`
	Status    string `json:"status,omitempty"`
	City      string `json:"city,omitempty"`
	State     string `json:"state,omitempty"`
	StartDate string `json:"start_date,omitempty"`
	EndDate   string `json:"end_date,omitempty"`
}

// CornerStoreCreateParams represents the parameters for creating a corner store
type CornerStoreCreateParams struct {
	Name        string                 `json:"name"`
	Address     string                 `json:"address"`
	City        string                 `json:"city"`
	State       string                 `json:"state"`
	PostalCode  string                 `json:"postal_code"`
	Country     string                 `json:"country"`
	PhoneNumber string                 `json:"phone_number,omitempty"`
	Email       string                 `json:"email,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

// CornerStoreUpdateParams represents the parameters for updating a corner store
type CornerStoreUpdateParams struct {
	Name        string                 `json:"name,omitempty"`
	Address     string                 `json:"address,omitempty"`
	City        string                 `json:"city,omitempty"`
	State       string                 `json:"state,omitempty"`
	PostalCode  string                 `json:"postal_code,omitempty"`
	Country     string                 `json:"country,omitempty"`
	PhoneNumber string                 `json:"phone_number,omitempty"`
	Email       string                 `json:"email,omitempty"`
	Status      string                 `json:"status,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

// CornerStoreListResponse represents the response when listing corner stores
type CornerStoreListResponse struct {
	Data       []CornerStore `json:"data"`
	TotalCount int           `json:"total_count"`
	Limit      int           `json:"limit"`
	Offset     int           `json:"offset"`
}
