//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestResolveGoModDeps — tracing off/on(otlp/stdout/default) 의존성 병합 분기 검증
package gogin

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestResolveGoModDeps_NoTracing(t *testing.T) {
	deps := resolveGoModDeps(&yongol.Fullstack{})
	if len(deps) != len(coreDeps) {
		t.Errorf("want only coreDeps (%d), got %d", len(coreDeps), len(deps))
	}
	if hasDepWith(deps, "opentelemetry") {
		t.Errorf("did not expect OTel deps when tracing disabled: %v", deps)
	}
}
