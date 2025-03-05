# Compatibility Standard Definitions

Golang CLI library for **validating compatibility** of any GraphQL implementation type system against the GraphQL reference implementation: [graphql-js](https://github.com/graphql/graphql-js).

Current implementation supports the following GraphQL implementations:
- [https://github.com/graphql-go/graphql](https://github.com/graphql-go/graphql)

## Use Cases

- Cross validationg of compatibility between implementation versions of the GraphQL type system.

## Quick Start

Running the library:

```
./bin/start.sh
```

### Implementation Details

1- `puller`: Git clones the **GraphQL specfication repository** and a choosen GraphQL implementation.

2- `extractor`: Pulls the type system definitions from the GraphQL specification and also from a GraphQL implementation via introspection.

