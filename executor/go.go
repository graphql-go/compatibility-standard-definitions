package executor

import (
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
)

// Go handles the go execution of a introspection query.
type Go struct {
	// rootQuery is the top root query object configuration of the graphql schema.
	rootQuery graphql.ObjectConfig

	// schemaConfig is the graphql schema configuration.
	schemaConfig graphql.SchemaConfig
}

// NewGo returns a pointer to the Go struct.
func NewGo() *Go {
	g := &Go{}

	g.rootQuery = graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"echo": &graphql.Field{
				Type: graphql.String,
				Resolve: func(_ graphql.ResolveParams) (interface{}, error) {
					return "ok", nil
				},
			},
		},
	}

	g.schemaConfig = graphql.SchemaConfig{
		Query: graphql.NewObject(g.rootQuery),
	}

	return g
}

// RunParams represents the params of the run method.
type RunParams struct {
	Query string
}

// RunResult represents the result of the run method.
type RunResult struct {
	Result string
}

// Run runs and returns a given introspection query.
func (g *Go) Run(params *RunParams) (*RunResult, error) {
	schema, err := graphql.NewSchema(g.schemaConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to run: %w", err)
	}

	gqlParams := graphql.Params{
		Schema:        schema,
		RequestString: params.Query,
	}

	doResult := graphql.Do(gqlParams)
	if doResult.Errors != nil {
		var joinedErrs error
		for _, err := range doResult.Errors {
			if joinedErrs == nil {
				joinedErrs = fmt.Errorf("%w", err)
				continue
			}

			joinedErrs = fmt.Errorf("%w: %w", err, joinedErrs)
		}

		return nil, joinedErrs
	}

	result, err := json.Marshal(doResult)
	if err != nil {
		return nil, fmt.Errorf("failed to do marshal: %w", err)
	}

	return &RunResult{
		Result: string(result),
	}, nil
}
