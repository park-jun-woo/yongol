//ff:func feature=gen-gogin type=util control=sequence topic=observability
//ff:what otelWrapCalls — @call 사이트마다 span 감쌀지 여부 (hasOtel + wrap_calls 모두 필요)

package boot

import "github.com/park-jun-woo/yongol/pkg/yongol"

// otelWrapCalls reports whether `@call` sites should be wrapped with an
// explicit ssac-tracer span. Requires hasOtel (master toggle) AND the
// wrap_calls opt-in so ordinary services don't get a span per call.
func otelWrapCalls(fs *yongol.Fullstack) bool {
	if !hasOtel(fs) {
		return false
	}
	return fs.Manifest.Backend.Observability.Tracing.WrapCalls
}
