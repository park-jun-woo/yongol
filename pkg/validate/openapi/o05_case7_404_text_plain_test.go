//ff:func feature=validate type=test control=sequence topic=response-body-required
//ff:what O-5 — 404 + content: text/plain (JSON 아님) 은 ERROR 1개

package openapi

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestO05_Case7_404TextPlain(t *testing.T) {
	op := opWithResponses("Op7", map[string]*openapi3.ResponseRef{
		"404": textPlainResponse("Not Found"),
	})
	paths := openapi3.NewPaths(openapi3.WithPath("/x", &openapi3.PathItem{Get: op}))
	fs := newFullstackWithPaths(paths)

	diags := o05ResponseBodyRequired(fs)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic (text/plain rejected), got %d: %+v", len(diags), diags)
	}
	if !strings.Contains(diags[0].Message, "404") {
		t.Errorf("message missing 404 status: %q", diags[0].Message)
	}
}
