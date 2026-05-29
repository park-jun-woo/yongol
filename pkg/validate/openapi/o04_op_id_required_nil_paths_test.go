//ff:func feature=validate type=test control=sequence topic=openapi-structural
//ff:what O-4 — Paths 가 nil 이면 규칙 침묵

package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestO04OpIdRequired_NilPaths(t *testing.T) {
	fs := &yongol.Fullstack{OpenAPIDoc: &openapi3.T{}}
	diags := o04OpIdRequired(fs)
	if len(diags) != 0 {
		t.Fatalf("expected 0 diagnostics for nil Paths, got %d: %+v", len(diags), diags)
	}

	fs2 := &yongol.Fullstack{}
	if diags := o04OpIdRequired(fs2); len(diags) != 0 {
		t.Fatalf("expected 0 diagnostics for nil OpenAPIDoc, got %d", len(diags))
	}
}
