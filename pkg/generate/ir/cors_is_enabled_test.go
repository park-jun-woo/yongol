//ff:func feature=gen-ir type=test control=sequence
//ff:what modulePath/projectID/rateLimitHasEntries/corsIsEnabled/otelEnabled/prometheusEnabled/securityHeadersEnabled/hasAuthSequence manifest 추출·판정 헬퍼
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestCorsIsEnabled(t *testing.T) {
	if corsIsEnabled(nil) {
		t.Errorf("nil should be false")
	}
	noCORS := &yongol.Fullstack{Manifest: &manifest.ProjectConfig{}}
	if corsIsEnabled(noCORS) {
		t.Errorf("nil CORS should be false")
	}
	on := &yongol.Fullstack{Manifest: &manifest.ProjectConfig{Backend: manifest.Backend{
		CORS: &manifest.CORSConfig{Enabled: true},
	}}}
	if !corsIsEnabled(on) {
		t.Errorf("enabled CORS should be true")
	}
}
