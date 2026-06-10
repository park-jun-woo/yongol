//ff:func feature=validate type=test control=iteration dimension=1 topic=manifest-structural
//ff:what c11IpKeyRequiresProxy — ip 키 rate_limit + trusted_proxies 미설정 WARNING 검증

package manifest

import (
	"testing"

	pm "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestC11IpKeyRequiresProxy(t *testing.T) {
	cases := []TestC11IpKeyRequiresProxyCase{
		{name: "nil_fs", fs: nil, wantCount: 0},
		{name: "nil_manifest", fs: &yongol.Fullstack{}, wantCount: 0},
		{name: "no_rate_limit", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{
			Backend: pm.Backend{Auth: &pm.Auth{}},
		}}, wantCount: 0},
		// key: ip + trusted_proxies unset → one aggregated WARNING.
		{name: "ip_key_no_proxies_warns", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{
			Backend: pm.Backend{
				Auth: &pm.Auth{},
				RateLimit: pm.RateLimitConfig{
					"Login": pm.RateLimitEntry{Rate: 5, Period: "1m", Key: "ip"},
				},
			},
		}}, wantCount: 1},
		// Empty key defaults to ip in codegen → same warning.
		{name: "empty_key_no_proxies_warns", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{
			Backend: pm.Backend{
				Auth: &pm.Auth{},
				RateLimit: pm.RateLimitConfig{
					"Login": pm.RateLimitEntry{Rate: 5, Period: "1m"},
				},
			},
		}}, wantCount: 1},
		// http section present but trusted_proxies empty → still unset.
		{name: "http_without_proxies_warns", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{
			Backend: pm.Backend{
				Auth: &pm.Auth{},
				HTTP: &pm.HTTPConfig{BodyLimit: "1MiB"},
				RateLimit: pm.RateLimitConfig{
					"Login": pm.RateLimitEntry{Rate: 5, Period: "1m", Key: "ip"},
				},
			},
		}}, wantCount: 1},
		// Multiple ip-keyed entries still aggregate into a single WARNING.
		{name: "multiple_ip_keys_single_warning", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{
			Backend: pm.Backend{
				Auth: &pm.Auth{},
				RateLimit: pm.RateLimitConfig{
					"Login":  pm.RateLimitEntry{Rate: 5, Period: "1m", Key: "ip"},
					"Signup": pm.RateLimitEntry{Rate: 3, Period: "1m"},
				},
			},
		}}, wantCount: 1},
		// trusted_proxies declared → silent.
		{name: "ip_key_with_proxies_ok", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{
			Backend: pm.Backend{
				Auth: &pm.Auth{},
				HTTP: &pm.HTTPConfig{TrustedProxies: []string{"10.0.0.0/8"}},
				RateLimit: pm.RateLimitConfig{
					"Login": pm.RateLimitEntry{Rate: 5, Period: "1m", Key: "ip"},
				},
			},
		}}, wantCount: 0},
		// Non-ip key only → rule does not apply.
		{name: "non_ip_key_no_proxies_ok", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{
			Backend: pm.Backend{
				Auth: &pm.Auth{},
				RateLimit: pm.RateLimitConfig{
					"Login": pm.RateLimitEntry{Rate: 5, Period: "1m", Key: "email"},
				},
			},
		}}, wantCount: 0},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runC11IpKeyRequiresProxy(t, c)
		})
	}
}
