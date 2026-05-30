//ff:func feature=validate type=test control=sequence topic=openapi-structural
//ff:what O-6 — nil path item + 빈 operation skip 시 진단 0

package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestO06_Case10_NilItem(t *testing.T) {
	paths := openapi3.NewPaths()
	paths.Set("/nilitem", nil)
	paths.Set("/empty", &openapi3.PathItem{})
	doc := &openapi3.T{Paths: paths}
	fs := &yongol.Fullstack{OpenAPIDoc: doc, OpenAPILines: o06EmptyLines()}
	if diags := o06RequiredPropertyConsistency(fs); len(diags) != 0 {
		t.Fatalf("expected 0, got %+v", diags)
	}
}
