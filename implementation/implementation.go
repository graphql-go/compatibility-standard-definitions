package implementation

import (
	"graphql-go/compatibility-standard-definitions/types"
)

const ImplementationPrefix = "Implementation"
const RefImplementationPrefix = "Reference Implementation"
const SpecificationPrefix = "Specification"

var GraphqlGoImplementation = types.Implementation{
	Repo: types.Repository{
		Name:          "graphql-go-graphql",
		URL:           "https://github.com/graphql-go/graphql",
		ReferenceName: "v0.8.1",
		Dir:           "./repos/graphql-go-graphql/",
	},
	Type: types.GoImplementationType,
}

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

var GraphqlSpecification = types.Specification{
	Repo: types.Repository{
		Name:          "graphql-specification",
		URL:           "https://github.com/graphql/graphql-spec",
		ReferenceName: "October2021",
		Dir:           "./repos/graphql-specification/",
	},
}

var RefImplementation = GraphqlJSImplementation

var Implementations = []types.Implementation{GraphqlGoImplementation}

var gqlGoImplURL = GraphqlGoImplementation.MapKey(ImplementationPrefix)
var jsImplURL = GraphqlJSImplementation.MapKey(ImplementationPrefix)

var ImplementationsMap = map[string]types.Implementation{
	gqlGoImplURL: GraphqlGoImplementation,
	jsImplURL:    GraphqlJSImplementation,
}
