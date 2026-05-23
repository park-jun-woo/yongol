//ff:func feature=validate type=test control=iteration dimension=1 topic=manifest-security-headers
//ff:what sec302HSTSShort — HSTS max_age가 180일 미만일 때 WARNING 검증

package manifest

import (
	"testing"

	pm "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestSec302HSTSShort(t *testing.T) {
	mk := func(maxAge int) *yongol.Fullstack {
		return &yongol.Fullstack{Manifest: &pm.ProjectConfig{Backend: pm.Backend{
			SecurityHeaders: &pm.SecurityHeadersConfig{HSTS: &pm.HSTSConfig{MaxAge: maxAge}},
		}}}
	}
	cases := []TestSec302HSTSShortCase{
		{name: "nil_fs", fs: nil, wantCount: 0},
		{name: "nil_manifest", fs: &yongol.Fullstack{}, wantCount: 0},
		{name: "nil_sh", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{}}, wantCount: 0},
		{name: "nil_hsts", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{Backend: pm.Backend{SecurityHeaders: &pm.SecurityHeadersConfig{}}}}, wantCount: 0},
		{name: "zero", fs: mk(0), wantCount: 0},
		{name: "negative", fs: mk(-1), wantCount: 0},
		{name: "at_minimum", fs: mk(hstsPreloadMinSeconds), wantCount: 0},
		{name: "above_minimum", fs: mk(31536000), wantCount: 0},
		{name: "below_minimum", fs: mk(86400), wantCount: 1},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runSec302HSTSShort(t, c)
		})
	}
}
