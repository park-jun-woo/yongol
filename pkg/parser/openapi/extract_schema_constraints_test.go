//ff:func feature=openapi-parse type=test control=sequence
//ff:what TestOpenAPIHelpers — unit tests for the pure openapi parser helper functions
package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestExtractSchemaConstraints(t *testing.T) {
	schema := &openapi3.Schema{
		Required: []string{"id"},
		Properties: openapi3.Schemas{
			"id":   {Value: &openapi3.Schema{Type: &openapi3.Types{"integer"}}},
			"name": {Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
			"bad":  {Value: nil}, // skipped
		},
	}
	fields := extractSchemaConstraints(schema)
	if !fields["id"].Required {
		t.Error("id should be required")
	}
	if fields["name"].Required {
		t.Error("name should not be required")
	}
	if fields["id"].Type != "integer" {
		t.Errorf("id type = %q", fields["id"].Type)
	}
	if _, ok := fields["bad"]; ok {
		t.Error("nil-value property should be skipped")
	}
}
