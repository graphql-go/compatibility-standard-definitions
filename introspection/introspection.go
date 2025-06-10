// Package introspection provides different types for interacting with GraphQL introspection operations.
package introspection

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"graphql-go/compatibility-standard-definitions/types"
)

// queryResultFilePath is the file path of the introspection result against the graphql javascript implementation.
const queryResultFilePath string = "../graphql-js-introspection/introspectionQueryResult.json"

// Introspection represents a wrapper for operations related to GraphQL introspection.
type Introspection struct {
}

// NewIntrospection returns a pointer to the Introspection struct.
func NewIntrospection() *Introspection {
	return &Introspection{}
}

// SpecificationQuery maps `queryResultFilePath` to `types.IntrospectionQueryResult`.
func (i *Introspection) SpecificationQuery() (*types.IntrospectionQueryResult, error) {
	queryResultFile, err := os.Open(queryResultFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer queryResultFile.Close()

	queryResult, err := io.ReadAll(queryResultFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	result := &types.IntrospectionQueryResult{}
	if err := json.Unmarshal(queryResult, result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return result, nil
}
