//ff:func feature=validate type=test control=sequence topic=openapi-structural
//ff:what O-6 — 배열 items 의 nested dangling required → ERROR

package openapi

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestO06_Case8_NestedItems(t *testing.T) {
	item := o06SchemaWithRequired([]string{"id"}, []string{"nested_phantom"})
	arr := &openapi3.SchemaRef{Value: &openapi3.Schema{Items: item}}
	doc := &openapi3.T{Components: &openapi3.Components{Schemas: openapi3.Schemas{"WorkflowList": arr}}}
	fs := &yongol.Fullstack{OpenAPIDoc: doc, OpenAPILines: o06EmptyLines()}
	diags := o06RequiredPropertyConsistency(fs)
	if len(diags) != 1 || !strings.Contains(diags[0].Message, "nested_phantom") {
		t.Fatalf("expected 1 nested_phantom diag, got %+v", diags)
	}
}
