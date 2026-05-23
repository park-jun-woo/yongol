//ff:func feature=validate type=test control=sequence topic=config-check
//ff:what sec04HTTPOverridesOperationID — nil/empty/매칭/누락 operationId 검증

package openapi_manifest

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestSec04HTTPOverridesOperationID(t *testing.T) {
	t.Run("nil manifest returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := sec04HTTPOverridesOperationID(fs)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("nil HTTP returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{
			Manifest: &pmanifest.ProjectConfig{},
		}
		diags := sec04HTTPOverridesOperationID(fs)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("empty overrides returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{
			Manifest: &pmanifest.ProjectConfig{
				Backend: pmanifest.Backend{
					HTTP: &pmanifest.HTTPConfig{},
				},
			},
		}
		diags := sec04HTTPOverridesOperationID(fs)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("matching override key passes", func(t *testing.T) {
		doc := &openapi3.T{
			Paths: openapi3.NewPaths(
				openapi3.WithPath("/upload", &openapi3.PathItem{
					Post: &openapi3.Operation{OperationID: "uploadFile"},
				}),
			),
		}
		fs := &yongol.Fullstack{
			Manifest: &pmanifest.ProjectConfig{
				Backend: pmanifest.Backend{
					HTTP: &pmanifest.HTTPConfig{
						Overrides: map[string]pmanifest.HTTPOverride{
							"uploadFile": {},
						},
					},
				},
			},
			OpenAPIDoc: doc,
		}
		diags := sec04HTTPOverridesOperationID(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("nonexistent override key raises diagnostic", func(t *testing.T) {
		doc := &openapi3.T{
			Paths: openapi3.NewPaths(),
		}
		fs := &yongol.Fullstack{
			Manifest: &pmanifest.ProjectConfig{
				Backend: pmanifest.Backend{
					HTTP: &pmanifest.HTTPConfig{
						Overrides: map[string]pmanifest.HTTPOverride{
							"nonExistent": {},
						},
					},
				},
			},
			OpenAPIDoc: doc,
		}
		diags := sec04HTTPOverridesOperationID(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d: %+v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, "SEC-04") {
			t.Errorf("Message missing SEC-04: %s", diags[0].Message)
		}
	})
}
