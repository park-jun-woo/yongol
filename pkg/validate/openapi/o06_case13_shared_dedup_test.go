//ff:func feature=validate type=test control=sequence topic=openapi-structural
//ff:what O-6 — components 와 response 가 같은 스키마를 공유해도 dangling 1회만 보고

package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestO06_Case13_SharedDedup(t *testing.T) {
	shared := o06SchemaWithRequired([]string{"id"}, []string{"phantom"})
	resps := openapi3.NewResponses()
	resps.Set("200", &openapi3.ResponseRef{Value: &openapi3.Response{Content: o06JSONBody(shared)}})
	doc := &openapi3.T{
		Components: &openapi3.Components{Schemas: openapi3.Schemas{"Shared": shared}},
		Paths: openapi3.NewPaths(openapi3.WithPath("/s", &openapi3.PathItem{
			Get: &openapi3.Operation{OperationID: "getS", Responses: resps},
		})),
	}
	fs := &yongol.Fullstack{OpenAPIDoc: doc, OpenAPILines: o06EmptyLines()}
	if diags := o06RequiredPropertyConsistency(fs); len(diags) != 1 {
		t.Fatalf("expected exactly 1 (deduped), got %d: %+v", len(diags), diags)
	}
}
