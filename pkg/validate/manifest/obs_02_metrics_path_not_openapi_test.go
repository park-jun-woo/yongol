//ff:func feature=validate type=test control=iteration dimension=1 topic=manifest-observability
//ff:what obs02MetricsPathNotOpenAPI — metrics.path가 OpenAPI path와 충돌하지 않는지 검증

package manifest

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	pm "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestObs02MetricsPathNotOpenAPI(t *testing.T) {
	docWithMetrics := &openapi3.T{Paths: &openapi3.Paths{}}
	docWithMetrics.Paths.Set("/metrics", &openapi3.PathItem{
		Get: &openapi3.Operation{Responses: &openapi3.Responses{}},
	})

	docWithoutMetrics := &openapi3.T{Paths: &openapi3.Paths{}}
	docWithoutMetrics.Paths.Set("/users", &openapi3.PathItem{
		Get: &openapi3.Operation{Responses: &openapi3.Responses{}},
	})

	cases := []TestObs02MetricsPathNotOpenAPICase{
		{name: "nil_fs", fs: nil, wantCount: 0},
		{name: "nil_manifest", fs: &yongol.Fullstack{}, wantCount: 0},
		{
			name: "no_collision",
			fs: &yongol.Fullstack{
				Manifest:   &pm.ProjectConfig{Backend: pm.Backend{Observability: &pm.Observability{Metrics: &pm.ObservabilityMetrics{Path: "/metrics"}}}},
				OpenAPIDoc: docWithoutMetrics,
			},
			wantCount: 0,
		},
		{
			name: "collision_produces_error",
			fs: &yongol.Fullstack{
				Manifest:   &pm.ProjectConfig{Backend: pm.Backend{Observability: &pm.Observability{Metrics: &pm.ObservabilityMetrics{Path: "/metrics"}}}},
				OpenAPIDoc: docWithMetrics,
			},
			wantCount: 1,
		},
		{
			name: "invalid_path_returns_empty_no_diag",
			fs: &yongol.Fullstack{
				Manifest:   &pm.ProjectConfig{Backend: pm.Backend{Observability: &pm.Observability{Metrics: &pm.ObservabilityMetrics{Path: "metrics"}}}},
				OpenAPIDoc: docWithMetrics,
			},
			wantCount: 0,
		},
		{
			name: "nil_openapi_doc",
			fs: &yongol.Fullstack{
				Manifest: &pm.ProjectConfig{Backend: pm.Backend{Observability: &pm.Observability{Metrics: &pm.ObservabilityMetrics{Path: "/metrics"}}}},
			},
			wantCount: 0,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runObs02MetricsPathNotOpenAPI(t, c)
		})
	}
}
