//ff:func feature=validate type=test control=iteration dimension=1 topic=manifest-structural
//ff:what c10RateLimitValueValid — rate_limit 항목 Rate≥1 + Period 파싱 가능 검증

package manifest

import (
	"testing"

	pm "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestC10RateLimitValueValid(t *testing.T) {
	cases := []TestC10RateLimitValueValidCase{
		{name: "nil_fs", fs: nil, wantCount: 0},
		{name: "nil_manifest", fs: &yongol.Fullstack{}, wantCount: 0},
		{name: "empty_rate_limit", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{
			Backend: pm.Backend{Auth: &pm.Auth{}},
		}}, wantCount: 0},
		{name: "valid_entry", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{
			Backend: pm.Backend{
				Auth: &pm.Auth{},
				RateLimit: pm.RateLimitConfig{
					"Login": pm.RateLimitEntry{Rate: 5, Period: "1m", Key: "ip"},
				},
			},
		}}, wantCount: 0},
		// Empty key is fine (codegen defaults to "ip"); only Rate/Period matter.
		{name: "valid_entry_empty_key", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{
			Backend: pm.Backend{
				Auth: &pm.Auth{},
				RateLimit: pm.RateLimitConfig{
					"Login": pm.RateLimitEntry{Rate: 5, Period: "1m"},
				},
			},
		}}, wantCount: 0},
		{name: "zero_rate", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{
			Backend: pm.Backend{
				Auth: &pm.Auth{},
				RateLimit: pm.RateLimitConfig{
					"Login": pm.RateLimitEntry{Rate: 0, Period: "1m", Key: "ip"},
				},
			},
		}}, wantCount: 1},
		{name: "negative_rate", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{
			Backend: pm.Backend{
				Auth: &pm.Auth{},
				RateLimit: pm.RateLimitConfig{
					"Login": pm.RateLimitEntry{Rate: -1, Period: "1m", Key: "ip"},
				},
			},
		}}, wantCount: 1},
		{name: "empty_period", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{
			Backend: pm.Backend{
				Auth: &pm.Auth{},
				RateLimit: pm.RateLimitConfig{
					"Login": pm.RateLimitEntry{Rate: 5, Period: "", Key: "ip"},
				},
			},
		}}, wantCount: 1},
		{name: "unparseable_period", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{
			Backend: pm.Backend{
				Auth: &pm.Auth{},
				RateLimit: pm.RateLimitConfig{
					"Login": pm.RateLimitEntry{Rate: 5, Period: "60", Key: "ip"},
				},
			},
		}}, wantCount: 1},
		// Zero-value entry (loose-decode escape): both Rate and Period invalid → 2 diags.
		{name: "zero_value_entry", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{
			Backend: pm.Backend{
				Auth: &pm.Auth{},
				RateLimit: pm.RateLimitConfig{
					"Login": pm.RateLimitEntry{},
				},
			},
		}}, wantCount: 2},
		// Multiple entries: one valid, one zero-value → 2 diags from the bad one.
		{name: "mixed_entries", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{
			Backend: pm.Backend{
				Auth: &pm.Auth{},
				RateLimit: pm.RateLimitConfig{
					"Login":  pm.RateLimitEntry{Rate: 5, Period: "1m", Key: "ip"},
					"Signup": pm.RateLimitEntry{},
				},
			},
		}}, wantCount: 2},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runC10RateLimitValueValid(t, c)
		})
	}
}
