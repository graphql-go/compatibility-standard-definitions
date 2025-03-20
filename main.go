package main

import (
	"log"

	mainApp "graphql-go/compatibility-standard-definitions/app"
	"graphql-go/compatibility-standard-definitions/cmd"
	"graphql-go/compatibility-standard-definitions/config"
	"graphql-go/compatibility-standard-definitions/implementation"
	"graphql-go/compatibility-standard-definitions/mapper"
)

func main() {
	header := implementation.GraphqlSpecification.Repo.String(implementation.SpecificationPrefix)

	cli := cmd.CLI{}
	if _, err := cli.Run(&cmd.RunParams{
		Choices: mapper.AvailableImplementations(),
		Header:  header,
	}); err != nil {
		log.Fatal(err)
	}

	app := mainApp.App{}
	runResult, err := app.Run(mainApp.RunParams{
		Specification:  implementation.GraphqlSpecification,
		Implementation: implementation.GraphqlGoImplementation,
	})
	if err != nil {
		log.Fatal(err)
	}

	cfg := config.New()

	log.Println(runResult.Status)

	if runResult.Details != "" && cfg.IsDebug == false {
		log.Println(runResult.Details)
	}
}
