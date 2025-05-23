package cmd

import (
	"fmt"

	"graphql-go/compatibility-standard-definitions/bubbletea"
)

type CLI struct {
}

type RunResult struct {
	Choice string
}

type RunParams struct {
	Choices []string
	Header  string
}

func (c *CLI) Run(p *RunParams) (*RunResult, error) {
	bt := bubbletea.New(&bubbletea.Params{
		Choices: p.Choices,
		UI: bubbletea.UIParams{
			Header: p.Header,
		},
	})

	btRunResult, err := bt.Run()
	if err != nil {
		return nil, fmt.Errorf("failed to run: %w", err)
	}

	return &RunResult{
		Choice: btRunResult.Choice,
	}, nil
}
