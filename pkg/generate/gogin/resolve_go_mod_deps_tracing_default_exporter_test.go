//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestResolveGoModDeps — tracing off/on(otlp/stdout/default) 의존성 병합 분기 검증
package gogin

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestResolveGoModDeps_TracingDefaultExporter(t *testing.T) {
	// Empty exporter falls into the "otlp"/"" case.
	deps := resolveGoModDeps(fsWithTracing(&pmanifest.ObservabilityTracing{Enabled: true}))
	if !hasDepWith(deps, "otlptracegrpc") {
		t.Errorf("expected default otlp exporter dep: %v", deps)
	}
}
