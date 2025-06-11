package types

type IntrospectionSchema struct {
	Description      *string                      `json:"description"`
	QueryType        *IntrospectionNamedTypeRef   `json:"queryType"`
	MutationType     *IntrospectionNamedTypeRef   `json:"mutationType"`
	SubscriptionType *IntrospectionNamedTypeRef   `json:"subscriptionType"`
	Types            []IntrospectionFullType      `json:"types"`
	Directives       []IntrospectionDirective     `json:"directives"`
}
