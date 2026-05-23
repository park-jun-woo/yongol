//ff:func feature=validate type=test control=sequence topic=config-check
//ff:what xno50SecuritySchemeMiddleware — nil/매칭/누락 검증

package openapi_manifest

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXno50SecuritySchemeMiddleware_Unit(t *testing.T) {
	t.Run("nil doc returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := xno50SecuritySchemeMiddleware(fs)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("nil manifest returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{
			OpenAPIDoc: &openapi3.T{
				Components: &openapi3.Components{
					SecuritySchemes: openapi3.SecuritySchemes{"bearerAuth": {}},
				},
			},
		}
		diags := xno50SecuritySchemeMiddleware(fs)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("matching middleware passes", func(t *testing.T) {
		fs := &yongol.Fullstack{
			OpenAPIDoc: &openapi3.T{
				Components: &openapi3.Components{
					SecuritySchemes: openapi3.SecuritySchemes{"bearerAuth": {}},
				},
			},
			Manifest: &pmanifest.ProjectConfig{
				Backend: pmanifest.Backend{Middleware: []string{"bearerAuth"}},
			},
		}
		diags := xno50SecuritySchemeMiddleware(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("missing middleware raises diagnostic", func(t *testing.T) {
		fs := &yongol.Fullstack{
			OpenAPIDoc: &openapi3.T{
				Components: &openapi3.Components{
					SecuritySchemes: openapi3.SecuritySchemes{"bearerAuth": {}},
				},
			},
			Manifest: &pmanifest.ProjectConfig{
				Backend: pmanifest.Backend{Middleware: []string{}},
			},
		}
		diags := xno50SecuritySchemeMiddleware(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d: %+v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, "XNO-50") {
			t.Errorf("Message missing XNO-50: %s", diags[0].Message)
		}
	})
}
