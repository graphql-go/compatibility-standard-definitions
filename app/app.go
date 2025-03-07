package app

import (
	"graphql-go/compatibility-standard-definitions/extractor"
	"graphql-go/compatibility-standard-definitions/puller"
	"graphql-go/compatibility-standard-definitions/types"
)

type App struct {
}

type AppResult struct {
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
	if _, err := ex.Extract(&extractor.ExtractorParams{}); err != nil {
		return nil, err
	}

	return &AppResult{}, nil
}
