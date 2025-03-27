package executor

import (
	"graphql-go/compatibility-standard-definitions/types"
)

// Executor handles the resolution of a graphql introspection query.
type Executor struct {
	goExecutor *Go
}

// New returns a pointer to the Executor struct.
func New() *Executor {
	return &Executor{
		goExecutor: NewGo(),
	}
}

// ExecuteResult is the result of the execute method.
type ExecuteResult struct {
	ImplementationIntrospection types.ImplementationIntrospection
}

// ExecuteParams is the params of the execute method.
type ExecuteParams struct {
	Implementation types.Implementation
}

// Execute executes and returns the resolution of a graphql introspection query.
func (e *Executor) Execute(params ExecuteParams) (*ExecuteResult, error) {
	runParams := &RunParams{
		Query: params.Implementation.Introspection.Query,
	}

	runResult, err := e.goExecutor.Run(runParams)
	if err != nil {
		return nil, err
	}

	return &ExecuteResult{
		ImplementationIntrospection: runResult.ImplementationIntrospection,
	}, nil
}
