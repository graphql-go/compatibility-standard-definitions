package types

type IntrospectionDirective struct {
	Name         string                    `json:"name"`
	Description  string                    `json:"description"`
	IsRepeatable bool                      `json:"isRepeatable"`
	Locations    []DirectiveLocation       `json:"locations"`
	Args         []IntrospectionInputValue `json:"args"`
}

type DirectiveLocation string

const (
	Query                DirectiveLocation = "QUERY"
	Mutation                               = "MUTATION"
	Subscription                           = "SUBSCRIPTION"
	Field                                  = "FIELD"
	FragmentDefinition                     = "FRAGMENT_DEFINITION"
	FragmentSpread                         = "FRAGMENT_SPREAD"
	InlineFragment                         = "INLINE_FRAGMENT"
	VariableDefinition                     = "VARIABLE_DEFINITION"
	Schema                                 = "SCHEMA"
	Scalar                                 = "SCALAR"
	Object                                 = "OBJECT"
	FieldDefinition                        = "FIELD_DEFINITION"
	ArgumentDefinition                     = "ARGUMENT_DEFINITION"
	Interface                              = "INTERFACE"
	Union                                  = "UNION"
	Enum                                   = "ENUM"
	EnumValue                              = "ENUM_VALUE"
	InputObject                            = "INPUT_OBJECT"
	InputFieldDefinition                   = "INPUT_FIELD_DEFINITION"
)

type IntrospectionInputValue struct {
	Name              string                    `json:"name"`
	Description       string                    `json:"description"`
	TypeRef           IntrospectionInputTypeRef `json:"typeRef"`
	DefaultValue      string                    `json:"defaultValue"`
	IsDeprecated      bool                      `json:"isDeprecated"`
	DeprecationReason string                    `json:"deprecationReason"`
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
