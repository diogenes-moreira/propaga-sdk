package models

// Account represents a user account in the Propaga system
type Account struct {
	ID             string                 `json:"id"`
	CustomerID     string                 `json:"customer_id"`
	Name           string                 `json:"name"`
	Email          string                 `json:"email,omitempty"`
	PhoneNumber    string                 `json:"phone_number"`
	Status         string                 `json:"status"`
	CreditLimit    float64                `json:"credit_limit"`
	CurrentBalance float64                `json:"current_balance"`
	CreatedAt      string                 `json:"created_at"`
	UpdatedAt      string                 `json:"updated_at"`
	Metadata       map[string]interface{} `json:"metadata,omitempty"`
}

// AccountStatus represents the possible states of an account
const (
	AccountStatusActive    = "active"
	AccountStatusInactive  = "inactive"
	AccountStatusSuspended = "suspended"
	AccountStatusPending   = "pending"
)

// AccountListParams represents the parameters for listing accounts
type AccountListParams struct {
	Limit      int    `json:"limit,omitempty"`
	Offset     int    `json:"offset,omitempty"`
	CustomerID string `json:"customer_id,omitempty"`
	Status     string `json:"status,omitempty"`
	StartDate  string `json:"start_date,omitempty"`
	EndDate    string `json:"end_date,omitempty"`
}

// AccountCreateParams represents the parameters for creating an account
type AccountCreateParams struct {
	CustomerID  string                 `json:"customer_id"`
	Name        string                 `json:"name"`
	Email       string                 `json:"email,omitempty"`
	PhoneNumber string                 `json:"phone_number"`
	CreditLimit float64                `json:"credit_limit,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

// AccountUpdateParams represents the parameters for updating an account
type AccountUpdateParams struct {
	Name        string                 `json:"name,omitempty"`
	Email       string                 `json:"email,omitempty"`
	PhoneNumber string                 `json:"phone_number,omitempty"`
	Status      string                 `json:"status,omitempty"`
	CreditLimit float64                `json:"credit_limit,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

// AccountListResponse represents the response when listing accounts
type AccountListResponse struct {
	Data       []Account `json:"data"`
	TotalCount int       `json:"total_count"`
	Limit      int       `json:"limit"`
	Offset     int       `json:"offset"`
}
