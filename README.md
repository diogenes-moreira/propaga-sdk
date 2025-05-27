# Propaga SDK for Golang

This SDK provides a Golang interface to interact with the Propaga API, making it easier to integrate Buy Now, Pay Later (BNPL) functionality into Go applications.

## Installation

```bash
go get github.com/diogenes-moreira/propaga-sdk
```

## Basic Usage

```go
package main

import (
    "fmt"
    "github.com/diogenes-moreira/propaga-sdk"
)

func main() {
    // Create a client with your API key
    client := propaga.NewClient("your_api_key_here")
    
    // Now you can use the client to interact with the Propaga API
    // For example, to list transactions:
    transactions, err := client.Transactions.List(nil)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    
    fmt.Printf("Found %d transactions\n", transactions.TotalCount)
}
```

## Features

- Authentication using API token
- CRUD operations for transactions
- Robust error handling
- Support for production and staging environments
- Customizable timeouts and base URLs

## SDK Structure

The SDK is organized into the following packages:

- `client`: Contains the base HTTP client for making requests to the API
- `auth`: Provides authentication functionality
- `models`: Defines the data models used in the API
- `transactions`: Implements transaction-related operations

## Transaction Operations

The SDK supports the following operations for transactions:

- List transactions with filters
- Get details of a specific transaction
- Create new transactions
- Update existing transactions
- Cancel transactions

## Examples

Check the `examples` folder for complete usage examples of the SDK.

## Important Notes

This SDK has been developed based on the available Propaga documentation. The endpoints used are placeholders and should be updated when the complete API documentation becomes available.

## License

MIT
