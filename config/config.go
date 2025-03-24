package config

import (
	"os"

	"graphql-go/compatibility-standard-definitions/implementation"
	"graphql-go/compatibility-standard-definitions/types"
)

// Config represents a set of configuration values.
type Config struct {
	IsDebug bool

	// GraphqlGoImplementation represents the graphql-go implementation.
	GraphqlGoImplementation types.Implementation

	// GraphqlJSImplementation represents the graphql-js implementation.
	GraphqlJSImplementation types.Implementation

	// GraphqlSpecification represents the graphql specification.
	GraphqlSpecification types.Specification

	// RefImplementation is the default reference implementation.
	RefImplementation types.Implementation

	// Implementations is the default list of graphql implementations.
	Implementations []types.Implementation

	// GraphqlSpecificationWithPrefix returns the graphql specification repository link with a prefix.
	GraphqlSpecificationWithPrefix string

	// AvailableImplementations returns a list of available implementations.
	AvailableImplementations []string
}

// New returns a pointer to a Config struct.
func New() *Config {
	debug := os.Getenv("DEBUG")

	isDebug := false
	if debug == "true" {
		isDebug = true
	}

	cfg := &Config{
		IsDebug: isDebug,
	}

	cfg.GraphqlGoImplementation = types.Implementation{
		Repo: types.Repository{
			Name:          "graphql-go-graphql",
			URL:           "https://github.com/graphql-go/graphql",
			ReferenceName: "v0.8.1",
			Dir:           "./repos/graphql-go-graphql/",
		},
		Type: types.GoImplementationType,
	}

	cfg.GraphqlJSImplementation = types.Implementation{
		Repo: types.Repository{
			Name:          "graphql-graphql-js",
			URL:           "https://github.com/graphql/graphql-js",
			ReferenceName: "v0.6.0",
			Dir:           "./repos/graphql-graphql-js/",
		},
		Type:              types.RefImplementationType,
		TestNamesFilePath: "./puller-js/unit-tests.txt",
	}

	cfg.GraphqlSpecification = types.Specification{
		Repo: types.Repository{
			Name:          "graphql-specification",
			URL:           "https://github.com/graphql/graphql-spec",
			ReferenceName: "October2021",
			Dir:           "./repos/graphql-specification/",
		},
	}

	cfg.RefImplementation = cfg.GraphqlJSImplementation

	cfg.Implementations = []types.Implementation{cfg.GraphqlGoImplementation}

	cfg.GraphqlSpecificationWithPrefix = cfg.GraphqlSpecification.Repo.String(implementation.SpecificationPrefix)

	cfg.AvailableImplementations = AvailableImplementations(cfg.Implementations)

	return cfg
}

// AvailableImplementations returns a list of available implementations.
func AvailableImplementations(implementations []types.Implementation) []string {
	var result = []string{}

	for _, i := range implementations {
		result = append(result, i.Repo.String(implementation.ImplementationPrefix))
	}

	return result
}
