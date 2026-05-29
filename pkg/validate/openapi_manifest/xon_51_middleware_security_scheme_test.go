//ff:func feature=validate type=test control=sequence topic=config-check
//ff:what xon51MiddlewareSecurityScheme — nil manifest/매칭/누락 검증

package openapi_manifest

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXon51MiddlewareSecurityScheme_Unit(t *testing.T) {
	t.Run("nil manifest returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := xon51MiddlewareSecurityScheme(fs)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("matching scheme passes", func(t *testing.T) {
		fs := &yongol.Fullstack{
			Manifest: &pmanifest.ProjectConfig{
				Backend: pmanifest.Backend{Middleware: []string{"bearerAuth"}},
			},
			OpenAPIDoc: &openapi3.T{
				Components: &openapi3.Components{
					SecuritySchemes: openapi3.SecuritySchemes{"bearerAuth": {}},
				},
			},
		}
		diags := xon51MiddlewareSecurityScheme(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("missing scheme raises diagnostic", func(t *testing.T) {
		fs := &yongol.Fullstack{
			Manifest: &pmanifest.ProjectConfig{
				Backend: pmanifest.Backend{Middleware: []string{"bearerAuth"}},
			},
			OpenAPIDoc: &openapi3.T{},
		}
		diags := xon51MiddlewareSecurityScheme(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d: %+v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, "XON-51") {
			t.Errorf("Message missing XON-51: %s", diags[0].Message)
		}
	})
}

func TestXon51MiddlewareSecurityScheme(t *testing.T) {
	_ = t
	_ = &yongol.Fullstack{}
}
