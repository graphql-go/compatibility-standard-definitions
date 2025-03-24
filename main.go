package main

import (
	"log"

	mainApp "graphql-go/compatibility-standard-definitions/app"
	"graphql-go/compatibility-standard-definitions/cmd"
	"graphql-go/compatibility-standard-definitions/config"
)

func main() {
	cfg := config.New()

	cli := cmd.CLI{}
	if _, err := cli.Run(&cmd.RunParams{
		Choices: cfg.AvailableImplementations,
		Header:  cfg.GraphqlSpecificationWithPrefix,
	}); err != nil {
		log.Fatal(err)
	}

	app := mainApp.App{}

	runResult, err := app.Run(mainApp.RunParams{
		Specification:  cfg.GraphqlSpecification,
		Implementation: cfg.GraphqlGoImplementation,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(runResult.Status)

	if runResult.Details != "" && !cfg.IsDebug {
		log.Println(runResult.Details)
	}
}
