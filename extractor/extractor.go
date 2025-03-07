package extractor

import (
	"bytes"
	"log"
	"os"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
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

	buf := bytes.Buffer{}
	if err := goldmark.Convert(rawMarkdown, &buf, parser.WithContext(parser.NewContext())); err != nil {
		return nil, err
	}

	log.Println(string(buf.String()))

	return &ExtractorResult{}, nil
}

func (e *Extractor) readTypeSystem() ([]byte, error) {
	f, err := os.ReadFile("./repos/graphql-specification/spec/Section 3 -- Type System.md")
	if err != nil {
		return []byte{}, err
	}

	return f, nil
}
