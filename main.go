package main

import (
	"log"

	"graphql-go/compatibility-standard-definitions/cmd"
)

func main() {
	cli := cmd.CLI{}
	runResult, err := cli.Run(&cmd.RunParams{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(runResult)
}
