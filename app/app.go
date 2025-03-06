package app

import (
	"log"

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
	extractorResult, err := ex.Extract(&extractor.ExtractorParams{})
	if err != nil {
		return nil, err
	}

	log.Println(extractorResult)

	return &AppResult{}, nil
}
