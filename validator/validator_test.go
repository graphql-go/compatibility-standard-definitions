package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"graphql-go/compatibility-standard-definitions/types"
)

func TestValidate_IgnoresSliceOrdering(t *testing.T) {
	validator := &Validator{}

	// Create two introspection results with same content but different ordering
	spec := &types.SpecificationIntrospection{
		QueryResult: types.IntrospectionQueryResult{
			Schema: types.IntrospectionSchema{
				Types: []types.IntrospectionFullType{
					{
						Kind: "OBJECT",
						Name: stringPtr("TestType"),
						Fields: []types.IntrospectionField{
							{Name: "fieldA", Description: stringPtr("Description A")},
							{Name: "fieldB", Description: stringPtr("Description B")},
						},
					},
				},
				Directives: []types.IntrospectionDirective{
					{Name: "directiveA", Description: stringPtr("Directive A")},
					{Name: "directiveB", Description: stringPtr("Directive B")},
				},
			},
		},
	}

	impl := &types.ImplementationIntrospection{
		QueryResult: types.IntrospectionQueryResult{
			Schema: types.IntrospectionSchema{
				Types: []types.IntrospectionFullType{
					{
						Kind: "OBJECT",
						Name: stringPtr("TestType"),
						Fields: []types.IntrospectionField{
							// Different order but same content
							{Name: "fieldB", Description: stringPtr("Description B")},
							{Name: "fieldA", Description: stringPtr("Description A")},
						},
					},
				},
				Directives: []types.IntrospectionDirective{
					// Different order but same content
					{Name: "directiveB", Description: stringPtr("Directive B")},
					{Name: "directiveA", Description: stringPtr("Directive A")},
				},
			},
		},
	}

	result, err := validator.Validate(&ValidateParams{
		Specification:  spec,
		Implementation: impl,
	})

	assert.NoError(t, err)
	assert.Equal(t, Success, result.Result)
	assert.Empty(t, result.Difference, "Should have no differences when slice ordering differs but content is the same")
}

func TestValidate_DetectsDifferences(t *testing.T) {
	validator := &Validator{}

	spec := &types.SpecificationIntrospection{
		QueryResult: types.IntrospectionQueryResult{
			Schema: types.IntrospectionSchema{
				Types: []types.IntrospectionFullType{
					{
						Kind: "OBJECT",
						Name: stringPtr("TestType"),
						Fields: []types.IntrospectionField{
							{Name: "fieldA", Description: stringPtr("Description A")},
						},
					},
				},
			},
		},
	}

	impl := &types.ImplementationIntrospection{
		QueryResult: types.IntrospectionQueryResult{
			Schema: types.IntrospectionSchema{
				Types: []types.IntrospectionFullType{
					{
						Kind: "OBJECT",
						Name: stringPtr("TestType"),
						Fields: []types.IntrospectionField{
							{Name: "fieldA", Description: stringPtr("Different Description")},
						},
					},
				},
			},
		},
	}

	result, err := validator.Validate(&ValidateParams{
		Specification:  spec,
		Implementation: impl,
	})

	assert.NoError(t, err)
	assert.Equal(t, Failure, result.Result)
	assert.NotEmpty(t, result.Difference, "Should detect actual content differences")
}

func stringPtr(s string) *string {
	return &s
}
