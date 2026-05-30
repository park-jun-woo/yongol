//ff:func feature=gen-gogin type=test control=branch topic=go-mod
//ff:what TestResolveGoModDeps — tracing off/on(otlp/stdout/default) 의존성 병합 분기 검증

package gogin

import (
	"strings"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func fsWithTracing(tr *pmanifest.ObservabilityTracing) *yongol.Fullstack {
	return &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Observability: &pmanifest.Observability{Tracing: tr},
			},
		},
	}
}

func hasDepWith(deps []string, substr string) bool {
	for _, d := range deps {
		if strings.Contains(d, substr) {
			return true
		}
	}
	return false
}

func TestResolveGoModDeps_NoTracing(t *testing.T) {
	deps := resolveGoModDeps(&yongol.Fullstack{})
	if len(deps) != len(coreDeps) {
		t.Errorf("want only coreDeps (%d), got %d", len(coreDeps), len(deps))
	}
	if hasDepWith(deps, "opentelemetry") {
		t.Errorf("did not expect OTel deps when tracing disabled: %v", deps)
	}
}

func TestResolveGoModDeps_TracingOTLP(t *testing.T) {
	deps := resolveGoModDeps(fsWithTracing(&pmanifest.ObservabilityTracing{Enabled: true, Exporter: "otlp"}))
	if !hasDepWith(deps, "go.opentelemetry.io/otel/sdk") {
		t.Errorf("expected otel sdk dep: %v", deps)
	}
	if !hasDepWith(deps, "otlptracegrpc") {
		t.Errorf("expected otlptracegrpc exporter dep: %v", deps)
	}
}

func TestResolveGoModDeps_TracingDefaultExporter(t *testing.T) {
	// Empty exporter falls into the "otlp"/"" case.
	deps := resolveGoModDeps(fsWithTracing(&pmanifest.ObservabilityTracing{Enabled: true}))
	if !hasDepWith(deps, "otlptracegrpc") {
		t.Errorf("expected default otlp exporter dep: %v", deps)
	}
}

func TestResolveGoModDeps_TracingStdout(t *testing.T) {
	deps := resolveGoModDeps(fsWithTracing(&pmanifest.ObservabilityTracing{Enabled: true, Exporter: "stdout"}))
	if !hasDepWith(deps, "stdouttrace") {
		t.Errorf("expected stdouttrace exporter dep: %v", deps)
	}
	if hasDepWith(deps, "otlptracegrpc") {
		t.Errorf("did not expect otlp dep for stdout exporter: %v", deps)
	}
}
