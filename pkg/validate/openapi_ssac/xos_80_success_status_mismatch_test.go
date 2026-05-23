//ff:func feature=validate type=test control=sequence topic=ssac-openapi
//ff:what xos80SuccessStatusMismatch — nil doc/empty funcs 검증

package openapi_ssac

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXos80SuccessStatusMismatch(t *testing.T) {
	t.Run("nil doc returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := xos80SuccessStatusMismatch(fs)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("empty funcs returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{
			OpenAPIDoc: &openapi3.T{Paths: openapi3.NewPaths()},
		}
		diags := xos80SuccessStatusMismatch(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d", len(diags))
		}
	})

	t.Run("func with response and matching 200 passes", func(t *testing.T) {
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
		diags := xos80SuccessStatusMismatch(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})
}
