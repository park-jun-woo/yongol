//ff:func feature=validate type=test control=sequence topic=response-body-required
//ff:what O-5 — OpenAPIDoc nil 시 panic 없이 진단 0

package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestO05_NilDoc(t *testing.T) {
	fs := newFullstackWithPaths(openapi3.NewPaths())
	fs.OpenAPIDoc = nil
	if diags := o05ResponseBodyRequired(fs); len(diags) != 0 {
		t.Fatalf("expected 0 diagnostics on nil doc, got %d: %+v", len(diags), diags)
	}
}
