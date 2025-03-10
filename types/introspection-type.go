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
	kind        string                     `json:"kind"`
	name        string                     `json:"name"`
	description string                     `json:"description"`
	fields      IntrospectionField         `json:"fields"`
	interfaces  IntrospectionInterfaceType `json:"interfaces"`
}

type IntrospectionInterfaceType struct {
	kind        string             `json:"kind"`
	name        string             `json:"name"`
	description string             `json:"description"`
	fields      IntrospectionField `json:"fields"`
	// TODO(@chris-ramon)
	// interfaces    IntrospectionInterfaceType `json:"interfaces"`
	// possibleTypes IntrospectionObjectType    `json:"possibleTypes"`
}

type IntrospectionUnionType struct {
	kind          string                  `json:"kind"`
	name          string                  `json:"name"`
	description   string                  `json:"description"`
	possibleTypes IntrospectionObjectType `json:"possibleTypes"`
}

type IntrospectionEnumType struct {
	kind        string `json:"kind"`
	name        string `json:"name"`
	description string `json:"description"`
	// TODO(@chris-ramon)
	// enumValues  IntrospectionEnumValue `json:"enumValues"`
}

type IntrospectionInputObjectType struct {
	kind        string `json:"kind"`
	name        string `json:"name"`
	description string `json:"description"`
	// TODO(@chris-ramon)
	// inputFields IntrospectionInputValue `json:"inputFields"`
	isOneOf bool `json:"isOneOf"`
}
