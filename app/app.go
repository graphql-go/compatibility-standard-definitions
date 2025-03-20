package app

import (
	"graphql-go/compatibility-standard-definitions/executor"
	"graphql-go/compatibility-standard-definitions/extractor"
	"graphql-go/compatibility-standard-definitions/puller"
	"graphql-go/compatibility-standard-definitions/types"
	"graphql-go/compatibility-standard-definitions/validator"
)

// App represents the high level entry point for the application.
type App struct {
}

// RunResult represents the result of the run method.
type RunResult struct {
	Status  string
	Details string
}

// RunParams represents the params of the run method.
type RunParams struct {
	// Specification is the graphql specification.
	Specification types.Specification

	// Implementation is the graphql implementation.
	Implementation types.Implementation
}

// Run runs and returns the application result.
func (app *App) Run(params RunParams) (*RunResult, error) {
	p := puller.Puller{}

	if _, err := p.Pull(&puller.PullerParams{
		Specification:  params.Specification.Repo,
		Implementation: params.Implementation.Repo,
	}); err != nil {
		return nil, err
	}

	executor := executor.New()

	ex := extractor.New(executor)
	extractResult, err := ex.Extract(&extractor.ExtractorParams{
		Implementation: params.Implementation,
		Specification:  params.Specification,
	})
	if err != nil {
		return nil, err
	}

	val := validator.Validator{}
	validateResult, err := val.Validate(&validator.ValidateParams{
		Specification:  extractResult.SpecificationIntrospection,
		Implementation: extractResult.ImplementationIntrospection,
	})
	if err != nil {
		return nil, err
	}

	return &RunResult{
		Status:  validateResult.Result.String(),
		Details: validateResult.Difference,
	}, nil
}
