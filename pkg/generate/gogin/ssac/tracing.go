//ff:func feature=gen-gogin type=util control=sequence topic=observability
//ff:what tracingWrapCalls — manifest 에서 @call span wrap 옵션 여부 조회

package ssac

import "github.com/park-jun-woo/yongol/pkg/yongol"

// tracingWrapCalls returns true only when BOTH master tracing and the
// wrap_calls opt-in are set in manifest.backend.observability.tracing.
// Keeping the two flags distinct lets services enable OTel broadly for
// HTTP + DB spans without paying the cost of a span per @call. Flip
// wrap_calls on only for deep-debug builds or low-traffic services.
func tracingWrapCalls(fs *yongol.Fullstack) bool {
	if fs == nil || fs.Manifest == nil {
		return false
	}
	obs := fs.Manifest.Backend.Observability
	if obs == nil || obs.Tracing == nil {
		return false
	}
	if !obs.Tracing.Enabled {
		return false
	}
	return obs.Tracing.WrapCalls
}
