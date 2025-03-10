package types

type IntrospectionSchema struct {
	description      string                  `json:"description"`
	queryType        IntrospectionObjectType `json:"queryType"`
	mutationType     IntrospectionObjectType `json:"mutationType"`
	subscriptionType IntrospectionObjectType `json:"subscriptionType"`
	types            IntrospectionType       `json:"types"`
	directives       IntrospectionDirective  `json:"directives"`
}
