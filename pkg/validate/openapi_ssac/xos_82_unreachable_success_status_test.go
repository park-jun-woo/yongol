//ff:func feature=validate type=test control=sequence topic=ssac-openapi
//ff:what xos82UnreachableSuccessStatus — nil doc/funcs with response 검증

package openapi_ssac

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXos82UnreachableSuccessStatus(t *testing.T) {
	t.Run("nil doc returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := xos82UnreachableSuccessStatus(fs)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("single 2xx no warning", func(t *testing.T) {
		resps := openapi3.NewResponses()
		resps.Set("200", &openapi3.ResponseRef{Value: &openapi3.Response{}})
		doc := &openapi3.T{Paths: openapi3.NewPaths(
			openapi3.WithPath("/users/{id}", &openapi3.PathItem{
				Get: &openapi3.Operation{OperationID: "getUser", Responses: resps},
			}),
		)}
		fs := &yongol.Fullstack{
			OpenAPIDoc: doc,
			ServiceFuncs: []ssac.ServiceFunc{
				{Name: "getUser", Sequences: []ssac.Sequence{{Type: "response"}}},
			},
		}
		diags := xos82UnreachableSuccessStatus(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})
}
