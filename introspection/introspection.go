// Package introspection provides different types for interacting with GraphQL introspection operations.
package introspection

import "graphql-go/compatibility-standard-definitions/types"

// queryResultFilePath is the file path of the introspection result against the graphql javascript implementation.
const queryResultFilePath string = "./graphql-js-introspection/introspectionQueryResult.json"

// Introspection represents a wrapper for operations related to GraphQL introspection.
type Introspection struct {
}

// NewIntrospection returns a pointer to the Introspection struct.
func NewIntrospection() *Introspection {
	return &Introspection{}
}

// SpecificationQuery maps `queryResultFilePath` to `types.IntrospectionQueryResult`.
func (i *Introspection) SpecificationQuery() (*types.IntrospectionQueryResult, error) {
	result := &types.IntrospectionQueryResult{}
	// TODO(@mentatbot): Map the `queryResultFilePath` json contents to `result` matching the exact fields.
	return result, nil
}
