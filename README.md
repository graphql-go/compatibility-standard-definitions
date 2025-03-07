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

1- `puller`: Git clones the **GraphQL specification repository**.

2- `puller`: Git clones a **GraphQL implementation repository**.

3- `extractor`: Pulls the type system definitions from the **GraphQL specification repository**.

4- `extractor`: Pulls the type system definitions from a **GraphQL implementation** via introspection.

5- `validator`: Compares the type system definitions of the implementation against the specification.


### Further Work

- The type system definition extraction from the **GraphQL specfication repository** can be used as a source for auto-generating any implementation in any programming language.


#### Dependencies Analysis

_Analysis date: March 7th 2025._

|               | blackfriday | goldmark |
| :---------------- | :------: | ----: |
| GitHub URL        |   https://github.com/russross/blackfriday   | https://github.com/yuin/goldmark |
| GitHub Stars           |   5.5k   | 3.9k |
| Last Commit Date           |   Oct 26, 2020   | Feb 18, 2025 |
| **Markdown Parsing Features** |      |  |
| Headings |   ✔️   | ✔️ |
| **Unit Tests Support** |      |  |
| Parser |   ✔️   | ✔️ |

