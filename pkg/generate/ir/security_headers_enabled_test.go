//ff:func feature=gen-ir type=test control=sequence
//ff:what modulePath/projectID/rateLimitHasEntries/corsIsEnabled/otelEnabled/prometheusEnabled/securityHeadersEnabled/hasAuthSequence manifest 추출·판정 헬퍼
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestSecurityHeadersEnabled(t *testing.T) {
	if !securityHeadersEnabled(nil) {
		t.Errorf("nil should default true")
	}
	if !securityHeadersEnabled(&yongol.Fullstack{Manifest: &manifest.ProjectConfig{}}) {
		t.Errorf("no config should default true")
	}
	disabled := &yongol.Fullstack{Manifest: &manifest.ProjectConfig{Backend: manifest.Backend{
		SecurityHeaders: &manifest.SecurityHeadersConfig{Enabled: boolPtr(false)},
	}}}
	if securityHeadersEnabled(disabled) {
		t.Errorf("explicit false should be false")
	}
}
