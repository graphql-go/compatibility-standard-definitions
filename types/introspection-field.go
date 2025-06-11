package types

type IntrospectionField struct {
	Name              string                    `json:"name"`
	Description       *string                   `json:"description"`
	Args              []IntrospectionInputValue `json:"args"`
	Type              IntrospectionTypeRef      `json:"type"`
	IsDeprecated      bool                      `json:"isDeprecated"`
	DeprecationReason *string                   `json:"deprecationReason"`
}
