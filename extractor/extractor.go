package extractor

import (
	"encoding/json"
	"io"
	"os"
	"strings"

	"go/doc/comment"
	"graphql-go/compatibility-standard-definitions/executor"
	"graphql-go/compatibility-standard-definitions/types"
)

// queryResultFilePath is the file path of the introspection result against the graphql javascript implementation.
const queryResultFilePath string = "./graphql-js-introspection/query-result.json"

// introspectionQueryFilePath is the file path of the introspection query of the graphql javascript implementation.
const introspectionQueryFilePath string = "./graphql-js-introspection/query.graphql"

// Extractor represents the component that handles the extraction of standard definitions.
type Extractor struct {
	// executor is the executor component that extractor delegates the execution of a graphql introspection query.
	executor *executor.Executor
}

// New returns a pointer to a Extractor struct.
func New(executor *executor.Executor) *Extractor {
	return &Extractor{
		executor: executor,
	}
}

// ExtractorParams represents the parameters of the extract method.
type ExtractorParams struct {
	Implementation types.Implementation
	Specification  types.Specification
}

// ExtractorResult represents the result of the extract method.
type ExtractorResult struct {
	// SpecificationIntrospection is the introspection types of the graphql specification.
	SpecificationIntrospection *types.SpecificationIntrospection

	// ImplementationIntrospection is the introspection types of a graphql implementation.
	ImplementationIntrospection *types.ImplementationIntrospection
}

// Extract extracts and return the introspection result from the specification and a given implementation.
func (e *Extractor) Extract(params *ExtractorParams) (*ExtractorResult, error) {
	specificationIntrospection, err := e.extractSpec()
	if err != nil {
		return nil, err
	}

	implementationIntrospection, err := e.extractImplementation(params.Implementation)
	if err != nil {
		return nil, err
	}

	return &ExtractorResult{
		SpecificationIntrospection:  specificationIntrospection,
		ImplementationIntrospection: implementationIntrospection,
	}, nil
}

// readTypeSystem reads and return the type system of the graphql specification.
func (e *Extractor) readTypeSystem() ([]byte, error) {
	f, err := os.ReadFile("./repos/graphql-specification/spec/Section 3 -- Type System.md")
	if err != nil {
		return []byte{}, err
	}

	return f, nil
}

// extractSpec extracts and returns the introspection result of the graphql specification.
func (e *Extractor) extractSpec() (*types.SpecificationIntrospection, error) {
	if _, err := e.parseSpec(); err != nil {
		return nil, err
	}

	spec, err := e.loadSpec()
	if err != nil {
		return nil, err
	}

	return spec, nil
}

// extractImplementation extracts and returns the introspection result of a graphql implementation.
func (e *Extractor) extractImplementation(implementation types.Implementation) (*types.ImplementationIntrospection, error) {
	introspectionQuery, err := e.loadIntrospectionQuery()
	if err != nil {
		return nil, err
	}

	implementation.Introspection = types.Introspection{
		Query: string(introspectionQuery),
	}

	if _, err := e.executor.Execute(executor.ExecuteParams{
		Implementation: implementation,
	}); err != nil {
		return nil, err
	}

	return &types.ImplementationIntrospection{}, nil
}

// parseSpec parses and returns the introspection result of the graphql specification from the specification github repository.
func (e *Extractor) parseSpec() (types.SpecificationIntrospection, error) {
	rawMarkdown, err := e.readTypeSystem()
	if err != nil {
		return types.SpecificationIntrospection{}, err
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

	return types.SpecificationIntrospection{}, err
}

// loadSpec loads and returns the introspection result of the graphql javascript implementation.
func (e *Extractor) loadSpec() (*types.SpecificationIntrospection, error) {
	queryResultFile, err := os.Open(queryResultFilePath)
	if err != nil {
		return nil, err
	}
	defer queryResultFile.Close()

	queryResult, err := io.ReadAll(queryResultFile)
	if err != nil {
		return nil, err
	}

	result := &types.IntrospectionQueryResult{}
	if err := json.Unmarshal(queryResult, result); err != nil {
		return nil, err
	}

	return &types.SpecificationIntrospection{
		QueryResult: *result,
	}, nil
}

// loadIntrospectionQuery loads and returns the introspection query of the graphql javascript implementation.
func (e *Extractor) loadIntrospectionQuery() ([]byte, error) {
	filePath, err := os.Open(introspectionQueryFilePath)
	if err != nil {
		return nil, err
	}
	defer filePath.Close()

	introspectionQuery, err := io.ReadAll(filePath)
	if err != nil {
		return nil, err
	}

	return introspectionQuery, nil
}
