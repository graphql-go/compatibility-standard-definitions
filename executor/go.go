package executor

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
	return &RunResult{}, nil
}
