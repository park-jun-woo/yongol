//ff:func feature=validate type=test control=sequence topic=ssac-openapi
//ff:what xos21ErrStatusNotInOpenAPI — nil doc/no matching op/status 검증
package openapi_ssac

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXos21ErrStatusNotInOpenAPI_Unit(t *testing.T) {
	t.Run("nil doc returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := xos21ErrStatusNotInOpenAPI(fs)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("no matching op skipped", func(t *testing.T) {
		doc := &openapi3.T{Paths: openapi3.NewPaths()}
		fs := &yongol.Fullstack{
			OpenAPIDoc: doc,
			ServiceFuncs: []ssac.ServiceFunc{
				{Name: "getUser"},
			},
		}
		diags := xos21ErrStatusNotInOpenAPI(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d", len(diags))
		}
	})

	t.Run("status defined passes", func(t *testing.T) {
		resps := openapi3.NewResponses()
		resps.Set("404", &openapi3.ResponseRef{Value: &openapi3.Response{}})
		doc := &openapi3.T{
			Paths: openapi3.NewPaths(
				openapi3.WithPath("/users/{id}", &openapi3.PathItem{
					Get: &openapi3.Operation{OperationID: "getUser", Responses: resps},
				}),
			),
		}
		fs := &yongol.Fullstack{
			OpenAPIDoc: doc,
			ServiceFuncs: []ssac.ServiceFunc{
				{
					Name: "getUser",
					Sequences: []ssac.Sequence{
						{Type: "empty", ErrStatus: 404},
					},
				},
			},
		}
		diags := xos21ErrStatusNotInOpenAPI(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})
}
