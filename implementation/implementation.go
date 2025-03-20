package implementation

import (
	"graphql-go/compatibility-standard-definitions/types"
)

// ImplementationPrefix is the default implementation prefix.
const ImplementationPrefix = "Implementation"

// RefImplementationPrefix is the reference implementation prefix.
const RefImplementationPrefix = "Reference Implementation"

// SpecificationPrefix is a implementation prefix.
const SpecificationPrefix = "Specification"

// GraphqlGoImplementation represents the graphql-go implementation.
var GraphqlGoImplementation = types.Implementation{
	Repo: types.Repository{
		Name:          "graphql-go-graphql",
		URL:           "https://github.com/graphql-go/graphql",
		ReferenceName: "v0.8.1",
		Dir:           "./repos/graphql-go-graphql/",
	},
	Type: types.GoImplementationType,
}

// GraphqlJSImplementation represents the graphql-js implementation.
var GraphqlJSImplementation = types.Implementation{
	Repo: types.Repository{
		Name:          "graphql-graphql-js",
		URL:           "https://github.com/graphql/graphql-js",
		ReferenceName: "v0.6.0",
		Dir:           "./repos/graphql-graphql-js/",
	},
	Type:              types.RefImplementationType,
	TestNamesFilePath: "./puller-js/unit-tests.txt",
}

// GraphqlSpecification represents the graphql specification.
var GraphqlSpecification = types.Specification{
	Repo: types.Repository{
		Name:          "graphql-specification",
		URL:           "https://github.com/graphql/graphql-spec",
		ReferenceName: "October2021",
		Dir:           "./repos/graphql-specification/",
	},
}

// GraphqlSpecificationWithPrefix returns the graphql specification repository link with a prefix.
func GraphqlSpecificationWithPrefix() string {
	return GraphqlSpecification.Repo.String(SpecificationPrefix)
}

// RefImplementation is the default reference implementation.
var RefImplementation = GraphqlJSImplementation

// Implementations is the default list of graphql implementations.
var Implementations = []types.Implementation{GraphqlGoImplementation}
