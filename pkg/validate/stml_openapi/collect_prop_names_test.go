//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-openapi
//ff:what TestStmlOpenAPIHelpers — unit tests for the pure stml_openapi helper functions
package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestCollectPropNames(t *testing.T) {
	out := map[string]struct{}{}
	s := &openapi3.Schema{
		Properties: openapi3.Schemas{
			"id":   {Value: &openapi3.Schema{}},
			"name": {Value: &openapi3.Schema{}},
		},
		AllOf: openapi3.SchemaRefs{
			{Value: &openapi3.Schema{Properties: openapi3.Schemas{"email": {Value: &openapi3.Schema{}}}}},
			nil, // skipped
		},
	}
	collectPropNames(out, s)
	for _, want := range []string{"id", "name", "email"} {
		if _, ok := out[want]; !ok {
			t.Errorf("missing prop %q", want)
		}
	}
	// nil schema is a no-op.
	collectPropNames(out, nil)
}
