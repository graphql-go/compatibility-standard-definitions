package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"graphql-go/compatibility-standard-definitions/types"
)

func TestValidate_HandlesKnownDifferences(t *testing.T) {
	validator := &Validator{}

	// Test the specific differences mentioned in the GitHub issue
	testCases := []struct {
		name            string
		specDesc        *string
		implDesc        *string
		shouldBeEqual   bool
		description     string
	}{
		{
			name:            "nil vs empty string",
			specDesc:        nil,
			implDesc:        stringPtr(""),
			shouldBeEqual:   true,
			description:     "Empty strings should be normalized to nil",
		},
		{
			name:            "missing period",
			specDesc:        stringPtr("An enum describing what kind of type a given `__Type` is."),
			implDesc:        stringPtr("An enum describing what kind of type a given `__Type` is"),
			shouldBeEqual:   true,
			description:     "Missing period should be normalized",
		},
		{
			name:            "extra space before newline",
			specDesc:        stringPtr("A Directive provides a way to describe alternate runtime execution and type validation behavior in a GraphQL document.\n\nIn some cases, you need to provide options to alter GraphQL's execution behavior in ways field arguments will not suffice, such as conditionally including or skipping a field. Directives provide this by describing additional information to the executor."),
			implDesc:        stringPtr("A Directive provides a way to describe alternate runtime execution and type validation behavior in a GraphQL document. \n\nIn some cases, you need to provide options to alter GraphQL's execution behavior in ways field arguments will not suffice, such as conditionally including or skipping a field. Directives provide this by describing additional information to the executor."),
			shouldBeEqual:   true,
			description:     "Extra space before newline should be normalized",
		},
		{
			name:            "unicode quote normalization",
			specDesc:        stringPtr("GraphQL's execution behavior"),
			implDesc:        stringPtr("GraphQL's execution behavior"),
			shouldBeEqual:   true,
			description:     "Unicode quotes should be normalized to ASCII",
		},
		{
			name:            "object definition text fix",
			specDesc:        stringPtr("Location adjacent to an object type definition."),
			implDesc:        stringPtr("Location adjacent to a object definition."),
			shouldBeEqual:   true,
			description:     "Grammar fix for object definition",
		},
		{
			name:            "subscription text fix",
			specDesc:        stringPtr("If this server support subscription, the type that subscription operations will be rooted at."),
			implDesc:        stringPtr("If this server supports subscription, the type that subscription operations will be rooted at."),
			shouldBeEqual:   true,
			description:     "Grammar fix for subscription text",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			spec := &types.SpecificationIntrospection{
				QueryResult: types.IntrospectionQueryResult{
					Schema: types.IntrospectionSchema{
						Types: []types.IntrospectionFullType{
							{
								Kind:        "OBJECT",
								Name:        stringPtr("TestType"),
								Description: tc.specDesc,
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
								Kind:        "OBJECT",
								Name:        stringPtr("TestType"),
								Description: tc.implDesc,
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
			
			if tc.shouldBeEqual {
				assert.Equal(t, Success, result.Result, tc.description)
				assert.Empty(t, result.Difference, "Should have no differences: %s", tc.description)
			} else {
				assert.Equal(t, Failure, result.Result, tc.description)
				assert.NotEmpty(t, result.Difference, "Should have differences: %s", tc.description)
			}
		})
	}
}

func TestNormalizeStringFields(t *testing.T) {
	testCases := []struct {
		name     string
		input    *string
		expected *string
	}{
		{
			name:     "nil input",
			input:    nil,
			expected: nil,
		},
		{
			name:     "empty string",
			input:    stringPtr(""),
			expected: nil,
		},
		{
			name:     "unicode quote normalization",
			input:    stringPtr("GraphQL's execution"),
			expected: stringPtr("GraphQL's execution"),
		},
		{
			name:     "space before newline",
			input:    stringPtr("text. \nmore text"),
			expected: stringPtr("text.\nmore text"),
		},
		{
			name:     "TypeKind period fix",
			input:    stringPtr("An enum describing what kind of type a given `__Type` is"),
			expected: stringPtr("An enum describing what kind of type a given `__Type` is."),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := normalizeStringFields(tc.input)
			
			if tc.expected == nil {
				assert.Nil(t, result)
			} else {
				assert.NotNil(t, result)
				assert.Equal(t, *tc.expected, *result)
			}
		})
	}
}

