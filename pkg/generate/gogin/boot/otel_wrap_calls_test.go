//ff:func feature=gen-gogin type=test control=sequence topic=observability
//ff:what otelWrapCalls — @call 사이트마다 span 감쌀지 여부 (hasOtel + wrap_calls 모두 필요)

package boot

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestOtelWrapCalls(t *testing.T) {
	if otelWrapCalls(nil) {
		t.Errorf("no otel should be false")
	}
	// otel disabled but wrap_calls=true → still false (master toggle gates it).
	disabled := fsTracing(&pmanifest.ObservabilityTracing{Enabled: false, WrapCalls: true})
	if otelWrapCalls(disabled) {
		t.Errorf("disabled tracing should be false")
	}
	// enabled but wrap_calls=false → false.
	noWrap := fsTracing(&pmanifest.ObservabilityTracing{Enabled: true, WrapCalls: false})
	if otelWrapCalls(noWrap) {
		t.Errorf("wrap_calls=false should be false")
	}
	both := fsTracing(&pmanifest.ObservabilityTracing{Enabled: true, WrapCalls: true})
	if !otelWrapCalls(both) {
		t.Errorf("enabled + wrap_calls should be true")
	}
}
