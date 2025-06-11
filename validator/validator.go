package validator

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	"graphql-go/compatibility-standard-definitions/types"
)

// Result represents a result type of the validator component.
type Result bool

const (
	// Success represents a success result.
	Success Result = true

	// Failure represents a failure result.
	Failure Result = false
)

// String returns the string representation of the result.
func (r Result) String() string {
	switch r {
	case Success:
		return "SUCCESS"
	default:
		return "FAILURE"
	}
}

// Validator represents the component that validates standard definitions.
type Validator struct {
}

// ValidateParams represents the parameters for the validate method.
type ValidateParams struct {
	Specification  *types.SpecificationIntrospection
	Implementation *types.ImplementationIntrospection
}

// ValidateResult represents the result of the validate method.
type ValidateResult struct {
	Result     Result
	Difference string
}

// Validate validates given graphql introspection query results.
func (v *Validator) Validate(params *ValidateParams) (*ValidateResult, error) {
	diff := cmp.Diff(params.Specification.QueryResult,
		params.Implementation.QueryResult,
		cmpopts.IgnoreUnexported(types.IntrospectionSchema{}),
		// Sort slices to ignore ordering differences
		cmpopts.SortSlices(func(a, b types.IntrospectionField) bool {
			return a.Name < b.Name
		}),
		cmpopts.SortSlices(func(a, b types.IntrospectionInputValue) bool {
			return a.Name < b.Name
		}),
		cmpopts.SortSlices(func(a, b types.IntrospectionEnumValue) bool {
			return a.Name < b.Name
		}),
		cmpopts.SortSlices(func(a, b types.IntrospectionDirective) bool {
			return a.Name < b.Name
		}),
		cmpopts.SortSlices(func(a, b types.DirectiveLocation) bool {
			return string(a) < string(b)
		}),
		cmpopts.SortSlices(func(a, b types.IntrospectionFullType) bool {
			// Handle nil name pointers
			if a.Name == nil && b.Name == nil {
				return a.Kind < b.Kind
			}
			if a.Name == nil {
				return true
			}
			if b.Name == nil {
				return false
			}
			return *a.Name < *b.Name
		}),
		cmpopts.SortSlices(func(a, b types.IntrospectionTypeRef) bool {
			// Handle nil name pointers
			if a.Name == nil && b.Name == nil {
				return a.Kind < b.Kind
			}
			if a.Name == nil {
				return true
			}
			if b.Name == nil {
				return false
			}
			return *a.Name < *b.Name
		}),
	)

	if diff != "" {
		return &ValidateResult{
			Result:     Failure,
			Difference: diff,
		}, nil
	}

	return &ValidateResult{
		Result:     Success,
		Difference: diff,
	}, nil
}
