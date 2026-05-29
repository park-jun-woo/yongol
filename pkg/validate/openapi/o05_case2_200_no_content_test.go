//ff:func feature=validate type=test control=sequence topic=response-body-required
//ff:what O-5 — 200 (content 없음) 은 본 룰 영역 외이므로 진단 0

package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestO05_Case2_200NoContent(t *testing.T) {
	op := opWithResponses("Op2", map[string]*openapi3.ResponseRef{
		"200": emptyResponse("OK"),
	})
	paths := openapi3.NewPaths(openapi3.WithPath("/x", &openapi3.PathItem{Get: op}))
	fs := newFullstackWithPaths(paths)
	if diags := o05ResponseBodyRequired(fs); len(diags) != 0 {
		t.Fatalf("expected 0 diagnostics (200 outside O-5 scope), got %d: %+v", len(diags), diags)
	}
}
