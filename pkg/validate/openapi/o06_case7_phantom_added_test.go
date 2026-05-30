//ff:func feature=validate type=test control=sequence topic=openapi-structural
//ff:what O-6 — phantom 을 properties 에 추가하면 진단 0(재현 케이스 대조군)

package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestO06_Case7_PhantomAdded(t *testing.T) {
	doc := &openapi3.T{Components: &openapi3.Components{Schemas: openapi3.Schemas{
		"Workflow": o06SchemaWithRequired([]string{"id", "phantom"}, []string{"phantom"}),
	}}}
	fs := &yongol.Fullstack{OpenAPIDoc: doc, OpenAPILines: o06EmptyLines()}
	if diags := o06RequiredPropertyConsistency(fs); len(diags) != 0 {
		t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
	}
}
