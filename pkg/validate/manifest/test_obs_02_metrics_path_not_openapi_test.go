//ff:func feature=validate type=test control=sequence topic=manifest-observability
//ff:what OBS-002 테스트 — metrics.path 와 OpenAPI path 충돌 양방향 검증

package manifest

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func buildDocWithPath(path string) *openapi3.T {
	doc := &openapi3.T{
		OpenAPI: "3.0.0",
		Info:    &openapi3.Info{Title: "t", Version: "1"},
		Paths:   openapi3.NewPaths(),
	}
	pi := &openapi3.PathItem{
		Get: &openapi3.Operation{OperationID: "Any", Responses: openapi3.NewResponses()},
	}
	doc.Paths.Set(path, pi)
	return doc
}

func TestObs02MetricsPathNotOpenAPI_NoCollision(t *testing.T) {
	fs := &yongol.Fullstack{
		OpenAPIDoc: buildDocWithPath("/users"),
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Observability: &pmanifest.Observability{
					Metrics: &pmanifest.ObservabilityMetrics{Path: "/metrics"},
				},
			},
		},
	}
	if got := obs02MetricsPathNotOpenAPI(fs); len(got) != 0 {
		t.Fatalf("expected 0 diagnostics, got %d: %+v", len(got), got)
	}
}

func TestObs02MetricsPathNotOpenAPI_Collision(t *testing.T) {
	fs := &yongol.Fullstack{
		OpenAPIDoc: buildDocWithPath("/metrics"),
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Observability: &pmanifest.Observability{
					Metrics: &pmanifest.ObservabilityMetrics{Path: "/metrics"},
				},
			},
		},
	}
	got := obs02MetricsPathNotOpenAPI(fs)
	if len(got) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d", len(got))
	}
	if !strings.Contains(got[0].Message, "[OBS-002]") {
		t.Fatalf("message missing [OBS-002] prefix: %q", got[0].Message)
	}
}

func TestObs02MetricsPathNotOpenAPI_DefaultCollision(t *testing.T) {
	// path unset falls back to "/metrics" — still collides.
	fs := &yongol.Fullstack{
		OpenAPIDoc: buildDocWithPath("/metrics"),
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Observability: &pmanifest.Observability{Metrics: &pmanifest.ObservabilityMetrics{}},
			},
		},
	}
	if got := obs02MetricsPathNotOpenAPI(fs); len(got) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d", len(got))
	}
}
