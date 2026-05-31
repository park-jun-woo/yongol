//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestResolveGoModDeps — tracing off/on(otlp/stdout/default) 의존성 병합 분기 검증
package gogin

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestResolveGoModDeps_TracingOTLP(t *testing.T) {
	deps := resolveGoModDeps(fsWithTracing(&pmanifest.ObservabilityTracing{Enabled: true, Exporter: "otlp"}))
	if !hasDepWith(deps, "go.opentelemetry.io/otel/sdk") {
		t.Errorf("expected otel sdk dep: %v", deps)
	}
	if !hasDepWith(deps, "otlptracegrpc") {
		t.Errorf("expected otlptracegrpc exporter dep: %v", deps)
	}
}
