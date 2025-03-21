package puller

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"

	"graphql-go/compatibility-standard-definitions/types"
)

// reposDirName is the code repository root directory name.
const reposDirName = "repos"

// Puller represents the puller component.
type Puller struct {
}

// PullerParams represents the parameters of the pull method.
type PullerParams struct {
	// Specification is the code repository of the graphql specification.
	Specification types.Repository

	// Implementation is the code repository of the graphql implementation.
	Implementation types.Repository
}

// PullerResult represents the result of the pull method.
type PullerResult struct {
}

// Pull pulls a set of code repositories and returns if it succeeded or not.
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
