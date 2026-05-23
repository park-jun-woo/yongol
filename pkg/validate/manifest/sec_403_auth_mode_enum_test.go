//ff:func feature=validate type=test control=iteration dimension=1 topic=manifest-auth
//ff:what sec403AuthModeEnum — auth.mode가 cookie/bearer/hybrid 중 하나인지 검증

package manifest

import (
	"testing"

	pm "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestSec403AuthModeEnum(t *testing.T) {
	mk := func(mode string) *yongol.Fullstack {
		return &yongol.Fullstack{Manifest: &pm.ProjectConfig{Backend: pm.Backend{Auth: &pm.Auth{Mode: mode}}}}
	}
	cases := []TestSec403AuthModeEnumCase{
		{name: "nil_fs", fs: nil, wantCount: 0},
		{name: "nil_auth", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{}}, wantCount: 0},
		{name: "empty_mode_ok", fs: mk(""), wantCount: 0},
		{name: "cookie", fs: mk("cookie"), wantCount: 0},
		{name: "bearer", fs: mk("bearer"), wantCount: 0},
		{name: "hybrid", fs: mk("hybrid"), wantCount: 0},
		{name: "invalid", fs: mk("jwt"), wantCount: 1},
		{name: "typo", fs: mk("cookies"), wantCount: 1},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runSec403AuthModeEnum(t, c)
		})
	}
}
