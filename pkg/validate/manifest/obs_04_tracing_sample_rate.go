//ff:func feature=validate type=rule control=sequence topic=manifest-observability
//ff:what OBS-004 — backend.observability.tracing.sample_rate 는 0.0~1.0 범위여야 함

package manifest

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// obs04TracingSampleRate rejects head-sampler ratios outside [0.0, 1.0].
// Zero is accepted (explicit "never sample" — useful to keep the SDK wired
// but silent in a canary deploy). Values below zero or above one are
// nonsensical to OTel's TraceIDRatioBased sampler and will be silently
// clamped by the SDK, which is a debugging hazard — better to fail fast.
//
// Rule only fires when tracing.enabled is true (same gating as OBS-003)
// and when the user explicitly set a non-default value. Omitting the
// field → codegen uses 1.0, which is trivially valid.
func obs04TracingSampleRate(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil || fs.Manifest == nil {
		return nil
	}
	obs := fs.Manifest.Backend.Observability
	if obs == nil || obs.Tracing == nil {
		return nil
	}
	if !obs.Tracing.Enabled {
		return nil
	}
	r := obs.Tracing.SampleRate
	// Default-zero (field omitted) is indistinguishable from explicit 0.0 at
	// the struct level; both are valid for the sampler, so no diagnostic.
	if r >= 0.0 && r <= 1.0 {
		return nil
	}
	return []diagnostic.Diagnostic{{
		File:    "manifest.yaml",
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelError,
		Message: fmt.Sprintf("[OBS-004] backend.observability.tracing.sample_rate must be in [0.0, 1.0] (got %v)", r),
		Advice:  "dev: 1.0 (모든 trace 수집), prod: 0.05~0.1 (head-based 5~10%) 권장",
	}}
}
