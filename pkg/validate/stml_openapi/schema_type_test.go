//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestStmlOpenAPIHelpers — unit tests for the pure stml_openapi helper functions
package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestSchemaType(t *testing.T) {
	ref := &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}}
	if got := schemaType(ref); got != "string" {
		t.Errorf("schemaType = %q, want string", got)
	}
	// nil ref / nil value / nil type → "".
	if got := schemaType(nil); got != "" {
		t.Errorf("nil ref: %q", got)
	}
	if got := schemaType(&openapi3.SchemaRef{}); got != "" {
		t.Errorf("nil value: %q", got)
	}
	if got := schemaType(&openapi3.SchemaRef{Value: &openapi3.Schema{}}); got != "" {
		t.Errorf("nil type: %q", got)
	}
	// empty type slice → "".
	if got := schemaType(&openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{}}}); got != "" {
		t.Errorf("empty type slice: %q", got)
	}
}
