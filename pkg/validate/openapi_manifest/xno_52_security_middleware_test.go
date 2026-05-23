//ff:func feature=validate type=test control=sequence topic=config-check
//ff:what xno52SecurityMiddleware — nil/security 매칭/누락 검증

package openapi_manifest

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXno52SecurityMiddleware_Unit(t *testing.T) {
	t.Run("nil doc returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := xno52SecurityMiddleware(fs)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("no security references no diagnostics", func(t *testing.T) {
		doc := &openapi3.T{
			Paths: openapi3.NewPaths(
				openapi3.WithPath("/users", &openapi3.PathItem{
					Get: &openapi3.Operation{OperationID: "listUsers"},
				}),
			),
		}
		fs := &yongol.Fullstack{
			OpenAPIDoc: doc,
			Manifest:   &pmanifest.ProjectConfig{},
			OpenAPILines: &oapiparser.LineIndex{
				Operations: map[string]int{},
				Paths:      map[string]int{},
			},
		}
		diags := xno52SecurityMiddleware(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("missing security middleware raises diagnostic", func(t *testing.T) {
		sec := openapi3.SecurityRequirements{{"oauth2": {}}}
		doc := &openapi3.T{
			Paths: openapi3.NewPaths(
				openapi3.WithPath("/users", &openapi3.PathItem{
					Get: &openapi3.Operation{
						OperationID: "listUsers",
						Security:    &sec,
					},
				}),
			),
		}
		fs := &yongol.Fullstack{
			OpenAPIDoc: doc,
			Manifest:   &pmanifest.ProjectConfig{},
			OpenAPILines: &oapiparser.LineIndex{
				Operations: map[string]int{},
				Paths:      map[string]int{},
			},
		}
		diags := xno52SecurityMiddleware(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d: %+v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, "XNO-52") {
			t.Errorf("Message missing XNO-52: %s", diags[0].Message)
		}
	})
}
