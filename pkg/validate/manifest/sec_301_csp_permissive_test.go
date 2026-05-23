//ff:func feature=validate type=test control=iteration dimension=1 topic=manifest-security-headers
//ff:what sec301CspPermissive — CSP default-src에 * / unsafe-eval 포함 시 WARNING 검증

package manifest

import (
	"testing"

	pm "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestSec301CspPermissive(t *testing.T) {
	mk := func(defaultSrc []string) *yongol.Fullstack {
		return &yongol.Fullstack{Manifest: &pm.ProjectConfig{Backend: pm.Backend{
			SecurityHeaders: &pm.SecurityHeadersConfig{CSP: &pm.CSPConfig{
				Directives: map[string][]string{"default-src": defaultSrc},
			}},
		}}}
	}
	cases := []TestSec301CspPermissiveCase{
		{name: "nil_fs", fs: nil, wantCount: 0},
		{name: "nil_manifest", fs: &yongol.Fullstack{}, wantCount: 0},
		{name: "nil_sh", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{}}, wantCount: 0},
		{name: "nil_csp", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{Backend: pm.Backend{SecurityHeaders: &pm.SecurityHeadersConfig{}}}}, wantCount: 0},
		{name: "no_default_src", fs: mk(nil), wantCount: 0},
		{name: "strict", fs: mk([]string{"'self'"}), wantCount: 0},
		{name: "wildcard", fs: mk([]string{"*"}), wantCount: 1},
		{name: "unsafe_eval", fs: mk([]string{"'unsafe-eval'"}), wantCount: 1},
		{name: "mixed_with_wildcard", fs: mk([]string{"'self'", "*"}), wantCount: 1},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runSec301CspPermissive(t, c)
		})
	}
}
