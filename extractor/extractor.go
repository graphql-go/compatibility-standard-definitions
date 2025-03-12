package extractor

import (
	"os"
	"strings"

	"go/doc/comment"
	"graphql-go/compatibility-standard-definitions/types"
)

type ExtractorParams struct{}

type ExtractorResult struct {
	SpecificationIntrospection  types.SpecificationIntrospection
	ImplementationIntrospection types.ImplementationIntrospection
}

type Extractor struct {
}

func (e *Extractor) Extract(params *ExtractorParams) (*ExtractorResult, error) {
	rawMarkdown, err := e.readTypeSystem()
	if err != nil {
		return nil, err
	}

	parser := comment.Parser{}
	doc := parser.Parse(string(rawMarkdown))
	for _, d := range doc.Content {
		p, ok := d.(*comment.Paragraph)
		if ok {
			for _, t := range p.Text {
				switch val := t.(type) {
				case comment.Plain:
					if strings.HasPrefix(string(val), "##") {
						// log.Println(string(val))
					}
				}
			}
		}
	}

	return &ExtractorResult{
		SpecificationIntrospection:  types.SpecificationIntrospection{},
		ImplementationIntrospection: types.ImplementationIntrospection{},
	}, nil
}

func (e *Extractor) readTypeSystem() ([]byte, error) {
	f, err := os.ReadFile("./repos/graphql-specification/spec/Section 3 -- Type System.md")
	if err != nil {
		return []byte{}, err
	}

	return f, nil
}
