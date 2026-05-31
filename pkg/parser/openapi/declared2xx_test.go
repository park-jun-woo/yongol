//ff:func feature=openapi-parse type=test control=sequence
//ff:what TestOpenAPIHelpers — unit tests for the pure openapi parser helper functions
package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestDeclared2xx(t *testing.T) {
	op := opWith2xx(200, 201)
	op.Responses.Set("404", &openapi3.ResponseRef{Value: &openapi3.Response{}})
	op.Responses.Set("default", &openapi3.ResponseRef{Value: &openapi3.Response{}})
	set := declared2xx(op)
	if !set[200] || !set[201] {
		t.Errorf("missing 2xx codes: %v", set)
	}
	if set[404] {
		t.Error("404 should not be in 2xx set")
	}
	if len(set) != 2 {
		t.Errorf("expected 2 codes, got %v", set)
	}
	// nil op → empty map.
	if got := declared2xx(nil); len(got) != 0 {
		t.Errorf("nil op → %v", got)
	}
	// nil responses → empty map.
	if got := declared2xx(&openapi3.Operation{}); len(got) != 0 {
		t.Errorf("nil responses → %v", got)
	}
}
