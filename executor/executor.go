package executor

import (
	"graphql-go/compatibility-standard-definitions/types"
	"log"
)

// Executor handles the resolution of a graphql introspection query.
type Executor struct {
}

// New returns a pointer to the Executor struct.
func New() *Executor {
	return &Executor{}
}

// ExecuteResult is the result of the execute method.
type ExecuteResult struct {
	Result string
}

// ExecuteParams is the params of the execute method.
type ExecuteParams struct {
	Implementation types.Implementation
}

// Execute executes and returns the resolution of a graphql introspection query.
func (e *Executor) Execute(params ExecuteParams) (*ExecuteResult, error) {
	log.Println(params.Implementation.Repo.Dir)
	log.Println(params.Implementation.Introspection.Query)
	return &ExecuteResult{
		Result: "",
	}, nil
}
