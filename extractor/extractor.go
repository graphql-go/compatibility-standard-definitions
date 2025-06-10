package extractor

import (
	"encoding/json"
	"fmt"
	"go/doc/comment"
	"io"
	"log"
	"os"
	"strings"

	"graphql-go/compatibility-standard-definitions/config"
	"graphql-go/compatibility-standard-definitions/executor"
	"graphql-go/compatibility-standard-definitions/types"
)

// queryResultFilePath is the file path of the introspection result against the graphql javascript implementation.
const queryResultFilePath string = "./graphql-js-introspection/introspectionQueryResult.json"

// introspectionQueryFilePath is the file path of the introspection query of the graphql javascript implementation.
const introspectionQueryFilePath string = "./graphql-js-introspection/introspectionQuery.graphql"

// Extractor represents the component that handles the extraction of standard definitions.
type Extractor struct {
	// executor is the executor component that extractor delegates the execution of a graphql introspection query.
	executor *executor.Executor

	// cfg is the configuration of the application.
	cfg *config.Config
}

// NewParams represents the paramters for the new method.
type NewParams struct {
	// Executor is the executor parameter.
	Executor *executor.Executor

	// Config is the configuration parameter.
	Config *config.Config
}

// New returns a pointer to an Extractor struct.
func New(params *NewParams) *Extractor {
	return &Extractor{
		executor: params.Executor,
		cfg:      params.Config,
	}
}

// ExtractParams represents the parameters of the extract method.
type ExtractParams struct {
	Implementation types.Implementation
	Specification  types.Specification
}

// ExtractResult represents the result of the extract method.
type ExtractResult struct {
	// SpecificationIntrospection is the introspection types of the graphql specification.
	SpecificationIntrospection *types.SpecificationIntrospection

	// ImplementationIntrospection is the introspection types of a graphql implementation.
	ImplementationIntrospection *types.ImplementationIntrospection
}

// Extract extracts and return the introspection result from the specification and a given implementation.
func (e *Extractor) Extract(params *ExtractParams) (*ExtractResult, error) {
	specificationIntrospection, err := e.extractSpec()
	if err != nil {
		return nil, fmt.Errorf("failed extract specification: %w", err)
	}

	implementationIntrospection, err := e.extractImplementation(params.Implementation)
	if err != nil {
		return nil, fmt.Errorf("failed extract implementation: %w", err)
	}

	return &ExtractResult{
		SpecificationIntrospection:  specificationIntrospection,
		ImplementationIntrospection: implementationIntrospection,
	}, nil
}

// readTypeSystem reads and return the type system of the graphql specification.
func (e *Extractor) readTypeSystem() ([]byte, error) {
	f, err := os.ReadFile("./repos/graphql-specification/spec/Section 3 -- Type System.md")
	if err != nil {
		return []byte{}, fmt.Errorf("failed to read file: %w", err)
	}

	return f, nil
}

// extractSpec extracts and returns the introspection result of the graphql specification.
func (e *Extractor) extractSpec() (*types.SpecificationIntrospection, error) {
	if _, err := e.parseSpec(); err != nil {
		return nil, fmt.Errorf("failed parse specification: %w", err)
	}

	spec, err := e.loadSpec()
	if err != nil {
		return nil, fmt.Errorf("failed load specification: %w", err)
	}

	return spec, nil
}

// extractImplementation extracts and returns the introspection result of a graphql implementation.
func (e *Extractor) extractImplementation(implementation types.Implementation) (
	*types.ImplementationIntrospection, error) {
	introspectionQuery, err := e.loadIntrospectionQuery()
	if err != nil {
		return nil, fmt.Errorf("failed to load introspection query: %w", err)
	}

	implementation.Introspection = types.Introspection{
		Query: string(introspectionQuery),
	}

	executeResult, err := e.executor.Execute(executor.ExecuteParams{
		Implementation: implementation,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to execute: %w", err)
	}

	return &executeResult.ImplementationIntrospection, nil
}

// parseSpec parses and returns the introspection result of the graphql specification
// from the specification github repository.
func (e *Extractor) parseSpec() (types.SpecificationIntrospection, error) {
	rawMarkdown, err := e.readTypeSystem()
	if err != nil {
		return types.SpecificationIntrospection{}, err
	}

	parser := comment.Parser{}

	headingsLevel2 := []string{}

	doc := parser.Parse(string(rawMarkdown))
	for _, d := range doc.Content {
		p, ok := d.(*comment.Paragraph)
		if ok {
			for _, t := range p.Text {
				val, ok := t.(comment.Plain)
				if ok {
					if strings.HasPrefix(string(val), "##") {
						headingsLevel2 = append(headingsLevel2, string(val))
					}
				}
			}
		}
	}

	if e.cfg.IsDebug {
		log.Println(headingsLevel2)
	}

	spec := types.SpecificationIntrospection{
		QueryResult: types.IntrospectionQueryResult{},
	}

	return spec, nil
}

// loadSpec loads and returns the introspection result of the graphql javascript implementation.
func (e *Extractor) loadSpec() (*types.SpecificationIntrospection, error) {
	queryResultFile, err := os.Open(queryResultFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer queryResultFile.Close()

	queryResult, err := io.ReadAll(queryResultFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	result := &types.IntrospectionQueryResult{}
	if err := json.Unmarshal(queryResult, result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return &types.SpecificationIntrospection{
		QueryResult: *result,
	}, nil
}

// loadIntrospectionQuery loads and returns the introspection query of the graphql javascript implementation.
func (e *Extractor) loadIntrospectionQuery() ([]byte, error) {
	filePath, err := os.Open(introspectionQueryFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer filePath.Close()

	introspectionQuery, err := io.ReadAll(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	return introspectionQuery, nil
}
