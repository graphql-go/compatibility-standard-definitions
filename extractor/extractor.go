package extractor

import (
	"log"
	"os"
)

type ExtractorParams struct{}
type ExtractorResult struct{}

type Extractor struct {
}

func (e *Extractor) Extract(params *ExtractorParams) (*ExtractorResult, error) {
	typeSystemStr, err := e.readTypeSystem()
	if err != nil {
		return nil, err
	}

	log.Println(typeSystemStr)

	return &ExtractorResult{}, nil
}

func (e *Extractor) readTypeSystem() (string, error) {
	f, err := os.ReadFile("./repos/graphql-specification/spec/Section 3 -- Type System.md")
	if err != nil {
		return "", err
	}

	log.Println(string(f))

	// return strings.Split(string(f), "\n"), nil

	return "", nil
}
