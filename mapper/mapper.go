package mapper

import (
	"graphql-go/compatibility-standard-definitions/implementation"
)

// AvailableImplementations returns a list of available implementations.
func AvailableImplementations() []string {
	var implementations = []string{}

	for _, i := range implementation.Implementations {
		implementations = append(implementations, i.Repo.String(implementation.ImplementationPrefix))
	}

	return implementations
}
