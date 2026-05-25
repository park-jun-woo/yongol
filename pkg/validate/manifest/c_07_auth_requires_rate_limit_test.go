//ff:func feature=validate type=test control=iteration dimension=1 topic=manifest-structural
//ff:what c07AuthRequiresRateLimit — manifest backend.auth 존재 시 backend.rate_limit 필수 검증

package manifest

import (
	"testing"

	pm "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestC07AuthRequiresRateLimit(t *testing.T) {
	cases := []TestC07AuthRequiresRateLimitCase{
		{name: "nil_fs", fs: nil, wantCount: 0},
		{name: "nil_manifest", fs: &yongol.Fullstack{}, wantCount: 0},
		{name: "no_auth", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{}}, wantCount: 0},
		{name: "auth_with_rate_limit", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{
			Backend: pm.Backend{
				Auth: &pm.Auth{},
				RateLimit: pm.RateLimitConfig{
					"Login": pm.RateLimitEntry{Rate: 5, Period: "1m", Key: "ip"},
				},
			},
		}}, wantCount: 0},
		{name: "auth_without_rate_limit", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{
			Backend: pm.Backend{Auth: &pm.Auth{}},
		}}, wantCount: 1},
		{name: "auth_with_empty_rate_limit", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{
			Backend: pm.Backend{
				Auth:      &pm.Auth{},
				RateLimit: pm.RateLimitConfig{},
			},
		}}, wantCount: 1},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runC07AuthRequiresRateLimit(t, c)
		})
	}
}
