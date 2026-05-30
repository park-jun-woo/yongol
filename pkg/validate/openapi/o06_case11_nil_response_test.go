//ff:func feature=validate type=test control=sequence topic=openapi-structural
//ff:what O-6 — nil response ref + nil response value skip 시 진단 0

package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestO06_Case11_NilResponse(t *testing.T) {
	resps := openapi3.NewResponses()
	resps.Set("200", nil)
	resps.Set("204", &openapi3.ResponseRef{})
	doc := &openapi3.T{Paths: openapi3.NewPaths(openapi3.WithPath("/r", &openapi3.PathItem{
		Get: &openapi3.Operation{OperationID: "getR", Responses: resps},
	}))}
	fs := &yongol.Fullstack{OpenAPIDoc: doc, OpenAPILines: o06EmptyLines()}
	if diags := o06RequiredPropertyConsistency(fs); len(diags) != 0 {
		t.Fatalf("expected 0, got %+v", diags)
	}
}
