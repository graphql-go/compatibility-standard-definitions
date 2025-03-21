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

2- `puller`: Git clones the **GraphQL JavaScript reference implementation repository**.

3- `puller`: Git clones a **GraphQL implementation repository**.

4- `extractor`: Pulls the type system definitions by parsing from the **GraphQL specification repository**.

5- `extractor`: Pulls the type system definitions by introspection from the **GraphQL JavaScript reference implementation**.

6- `extractor`: Pulls the type system definitions by introspection from a **GraphQL implementation**.

7- `executor`: Executes the type system definitions introspection result on a **GraphQL implementation**.

8- `validator`: Validates by comparing the schema of the **GraphQL specification** against the **GraphQL JavaScript reference implementation** schema.

9- `validator`: Validates by comparing the schema of the **GraphQL implementation** against the **GraphQL specification** schema.

10- `validator`: Compares the type system definitions of the specification against an implementation.


Notes: When validating the schema we use the JavaScript reference implementation as middle layer for confirming that the specification is up-to-date. 

### Further Work

- The type system definition extraction from the **GraphQL specfication repository** can be used as a source for auto-generating any implementation in any programming language.

- Add functionality to be used through GitHub app so it can analyze any implementation in any programming language.


#### Dependencies Analysis

##### Markdown

_Analysis date: March 12th 2025._

|               | blackfriday | goldmark | go/doc/comment |
| :---------------- | :------: | ----: |----: |
| GitHub URL        |   https://github.com/russross/blackfriday   | https://github.com/yuin/goldmark | https://pkg.go.dev/go/doc/comment |
| GitHub Stars           |   5.5k   | 3.9k | 126k |
| Last Commit Date           |   Oct 26, 2020   | Feb 18, 2025 | Apr 11, 2022 |
| **Markdown Parsing Features** |      |  | |
| Headings |   ✔️   | ✔️ | ✔️ |
| **Unit Tests Support** |      |  | |
| Parser |   ✔️   | ✔️ | ✔️ |
| **AST Walk** |   ✔️   | ✔️ | ✔️ | 
| Node Operations | Do have operations at node level, but obtaining the node value as a string were not possible. | Do have operations at node level, but obtaining the node value as a string were not possible. | Do have operations at node level, and it is possible to obtain the node value. |

##### JSON Assert

_Analysis date: March 12th 2025._

|               | github.com/stretchr/testify | github.com/google/go-cmp |
| :---------------- | ----: | ----: |
| GitHub URL        |  https://github.com/stretchr/testify | https://github.com/google/go-cmp |
| GitHub Stars           |   24.2k | 4.3k |
| Last Commit Date           |   Dec 21, 2024 | Jan 14, 2025 |
| JSON Assertion |  ✔️ | ✔️ |
| Dependant of testing.T | ✔️ | ❌ |



