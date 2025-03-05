package main

import (
	"log"

	"graphql-go/compatibility-standard-definitions/cmd"
	"graphql-go/compatibility-standard-definitions/implementation"
)

var choices = []string{}

func init() {
	for _, i := range implementation.Implementations {
		choices = append(choices, i.Repo.String(implementation.ImplementationPrefix))
	}
}

func main() {
	header := implementation.GraphqlSpecification.Repo.String(implementation.SpecificationPrefix)

	cli := cmd.CLI{}
	runResult, err := cli.Run(&cmd.RunParams{
		Choices: choices,
		Header:  header,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(runResult)
}
