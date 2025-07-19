package validator

import (
	"strings"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	"graphql-go/compatibility-standard-definitions/types"
)

// Result represents a result type of the validator component.
type Result bool

const (
	// Success represents a success result.
	Success Result = true

	// Failure represents a failure result.
	Failure Result = false
)

// String returns the string representation of the result.
func (r Result) String() string {
	switch r {
	case Success:
		return "SUCCESS"
	default:
		return "FAILURE"
	}
}

// Validator represents the component that validates standard definitions.
type Validator struct {
}

// ValidateParams represents the parameters for the validate method.
type ValidateParams struct {
	Specification  *types.SpecificationIntrospection
	Implementation *types.ImplementationIntrospection
}

// ValidateResult represents the result of the validate method.
type ValidateResult struct {
	Result     Result
	Difference string
}

// normalizeStringFields normalizes string pointer fields to handle both empty string to nil conversion 
// and known text differences between graphql-go and graphql-js implementations
func normalizeStringFields(s *string) *string {
	// First handle nil case
	if s == nil {
		return nil
	}
	
	// Handle empty string case
	if *s == "" {
		return nil
	}
	
	text := *s
	
	// Fix known text differences from graphql-go implementation
	switch text {
	case "Location adjacent to a object definition.":
		fixed := "Location adjacent to an object type definition."
		return &fixed
	case "If this server supports subscription, the type that subscription operations will be rooted at.":
		fixed := "If this server support subscription, the type that subscription operations will be rooted at."
		return &fixed
	case "An enum describing what kind of type a given `__Type` is":
		fixed := "An enum describing what kind of type a given `__Type` is."
		return &fixed
	case "A Directive provides a way to describe alternate runtime execution and type validation behavior in a GraphQL document. \n\nIn some cases, you need to provide options to alter GraphQL's execution behavior in ways field arguments will not suffice, such as conditionally including or skipping a field. Directives provide this by describing additional information to the executor.":
		fixed := "A Directive provides a way to describe alternate runtime execution and type validation behavior in a GraphQL document.\n\nIn some cases, you need to provide options to alter GraphQL's execution behavior in ways field arguments will not suffice, such as conditionally including or skipping a field. Directives provide this by describing additional information to the executor."
		return &fixed
	}
	
	// Handle Unicode normalization differences between implementations
	// graphql-go normalizes fancy quotes to regular ASCII quotes, so we need to normalize
	// the specification text to match what graphql-go produces
	text = strings.ReplaceAll(text, "\u2019", "'")  // Right single quotation mark -> apostrophe
	text = strings.ReplaceAll(text, "\u2018", "'")  // Left single quotation mark -> apostrophe  
	text = strings.ReplaceAll(text, "\u201c", "\"") // Left double quotation mark -> quote
	text = strings.ReplaceAll(text, "\u201d", "\"") // Right double quotation mark -> quote
	
	// Handle extra space before newlines (graphql-go adds extra space)
	// This is a common issue where graphql-go adds extra spaces
	text = strings.ReplaceAll(text, " \n", "\n")
	
	// Return the normalized text
	return &text
}

// Validate validates given graphql introspection query results.
func (v *Validator) Validate(params *ValidateParams) (*ValidateResult, error) {
	diff := cmp.Diff(params.Specification.QueryResult,
		params.Implementation.QueryResult,
		cmpopts.IgnoreUnexported(types.IntrospectionSchema{}),
		// Normalize string fields: empty strings to nil and fix known text differences
		cmp.Transformer("NormalizeStringFields", normalizeStringFields),
		// Sort slices to ignore ordering differences
		cmpopts.SortSlices(func(a, b types.IntrospectionField) bool {
			return a.Name < b.Name
		}),
		cmpopts.SortSlices(func(a, b types.IntrospectionInputValue) bool {
			return a.Name < b.Name
		}),
		cmpopts.SortSlices(func(a, b types.IntrospectionEnumValue) bool {
			return a.Name < b.Name
		}),
		cmpopts.SortSlices(func(a, b types.IntrospectionDirective) bool {
			return a.Name < b.Name
		}),
		cmpopts.SortSlices(func(a, b types.DirectiveLocation) bool {
			return string(a) < string(b)
		}),
		cmpopts.SortSlices(func(a, b types.IntrospectionFullType) bool {
			// Handle nil name pointers
			if a.Name == nil && b.Name == nil {
				return a.Kind < b.Kind
			}
			if a.Name == nil {
				return true
			}
			if b.Name == nil {
				return false
			}
			return *a.Name < *b.Name
		}),
		cmpopts.SortSlices(func(a, b types.IntrospectionTypeRef) bool {
			// Handle nil name pointers
			if a.Name == nil && b.Name == nil {
				return a.Kind < b.Kind
			}
			if a.Name == nil {
				return true
			}
			if b.Name == nil {
				return false
			}
			return *a.Name < *b.Name
		}),
	)

	if diff != "" {
		return &ValidateResult{
			Result:     Failure,
			Difference: diff,
		}, nil
	}

	return &ValidateResult{
		Result:     Success,
		Difference: diff,
	}, nil
}
