//ff:func feature=validate type=test control=sequence topic=openapi-structural
//ff:what O-6 — additionalProperties 스키마의 dangling required → ERROR

package openapi

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestO06_Case9_AdditionalProps(t *testing.T) {
	ap := o06SchemaWithRequired([]string{"id"}, []string{"ap_phantom"})
	root := &openapi3.SchemaRef{Value: &openapi3.Schema{
		AdditionalProperties: openapi3.AdditionalProperties{Schema: ap},
	}}
	doc := &openapi3.T{Components: &openapi3.Components{Schemas: openapi3.Schemas{"Map": root}}}
	fs := &yongol.Fullstack{OpenAPIDoc: doc, OpenAPILines: o06EmptyLines()}
	diags := o06RequiredPropertyConsistency(fs)
	if len(diags) != 1 || !strings.Contains(diags[0].Message, "ap_phantom") {
		t.Fatalf("expected 1 ap_phantom diag, got %+v", diags)
	}
}
