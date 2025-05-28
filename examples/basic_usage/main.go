// Package main demonstrates how to use the Propaga SDK for managing transactions.
// It includes examples of listing transactions, retrieving specific transaction details,
// creating, updating, and canceling transactions.
package main

import (
	"fmt"
	"log"

	"github.com/diogenes-moreira/propaga-sdk"
	"github.com/diogenes-moreira/propaga-sdk/models"
)

func main() {
	// NewClient creates and returns a new Propaga SDK client.
	// It requires an API key to authenticate API requests.
	client := propaga.NewClient("your_api_key_here", false) // Set to true for staging environment

	// Example 1: List transactions
	// This example retrieves a list of transactions using the List method.
	// It demonstrates filtering transactions by status, as well as retrieving paginated results.
	fmt.Println("Example 1: List transactions")
	params := &models.TransactionListParams{
		Limit:  10,                              // Number of transactions to retrieve per page
		Offset: 0,                               // Starting point for pagination
		Status: models.TransactionStatusPending, // Filter results for pending transactions
	}

	// Transactions.List fetches a list of transactions matching the provided parameters.
	transactions, err := client.Transactions.List(params)
	if err != nil {
		log.Fatalf("Error listing transactions: %v", err)
	}

	// Output the transactions
	fmt.Printf("Found %d transactions\n", transactions.TotalCount)
	for i, tx := range transactions.Data {
		fmt.Printf("%d. ID: %s, Amount: %.2f %s, Status: %s\n",
			i+1, tx.ID, tx.Amount, tx.Currency, tx.Status)
	}

	// Example 2: Get a specific transaction
	// This example demonstrates how to retrieve detailed information for a specific transaction
	// using its unique identifier (ID).
	fmt.Println("\nExample 2: Get a specific transaction")
	if len(transactions.Data) > 0 {
		// Retrieve the first Transaction ID from the list
		txID := transactions.Data[0].ID

		// Transactions.Get fetches the details of a transaction by its ID.
		tx, err := client.Transactions.Get(txID)
		if err != nil {
			log.Fatalf("Error getting transaction %s: %v", txID, err)
		}

		// Output details of the retrieved transaction.
		fmt.Printf("Transaction ID: %s\n", tx.ID)
		fmt.Printf("  Amount: %.2f %s\n", tx.Amount, tx.Currency)
		fmt.Printf("  Status: %s\n", tx.Status)
		fmt.Printf("  Customer: %s\n", tx.CustomerID)
		fmt.Printf("  Created: %s\n", tx.CreatedAt)
		fmt.Printf("  Updated: %s\n", tx.UpdatedAt)
	}

	// Example 3: Create a new transaction
	// This example creates a new transaction with various attributes, such as amount,
	// currency, description, and metadata. The Create method is demonstrated here.
	fmt.Println("\nExample 3: Create a new transaction")
	createParams := &models.TransactionCreateParams{
		Amount:        100.50,                       // The amount of the transaction
		Currency:      "MXN",                        // Currency represented in ISO 4217 format
		CustomerID:    "cust_123456",                // Unique identifier for the customer
		Description:   "Payment for invoice #12345", // Details of the transaction
		PaymentMethod: "credit_card",                // Method used for the payment
		Metadata: map[string]interface{}{
			"invoice_id": "inv_12345",
			"product_id": "prod_67890",
		}, // Arbitrary metadata to accompany the transaction
	}

	// Transactions.Create submits the above parameters to create a new transaction.
	newTx, err := client.Transactions.Create(createParams)
	if err != nil {
		log.Fatalf("Error creating transaction: %v", err)
	}

	// Output details of the newly created transaction
	fmt.Printf("New transaction created with ID: %s\n", newTx.ID)
	fmt.Printf("  Amount: %.2f %s\n", newTx.Amount, newTx.Currency)
	fmt.Printf("  Status: %s\n", newTx.Status)

	// Example 4: Update a transaction
	// This example demonstrates how to update attributes of a transaction,
	// such as its description or metadata, using the Update method.
	fmt.Println("\nExample 4: Update a transaction")
	updateParams := &models.TransactionUpdateParams{
		Description: "Payment for invoice #12345 - Updated", // Updated transaction description
		Metadata: map[string]interface{}{
			"updated_by": "user_admin",             // Metadata indicating who updated the transaction
			"reason":     "description correction", // Reason for the update
		},
	}

	// Transactions.Update applies the changes to an existing transaction.
	updatedTx, err := client.Transactions.Update(newTx.ID, updateParams)
	if err != nil {
		log.Fatalf("Error updating transaction: %v", err)
	}

	// Output the updated transaction details
	fmt.Printf("Updated transaction ID: %s\n", updatedTx.ID)
	fmt.Printf("  Description: %s\n", updatedTx.Description)
	fmt.Printf("  Status: %s\n", updatedTx.Status)

	// Example 5: Cancel a transaction
	// This example demonstrates how to cancel an existing transaction using the Cancel method.
	// Cancelling a transaction changes its status to a canceled state.
	fmt.Println("\nExample 5: Cancel a transaction")

	// Transactions.Cancel invalidates an existing transaction by marking it as canceled.
	canceledTx, err := client.Transactions.Cancel(newTx.ID)
	if err != nil {
		log.Fatalf("Error canceling transaction: %v", err)
	}

	// Output details of the canceled transaction
	fmt.Printf("Canceled transaction ID: %s\n", canceledTx.ID)
	fmt.Printf("  Status: %s\n", canceledTx.Status)
}
