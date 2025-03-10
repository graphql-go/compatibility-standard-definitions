package types

type IntrospectionField struct {
}

type IntrospectionInterfaceType struct {
}

type IntrospectionObjectType struct {
	kind        string                     `json:"kind"`
	name        string                     `json:"name"`
	description string                     `json:"description"`
	fields      IntrospectionField         `json:"fields"`
	interfaces  IntrospectionInterfaceType `json:"interfaces"`
}

type IntrospectionSchema struct {
	description string                  `json:"description"`
	queryType   IntrospectionObjectType `json:"queryType"`
}
