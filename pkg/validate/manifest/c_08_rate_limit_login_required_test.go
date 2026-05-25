//ff:func feature=validate type=test control=iteration dimension=1 topic=manifest-structural
//ff:what c08RateLimitLoginRequired — rate_limit 에 Login 항목 필수 검증

package manifest

import (
	"testing"

	pm "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestC08RateLimitLoginRequired(t *testing.T) {
	cases := []TestC08RateLimitLoginRequiredCase{
		{name: "nil_fs", fs: nil, wantCount: 0},
		{name: "nil_manifest", fs: &yongol.Fullstack{}, wantCount: 0},
		{name: "no_auth", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{}}, wantCount: 0},
		{name: "no_rate_limit", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{
			Backend: pm.Backend{Auth: &pm.Auth{}},
		}}, wantCount: 0},
		{name: "rate_limit_with_login", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{
			Backend: pm.Backend{
				Auth: &pm.Auth{},
				RateLimit: pm.RateLimitConfig{
					"Login": pm.RateLimitEntry{Rate: 5, Period: "1m", Key: "ip"},
				},
			},
		}}, wantCount: 0},
		{name: "rate_limit_without_login", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{
			Backend: pm.Backend{
				Auth: &pm.Auth{},
				RateLimit: pm.RateLimitConfig{
					"Signup": pm.RateLimitEntry{Rate: 3, Period: "1m", Key: "ip"},
				},
			},
		}}, wantCount: 1},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runC08RateLimitLoginRequired(t, c)
		})
	}
}
