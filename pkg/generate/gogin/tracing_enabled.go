//ff:func feature=gen-gogin type=util control=sequence
//ff:what tracingEnabled — tracing 활성 상태의 ObservabilityTracing 반환 (아니면 nil)

package gogin

import (
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// tracingEnabled returns the tracing block when enabled, else nil. Used by
// generateGoMod to gate OTel dependency injection and by downstream
// generators that need the exporter kind for require-list shaping.
func tracingEnabled(fs *yongol.Fullstack) *pmanifest.ObservabilityTracing {
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
	return obs.Tracing
}
