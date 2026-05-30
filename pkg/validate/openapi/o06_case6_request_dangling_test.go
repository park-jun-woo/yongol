//ff:func feature=validate type=test control=sequence topic=openapi-structural
//ff:what O-6 — requestBody inline 스키마 dangling required → ERROR

package openapi

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestO06_Case6_RequestDangling(t *testing.T) {
	doc := &openapi3.T{Paths: openapi3.NewPaths(openapi3.WithPath("/workflows", &openapi3.PathItem{
		Post: &openapi3.Operation{
			OperationID: "createWorkflow",
			RequestBody: &openapi3.RequestBodyRef{Value: &openapi3.RequestBody{
				Content: o06JSONBody(o06SchemaWithRequired([]string{"actions_json"}, []string{"actions_json", "ghost_field"})),
			}},
		},
	}))}
	fs := &yongol.Fullstack{OpenAPIDoc: doc, OpenAPILines: o06EmptyLines()}
	diags := o06RequiredPropertyConsistency(fs)
	if len(diags) != 1 || !strings.Contains(diags[0].Message, "ghost_field") {
		t.Fatalf("expected 1 ghost_field diag, got %+v", diags)
	}
}
