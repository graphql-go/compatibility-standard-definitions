package introspection

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"graphql-go/compatibility-standard-definitions/types"
)

func TestSpecificationQuery(t *testing.T) {
	introspection := NewIntrospection()

	result, err := introspection.SpecificationQuery()
	if err != nil {
		log.Fatalf("expected no error, got: %v", err)
	}

	expected := &types.IntrospectionQueryResult{}

	assert.Equal(t, expected, result)
}
