package cmd

import (
	"graphql-go/compatibility-standard-definitions/bubbletea"
)

type CLI struct {
}

type model struct {
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
		return nil, err
	}

	return &RunResult{
		Choice: btRunResult.Choice,
	}, nil
}
