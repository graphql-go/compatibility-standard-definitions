package extractor

import (
	"log"
	"os"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/text"
)

type ExtractorParams struct{}
type ExtractorResult struct{}

type Extractor struct {
}

func (e *Extractor) Extract(params *ExtractorParams) (*ExtractorResult, error) {
	rawMarkdown, err := e.readTypeSystem()
	if err != nil {
		return nil, err
	}

	parser := goldmark.DefaultParser()
	node := parser.Parse(text.NewReader(rawMarkdown))

	log.Println(node)

	return &ExtractorResult{}, nil
}

func (e *Extractor) readTypeSystem() ([]byte, error) {
	f, err := os.ReadFile("./repos/graphql-specification/spec/Section 3 -- Type System.md")
	if err != nil {
		return []byte{}, err
	}

	return f, nil
}
