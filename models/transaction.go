package models

import "time"

// Transaction represents a transaction in the Propaga system
type Transaction struct {
	TransactionId            string    `json:"transactionId"`
	CornerStoreId            string    `json:"cornerStoreId"`
	UserId                   string    `json:"userId"`
	TransactionStatus        string    `json:"transactionStatus"`
	WholesalerTransactionId  string    `json:"wholesalerTransactionId"`
	MovementDate             time.Time `json:"movementDate"`
	TotalAmount              float64   `json:"totalAmount"`
	WholesalerFees           float64   `json:"wholesalerFees"`
	Interests                float64   `json:"interests"`
	IVAAmount                float64   `json:"IVAAmount"`
	TotalAmountWithInterests float64   `json:"totalAmountWithInterests"`
	PaymentDate              time.Time `json:"paymentDate"`
	DeliveryDate             time.Time `json:"deliveryDate"`
	Wholesaler               string    `json:"wholesaler"`
	Products                 []Product `json:"products"`
	Metadata                 Metadata  `json:"metadata,omitempty"`
}

type Product struct {
	Id          string    `json:"id"`
	ExternalSKU string    `json:"externalSKU"`
	Name        string    `json:"name"`
	Quantity    int       `json:"quantity"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type Metadata struct {
	SuccessUrl string `json:"success_url"`
	ErrorUrl   string `json:"error_url"`
}

// TransactionStatus represents the possible states of a transaction
const (
	TransactionStatusPending   = "pending-verification"
	TransactionStatusOnHold    = "on-hold"
	TransactionStatusDelivery  = "delivery"
	TransactionStatusCancelled = "cancel"
	TransactionStatusPaid      = "paid"
	TransactionStatusExpired   = "expired"
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
	CornerStoreId           string                 `json:"cornerStoreId"`
	TotalAmount             float64                `json:"totalAmount"`
	WholesalerTransactionId string                 `json:"wholesalerTransactionId"`
	DeliveryDate            string                 `json:"deliveryDate"`
	Products                []Product              `json:"products"`
	Metadata                map[string]interface{} `json:"metadata,omitempty"`
}

// TransactionUpdateParams represents the parameters for updating a transaction
type TransactionUpdateParams struct {
	Status              string                 `json:"status,omitempty"`
	TransactionAmount   float64                `json:"transactionAmount,omitempty"`
	Latitude            float64                `json:"latitude,omitempty"`
	Longitude           float64                `json:"longitude,omitempty"`
	LocationDescription string                 `json:"locationDescription,omitempty"`
	Metadata            map[string]interface{} `json:"metadata,omitempty"`
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
