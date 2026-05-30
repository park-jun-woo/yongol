//ff:func feature=validate type=test control=sequence topic=openapi-structural
//ff:what O-6 — required ⊆ properties 시 진단 0

package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestO06_Case2_SubsetPasses(t *testing.T) {
	doc := &openapi3.T{Components: &openapi3.Components{Schemas: openapi3.Schemas{
		"Workflow": o06SchemaWithRequired([]string{"id", "name"}, []string{"id"}),
	}}}
	fs := &yongol.Fullstack{OpenAPIDoc: doc, OpenAPILines: o06EmptyLines()}
	if diags := o06RequiredPropertyConsistency(fs); len(diags) != 0 {
		t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
	}
}
