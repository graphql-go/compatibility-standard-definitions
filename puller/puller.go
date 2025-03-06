package puller

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"

	"graphql-go/compatibility-standard-definitions/types"
)

const reposDirName = "repos"

type Puller struct {
}

type PullerResult struct {
}

type PullerParams struct {
	Specification  types.Repository
	Implementation types.Repository
}

func (p *Puller) Pull(params *PullerParams) (*PullerResult, error) {
	repos := []types.Repository{
		params.Specification,
		params.Implementation,
	}

	if _, err := os.Stat(reposDirName); err != nil {
		if os.IsNotExist(err) {
			if err := os.Mkdir(reposDirName, os.ModePerm); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	for _, r := range repos {
		name := filepath.Join(reposDirName, r.Name)
		if _, err := os.Stat(name); os.IsNotExist(err) {
			if err := os.Mkdir(name, os.ModePerm); err != nil {
				return nil, err
			}
		}
		if _, err := git.PlainClone(name, false, &git.CloneOptions{
			URL:      r.URL,
			Progress: os.Stdout,
		}); err != nil {
			if strings.Contains(err.Error(), "repository already exists") {
				return nil, nil
			}
			return nil, err
		}
	}

	return &PullerResult{}, nil
}
