package app

import (
	"graphql-go/compatibility-standard-definitions/extractor"
	"graphql-go/compatibility-standard-definitions/puller"
	"graphql-go/compatibility-standard-definitions/types"
	"graphql-go/compatibility-standard-definitions/validator"
)

type App struct {
}

type AppResult struct {
	Status  string
	Details string
}

type AppParams struct {
	Specification  types.Repository
	Implementation types.Repository
}

func (app *App) Run(params AppParams) (*AppResult, error) {
	p := puller.Puller{}

	if _, err := p.Pull(&puller.PullerParams{
		Specification:  params.Specification,
		Implementation: params.Implementation,
	}); err != nil {
		return nil, err
	}

	ex := extractor.Extractor{}
	extractResult, err := ex.Extract(&extractor.ExtractorParams{})
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

	return &AppResult{
		Status:  validateResult.Result.String(),
		Details: validateResult.Difference,
	}, nil
}
