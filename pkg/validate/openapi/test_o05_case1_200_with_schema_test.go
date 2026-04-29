//ff:func feature=validate type=test control=sequence topic=response-body-required
//ff:what O-5 — 200 + content+schema 시 진단 0

package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestO05_Case1_200WithSchema(t *testing.T) {
	op := opWithResponses("Op1", map[string]*openapi3.ResponseRef{
		"200": jsonResponseWithSchema("OK"),
	})
	paths := openapi3.NewPaths(openapi3.WithPath("/x", &openapi3.PathItem{Get: op}))
	fs := newFullstackWithPaths(paths)
	if diags := o05ResponseBodyRequired(fs); len(diags) != 0 {
		t.Fatalf("expected 0 diagnostics, got %d: %+v", len(diags), diags)
	}
}
