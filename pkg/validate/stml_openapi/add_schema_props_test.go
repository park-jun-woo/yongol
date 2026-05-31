//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestStmlOpenAPIHelpers — unit tests for the pure stml_openapi helper functions
package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestAddSchemaProps(t *testing.T) {
	out := map[string]responseFieldInfo{}
	s := &openapi3.Schema{
		Properties: openapi3.Schemas{
			"id": {Value: &openapi3.Schema{Type: &openapi3.Types{"integer"}}},
		},
		AllOf: openapi3.SchemaRefs{
			{Value: &openapi3.Schema{Properties: openapi3.Schemas{
				"name": {Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
			}}},
			nil,          // nil allOf ref → skipped
			{Value: nil}, // nil allOf value → skipped
		},
	}
	addSchemaProps(out, s)
	if out["id"].typ != "integer" {
		t.Errorf("id typ = %q", out["id"].typ)
	}
	if out["name"].typ != "string" {
		t.Errorf("name typ = %q", out["name"].typ)
	}
	addSchemaProps(out, nil) // no-op
}
