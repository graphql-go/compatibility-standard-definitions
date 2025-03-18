package executor

import (
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
)

var RootQuery = graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"echo": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "ok", nil
			},
		},
	},
}

var SchemaConfig = graphql.SchemaConfig{
	Query: graphql.NewObject(RootQuery),
}

// Go handles the go execution of a introspection query.
type Go struct {
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
	schema, err := graphql.NewSchema(SchemaConfig)
	if err != nil {
		return nil, err
	}

	gqlParams := graphql.Params{
		Schema:        schema,
		RequestString: params.Query,
	}
	doResult := graphql.Do(gqlParams)
	if doResult.Errors != nil {
		return nil, fmt.Errorf("%+v", doResult.Errors)
	}

	result, err := json.Marshal(doResult)
	if err != nil {
		return nil, err
	}

	return &RunResult{
		Result: string(result),
	}, nil
}
