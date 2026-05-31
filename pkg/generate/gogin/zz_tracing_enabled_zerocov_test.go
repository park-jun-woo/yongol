//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestTracingEnabled_ZeroCov — tracingEnabled 분기(nil/manifest nil/disabled/enabled) 직접 호출

package gogin

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestTracingEnabled_ZeroCov(t *testing.T) {
	// nil fs.
	if tracingEnabled(nil) != nil {
		t.Errorf("nil fs should be nil")
	}
	// nil manifest.
	if tracingEnabled(&yongol.Fullstack{}) != nil {
		t.Errorf("nil manifest should be nil")
	}
	// nil observability.
	fs := &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{}}
	if tracingEnabled(fs) != nil {
		t.Errorf("nil observability should be nil")
	}
	// observability present, tracing disabled.
	fs.Manifest.Backend.Observability = &pmanifest.Observability{
		Tracing: &pmanifest.ObservabilityTracing{Enabled: false},
	}
	if tracingEnabled(fs) != nil {
		t.Errorf("disabled tracing should be nil")
	}
	// enabled.
	fs.Manifest.Backend.Observability.Tracing.Enabled = true
	if got := tracingEnabled(fs); got == nil || !got.Enabled {
		t.Errorf("enabled tracing should be returned, got %v", got)
	}
}
