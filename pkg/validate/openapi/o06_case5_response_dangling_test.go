//ff:func feature=validate type=test control=sequence topic=openapi-structural
//ff:what O-6 — response inline 스키마 dangling required → ERROR

package openapi

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestO06_Case5_ResponseDangling(t *testing.T) {
	resps := openapi3.NewResponses()
	resps.Set("200", &openapi3.ResponseRef{Value: &openapi3.Response{
		Content: o06JSONBody(o06SchemaWithRequired([]string{"id"}, []string{"phantom"})),
	}})
	doc := &openapi3.T{Paths: openapi3.NewPaths(openapi3.WithPath("/workflows/{id}", &openapi3.PathItem{
		Get: &openapi3.Operation{OperationID: "getWorkflow", Responses: resps},
	}))}
	fs := &yongol.Fullstack{OpenAPIDoc: doc, OpenAPILines: o06EmptyLines()}
	diags := o06RequiredPropertyConsistency(fs)
	if len(diags) != 1 || !strings.Contains(diags[0].Message, "phantom") {
		t.Fatalf("expected 1 phantom diag, got %+v", diags)
	}
}
