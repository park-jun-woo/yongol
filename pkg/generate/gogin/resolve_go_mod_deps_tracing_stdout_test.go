//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestResolveGoModDeps — tracing off/on(otlp/stdout/default) 의존성 병합 분기 검증
package gogin

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestResolveGoModDeps_TracingStdout(t *testing.T) {
	deps := resolveGoModDeps(fsWithTracing(&pmanifest.ObservabilityTracing{Enabled: true, Exporter: "stdout"}))
	if !hasDepWith(deps, "stdouttrace") {
		t.Errorf("expected stdouttrace exporter dep: %v", deps)
	}
	if hasDepWith(deps, "otlptracegrpc") {
		t.Errorf("did not expect otlp dep for stdout exporter: %v", deps)
	}
}
