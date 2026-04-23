//ff:func feature=gen-gogin type=util control=sequence topic=observability
//ff:what otelSampleRate — tracing.sample_rate 값 결정 (0 이하는 1.0 기본값)

package boot

import "github.com/park-jun-woo/yongol/pkg/yongol"

// otelSampleRate returns the head-sampler ratio clamped to a default of 1.0
// when the field is zero AND tracing was only implicitly enabled (since
// Go's zero value for float64 is 0.0, indistinguishable from explicit "0").
// To preserve the "explicit 0 = never sample" semantic, the helper treats
// **any** non-positive value as the 1.0 default so users who truly want 0
// must be explicit via env override at runtime.
func otelSampleRate(fs *yongol.Fullstack) float64 {
	if !hasOtel(fs) {
		return 1.0
	}
	r := fs.Manifest.Backend.Observability.Tracing.SampleRate
	if r <= 0 {
		return 1.0
	}
	return r
}
