package types

// IntrospectionFullType represents a complete type from introspection
type IntrospectionFullType struct {
	Kind          string                     `json:"kind"`
	Name          *string                    `json:"name"`
	Description   *string                    `json:"description"`
	Fields        []IntrospectionField       `json:"fields"`
	InputFields   []IntrospectionInputValue  `json:"inputFields"`
	Interfaces    []IntrospectionTypeRef     `json:"interfaces"`
	EnumValues    []IntrospectionEnumValue   `json:"enumValues"`
	PossibleTypes []IntrospectionTypeRef     `json:"possibleTypes"`
	OfType        *IntrospectionTypeRef      `json:"ofType"`
}

// IntrospectionTypeRef represents a type reference
type IntrospectionTypeRef struct {
	Kind   string                `json:"kind"`
	Name   *string               `json:"name"`
	OfType *IntrospectionTypeRef `json:"ofType"`
}

// IntrospectionNamedTypeRef represents a simple named type reference
type IntrospectionNamedTypeRef struct {
	Name string `json:"name"`
}

// IntrospectionEnumValue represents an enum value
type IntrospectionEnumValue struct {
	Name              string  `json:"name"`
	Description       *string `json:"description"`
	IsDeprecated      bool    `json:"isDeprecated"`
	DeprecationReason *string `json:"deprecationReason"`
}
