package types

type IntrospectionType interface {
}

type IntrospectionScalarType struct {
	kind           string `json:"kind"`
	name           string `json:"name"`
	description    string `json:"description"`
	specifiedByURL string `json:"specifiedByURL"`
}

type IntrospectionObjectType struct {
	Kind        string             `json:"kind"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Fields      IntrospectionField `json:"fields"`
	// TODO(@chris-ramon): Replace interface{} with other strategy.
	Interfaces interface{} `json:"interfaces"`
}

type IntrospectionInterfaceType struct {
	Kind        string             `json:"kind"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Fields      IntrospectionField `json:"fields"`
	// TODO(@chris-ramon): Replace interface{} with other strategy.
	Interfaces interface{} `json:"interfaces"`
	// possibleTypes IntrospectionObjectType    `json:"possibleTypes"`
}

type IntrospectionUnionType struct {
	Kind          string                  `json:"kind"`
	Name          string                  `json:"name"`
	Description   string                  `json:"description"`
	PossibleTypes IntrospectionObjectType `json:"possibleTypes"`
}

type IntrospectionEnumType struct {
	Kind        string `json:"kind"`
	Name        string `json:"name"`
	Description string `json:"description"`
	// TODO(@chris-ramon)
	// enumValues  IntrospectionEnumValue `json:"enumValues"`
}

type IntrospectionInputObjectType struct {
	Kind        string `json:"kind"`
	Name        string `json:"name"`
	Description string `json:"description"`
	// TODO(@chris-ramon)
	// inputFields IntrospectionInputValue `json:"inputFields"`
	IsOneOf bool `json:"isOneOf"`
}
