//ff:func feature=validate type=test control=sequence topic=config-check
//ff:what sec05RateLimitOpRoutable — nil/empty/매칭/누락 operationId 검증

package openapi_manifest

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestSec05RateLimitOpRoutable(t *testing.T) {
	t.Run("nil manifest returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := sec05RateLimitOpRoutable(fs)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("empty rate_limit returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{
			Manifest: &pmanifest.ProjectConfig{},
		}
		diags := sec05RateLimitOpRoutable(fs)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("rate_limit operationId mapping to a route passes", func(t *testing.T) {
		doc := &openapi3.T{
			Paths: openapi3.NewPaths(
				openapi3.WithPath("/login", &openapi3.PathItem{
					Post: &openapi3.Operation{OperationID: "Login"},
				}),
			),
		}
		fs := &yongol.Fullstack{
			Manifest: &pmanifest.ProjectConfig{
				Backend: pmanifest.Backend{
					RateLimit: pmanifest.RateLimitConfig{
						"Login": {Rate: 5, Period: "1m"},
					},
				},
			},
			OpenAPIDoc: doc,
		}
		diags := sec05RateLimitOpRoutable(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("rate_limit operationId with no route raises diagnostic", func(t *testing.T) {
		doc := &openapi3.T{
			Paths: openapi3.NewPaths(),
		}
		fs := &yongol.Fullstack{
			Manifest: &pmanifest.ProjectConfig{
				Backend: pmanifest.Backend{
					RateLimit: pmanifest.RateLimitConfig{
						"Missing": {Rate: 5, Period: "1m"},
					},
				},
			},
			OpenAPIDoc: doc,
		}
		diags := sec05RateLimitOpRoutable(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d: %+v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, "SEC-05") {
			t.Errorf("Message missing SEC-05: %s", diags[0].Message)
		}
		if diags[0].Level != diagnostic.LevelError {
			t.Errorf("expected error level, got %s", diags[0].Level)
		}
	})

	t.Run("diagnostics are emitted in sorted key order", func(t *testing.T) {
		doc := &openapi3.T{Paths: openapi3.NewPaths()}
		fs := &yongol.Fullstack{
			Manifest: &pmanifest.ProjectConfig{
				Backend: pmanifest.Backend{
					RateLimit: pmanifest.RateLimitConfig{
						"Zebra":  {Rate: 1, Period: "1s"},
						"Alpha":  {Rate: 1, Period: "1s"},
						"Middle": {Rate: 1, Period: "1s"},
					},
				},
			},
			OpenAPIDoc: doc,
		}
		diags := sec05RateLimitOpRoutable(fs)
		if len(diags) != 3 {
			t.Fatalf("expected 3, got %d: %+v", len(diags), diags)
		}
		wantOrder := []string{"Alpha", "Middle", "Zebra"}
		for i, w := range wantOrder {
			if !strings.Contains(diags[i].Message, "\""+w+"\"") {
				t.Errorf("diag[%d] = %q, want key %q", i, diags[i].Message, w)
			}
		}
	})
}
