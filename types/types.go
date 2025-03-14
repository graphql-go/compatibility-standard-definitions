package types

import "fmt"

const taggedRepoURL string = "%s/releases/tag/%s"

type ImplementationType uint

const (
	GoImplementationType = iota + 1
	RefImplementationType
)

type Repository struct {
	Name          string
	URL           string
	ReferenceName string
	Dir           string
}

func (r *Repository) String(prefix string) string {
	base := fmt.Sprintf("%s: %s\n", prefix, taggedRepoURL)
	return fmt.Sprintf(base, r.URL, r.ReferenceName)
}

type Implementation struct {
	Repo              Repository
	Type              ImplementationType
	TestNames         []string
	TestNamesFilePath string
}

func (i *Implementation) MapKey(prefix string) string {
	return i.Repo.String(prefix)
}

type Specification struct {
	Repo Repository
}

// SpecificationIntrospection represents the introspection result of the graphql specification.
type SpecificationIntrospection struct {
	// QueryResult contains the result of the introspection query.
	QueryResult IntrospectionQueryResult
}

// ImplementationIntrospection represents the introspection result of a graphql implementation.
type ImplementationIntrospection struct {
	// QueryResult contains the result of the introspection query.
	QueryResult IntrospectionQueryResult
}
