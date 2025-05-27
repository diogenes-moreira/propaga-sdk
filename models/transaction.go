package models

// Transaction represents a transaction in the Propaga system
type Transaction struct {
	ID            string                 `json:"id"`
	Amount        float64                `json:"amount"`
	Currency      string                 `json:"currency"`
	Status        string                 `json:"status"`
	CustomerID    string                 `json:"customer_id"`
	CreatedAt     string                 `json:"created_at"`
	UpdatedAt     string                 `json:"updated_at"`
	Description   string                 `json:"description,omitempty"`
	PaymentMethod string                 `json:"payment_method,omitempty"`
	Metadata      map[string]interface{} `json:"metadata,omitempty"`
}

// TransactionStatus represents the possible states of a transaction
const (
	TransactionStatusPending   = "pending"
	TransactionStatusApproved  = "approved"
	TransactionStatusRejected  = "rejected"
	TransactionStatusCancelled = "cancelled"
	TransactionStatusCompleted = "completed"
)

// TransactionListParams represents the parameters for listing transactions
type TransactionListParams struct {
	Limit      int    `json:"limit,omitempty"`
	Offset     int    `json:"offset,omitempty"`
	CustomerID string `json:"customer_id,omitempty"`
	Status     string `json:"status,omitempty"`
	StartDate  string `json:"start_date,omitempty"`
	EndDate    string `json:"end_date,omitempty"`
}

// TransactionCreateParams represents the parameters for creating a transaction
type TransactionCreateParams struct {
	Amount        float64                `json:"amount"`
	Currency      string                 `json:"currency"`
	CustomerID    string                 `json:"customer_id"`
	Description   string                 `json:"description,omitempty"`
	PaymentMethod string                 `json:"payment_method,omitempty"`
	Metadata      map[string]interface{} `json:"metadata,omitempty"`
}

// TransactionUpdateParams represents the parameters for updating a transaction
type TransactionUpdateParams struct {
	Status      string                 `json:"status,omitempty"`
	Description string                 `json:"description,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

// TransactionListResponse represents the response when listing transactions
type TransactionListResponse struct {
	Data       []Transaction `json:"data"`
	TotalCount int           `json:"total_count"`
	Limit      int           `json:"limit"`
	Offset     int           `json:"offset"`
}

// APIError represents an error returned by the API
type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}
