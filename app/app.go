package app

import (
	"graphql-go/compatibility-standard-definitions/executor"
	"graphql-go/compatibility-standard-definitions/extractor"
	"graphql-go/compatibility-standard-definitions/puller"
	"graphql-go/compatibility-standard-definitions/types"
	"graphql-go/compatibility-standard-definitions/validator"
)

type App struct {
}

type RunResult struct {
	Status  string
	Details string
}

type RunParams struct {
	Specification  types.Specification
	Implementation types.Implementation
}

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
