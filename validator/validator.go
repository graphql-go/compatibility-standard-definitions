package validator

// Validator represents the component that validates standard definitions.
type Validator struct {
}

// ValidatorParams represents the parameters for the validate method.
type ValidatorParams struct {
}

// ValidatorResult represents the result of the validate method.
type ValidatorResult struct {
}

// Validates validates given graphql introspection query results.
func (v *Validator) Validate(params *ValidatorParams) (*ValidatorResult, error) {
	return nil, nil
}
