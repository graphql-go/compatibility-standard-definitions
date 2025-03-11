package types

type IntrospectionDirective struct {
	name         string                  `json:"name"`
	description  string                  `json:"description"`
	isRepeatable bool                    `json:"isRepeatable"`
	locations    DirectiveLocation       `json:"locations"`
	args         IntrospectionInputValue `json:"args"`
}

type DirectiveLocation string

const (
	QUERY                  DirectiveLocation = "QUERY"
	MUTATION                                 = "MUTATION"
	SUBSCRIPTION                             = "SUBSCRIPTION"
	FIELD                                    = "FIELD"
	FRAGMENT_DEFINITION                      = "FRAGMENT_DEFINITION"
	FRAGMENT_SPREAD                          = "FRAGMENT_SPREAD"
	INLINE_FRAGMENT                          = "INLINE_FRAGMENT"
	VARIABLE_DEFINITION                      = "VARIABLE_DEFINITION"
	SCHEMA                                   = "SCHEMA"
	SCALAR                                   = "SCALAR"
	OBJECT                                   = "OBJECT"
	FIELD_DEFINITION                         = "FIELD_DEFINITION"
	ARGUMENT_DEFINITION                      = "ARGUMENT_DEFINITION"
	INTERFACE                                = "INTERFACE"
	UNION                                    = "UNION"
	ENUM                                     = "ENUM"
	ENUM_VALUE                               = "ENUM_VALUE"
	INPUT_OBJECT                             = "INPUT_OBJECT"
	INPUT_FIELD_DEFINITION                   = "INPUT_FIELD_DEFINITION"
)

type IntrospectionInputValue struct {
	name              string                    `json:"name"`
	description       string                    `json:"description"`
	typeRef           IntrospectionInputTypeRef `json:"typeRef"`
	defaultValue      string                    `json:"defaultValue"`
	isDeprecated      bool                      `json:"isDeprecated"`
	deprecationReason string                    `json:"deprecationReason"`
}

type IntrospectionInputType interface {
}

type IntrospectionInputTypeRef interface {
}

type IntrospectionNamedTypeRef interface {
	IntrospectionInputType
}

type IntrospectionListTypeRef interface {
	IntrospectionInputTypeRef
}

type IntrospectionNonNullTypeRef interface {
	IntrospectionNamedTypeRef
	IntrospectionListTypeRef
}
