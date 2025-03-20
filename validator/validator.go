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
