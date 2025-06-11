package introspection

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"graphql-go/compatibility-standard-definitions/types"
)

func TestSpecificationQuery(t *testing.T) {
	introspection := NewIntrospection()

	result, err := introspection.SpecificationQuery()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	// Verify the result is not nil and is of correct type
	assert.NotNil(t, result)
	assert.IsType(t, &types.IntrospectionQueryResult{}, result)
	
	// Verify that the schema was parsed
	assert.NotNil(t, result.Schema)
	
	// Verify that query type exists and has the expected name
	assert.NotNil(t, result.Schema.QueryType)
	assert.Equal(t, "RootQueryType", result.Schema.QueryType.Name)
	
	// Verify that mutation and subscription types are nil as expected
	assert.Nil(t, result.Schema.MutationType)
	assert.Nil(t, result.Schema.SubscriptionType)
	
	// Verify directives were parsed (should contain at least include, skip, deprecated)
	assert.NotNil(t, result.Schema.Directives)
	assert.Greater(t, len(result.Schema.Directives), 0)
	
	// Check for expected directive names
	directiveNames := make([]string, len(result.Schema.Directives))
	for i, directive := range result.Schema.Directives {
		directiveNames[i] = directive.Name
	}
	assert.Contains(t, directiveNames, "include")
	assert.Contains(t, directiveNames, "skip")
	assert.Contains(t, directiveNames, "deprecated")
}
