//ff:func feature=validate type=test control=sequence topic=response-body-required
//ff:what O-5 — 500 (description 만) 은 ERROR 1개

package openapi

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestO05_Case8_500NoContent(t *testing.T) {
	op := opWithResponses("Op8", map[string]*openapi3.ResponseRef{
		"500": emptyResponse("Internal Server Error"),
	})
	paths := openapi3.NewPaths(openapi3.WithPath("/x", &openapi3.PathItem{Post: op}))
	fs := newFullstackWithPaths(paths)

	diags := o05ResponseBodyRequired(fs)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d: %+v", len(diags), diags)
	}
	if !strings.Contains(diags[0].Message, "500") {
		t.Errorf("message missing 500 status: %q", diags[0].Message)
	}
}
