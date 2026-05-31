//ff:func feature=gen-ir type=test control=sequence
//ff:what TestConfigBuildersZeroCov — boot/middleware config 빌더 전 분기 직접 커버
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBuildRateLimitConfig_ZeroCov(t *testing.T) {
	if c := buildRateLimitConfig(nil); c != nil {
		t.Errorf("nil fullstack should give nil rate limit")
	}
	// empty rate limit map → nil.
	fsEmpty := &yongol.Fullstack{Manifest: &manifest.ProjectConfig{}}
	if c := buildRateLimitConfig(fsEmpty); c != nil {
		t.Errorf("empty rate limit should be nil")
	}
	// non-empty → config.
	fs := &yongol.Fullstack{
		Manifest: &manifest.ProjectConfig{
			Backend: manifest.Backend{
				RateLimit: manifest.RateLimitConfig{
					"Login": manifest.RateLimitEntry{Rate: 5, Period: "1m", Key: "ip"},
				},
			},
		},
	}
	if c := buildRateLimitConfig(fs); c == nil {
		t.Errorf("non-empty rate limit should give config")
	}
}
