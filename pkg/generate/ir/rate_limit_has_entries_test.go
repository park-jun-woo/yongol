//ff:func feature=gen-ir type=test control=sequence
//ff:what modulePath/projectID/rateLimitHasEntries/corsIsEnabled/otelEnabled/prometheusEnabled/securityHeadersEnabled/hasAuthSequence manifest 추출·판정 헬퍼
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestRateLimitHasEntries(t *testing.T) {
	if rateLimitHasEntries(nil) {
		t.Errorf("nil should be false")
	}
	empty := &yongol.Fullstack{Manifest: &manifest.ProjectConfig{}}
	if rateLimitHasEntries(empty) {
		t.Errorf("empty should be false")
	}
	withRL := &yongol.Fullstack{Manifest: &manifest.ProjectConfig{Backend: manifest.Backend{
		RateLimit: manifest.RateLimitConfig{"op": {Rate: 10, Period: "1m"}},
	}}}
	if !rateLimitHasEntries(withRL) {
		t.Errorf("with entry should be true")
	}
}
