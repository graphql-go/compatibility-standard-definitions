package validator

import "graphql-go/compatibility-standard-definitions/types"

// Validator represents the component that validates standard definitions.
type Validator struct {
}

// ValidateParams represents the parameters for the validate method.
type ValidateParams struct {
	Specification  types.SpecificationIntrospection
	Implementation types.ImplementationIntrospection
}

// ValidateResult represents the result of the validate method.
type ValidateResult struct {
}

// Validates validates given graphql introspection query results.
func (v *Validator) Validate(params *ValidateParams) (*ValidateResult, error) {
	return nil, nil
}
