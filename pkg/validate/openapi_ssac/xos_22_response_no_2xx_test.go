//ff:func feature=validate type=test control=sequence topic=ssac-openapi
//ff:what xos22ResponseNo2xx — nil doc/no response/no 2xx/with 2xx 검증

package openapi_ssac

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXos22ResponseNo2xx_Unit(t *testing.T) {
	t.Run("nil doc returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := xos22ResponseNo2xx(fs)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("no response sequence skipped", func(t *testing.T) {
		doc := &openapi3.T{Paths: openapi3.NewPaths(
			openapi3.WithPath("/users", &openapi3.PathItem{
				Get: &openapi3.Operation{OperationID: "getUser"},
			}),
		)}
		fs := &yongol.Fullstack{
			OpenAPIDoc:   doc,
			ServiceFuncs: []ssac.ServiceFunc{{Name: "getUser", Sequences: []ssac.Sequence{{Type: "get"}}}},
		}
		diags := xos22ResponseNo2xx(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d", len(diags))
		}
	})

	t.Run("no matching op skipped", func(t *testing.T) {
		doc := &openapi3.T{Paths: openapi3.NewPaths()}
		fs := &yongol.Fullstack{
			OpenAPIDoc:   doc,
			ServiceFuncs: []ssac.ServiceFunc{{Name: "getUser", Sequences: []ssac.Sequence{{Type: "response"}}}},
		}
		diags := xos22ResponseNo2xx(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d", len(diags))
		}
	})

	t.Run("with 2xx passes", func(t *testing.T) {
		resps := openapi3.NewResponses()
		resps.Set("200", &openapi3.ResponseRef{Value: &openapi3.Response{}})
		doc := &openapi3.T{Paths: openapi3.NewPaths(
			openapi3.WithPath("/users/{id}", &openapi3.PathItem{
				Get: &openapi3.Operation{OperationID: "getUser", Responses: resps},
			}),
		)}
		fs := &yongol.Fullstack{
			OpenAPIDoc:   doc,
			ServiceFuncs: []ssac.ServiceFunc{{Name: "getUser", Sequences: []ssac.Sequence{{Type: "response"}}}},
		}
		diags := xos22ResponseNo2xx(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("no 2xx raises diagnostic", func(t *testing.T) {
		resps := openapi3.NewResponses()
		resps.Set("404", &openapi3.ResponseRef{Value: &openapi3.Response{}})
		doc := &openapi3.T{Paths: openapi3.NewPaths(
			openapi3.WithPath("/users/{id}", &openapi3.PathItem{
				Get: &openapi3.Operation{OperationID: "getUser", Responses: resps},
			}),
		)}
		fs := &yongol.Fullstack{
			OpenAPIDoc:   doc,
			ServiceFuncs: []ssac.ServiceFunc{{Name: "getUser", FileName: "user.ssac", Sequences: []ssac.Sequence{{Type: "response"}}}},
		}
		diags := xos22ResponseNo2xx(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d: %+v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, "XOS-22") {
			t.Errorf("Message missing XOS-22: %s", diags[0].Message)
		}
	})
}
