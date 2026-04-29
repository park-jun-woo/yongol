//ff:func feature=validate type=test control=sequence topic=response-body-required
//ff:what O-5 — 204 No Content 는 의도된 빈 body 로 진단 0

package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestO05_Case3_204Exempt(t *testing.T) {
	op := opWithResponses("Op3", map[string]*openapi3.ResponseRef{
		"204": emptyResponse("No Content"),
	})
	paths := openapi3.NewPaths(openapi3.WithPath("/x", &openapi3.PathItem{Delete: op}))
	fs := newFullstackWithPaths(paths)
	if diags := o05ResponseBodyRequired(fs); len(diags) != 0 {
		t.Fatalf("expected 0 diagnostics (204 exempt), got %d: %+v", len(diags), diags)
	}
}
