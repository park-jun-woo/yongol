//ff:func feature=validate type=test control=sequence topic=response-body-required
//ff:what O-5 — 403 + content+schema 시 진단 0

package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestO05_Case6_403WithSchema(t *testing.T) {
	op := opWithResponses("Op6", map[string]*openapi3.ResponseRef{
		"403": jsonResponseWithSchema("Forbidden"),
	})
	paths := openapi3.NewPaths(openapi3.WithPath("/x", &openapi3.PathItem{Get: op}))
	fs := newFullstackWithPaths(paths)
	if diags := o05ResponseBodyRequired(fs); len(diags) != 0 {
		t.Fatalf("expected 0 diagnostics, got %d: %+v", len(diags), diags)
	}
}
