//ff:func feature=validate type=test control=iteration dimension=1 topic=manifest-auth
//ff:what sec402AccessTTLUpperBound — access_token_ttl 30분 상한 검증

package manifest

import (
	"testing"

	pm "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestSec402AccessTTLUpperBound(t *testing.T) {
	mk := func(ttl string) *yongol.Fullstack {
		return &yongol.Fullstack{Manifest: &pm.ProjectConfig{Backend: pm.Backend{Auth: &pm.Auth{AccessTokenTTL: ttl}}}}
	}
	cases := []TestSec402AccessTTLUpperBoundCase{
		{name: "nil_fs", fs: nil, wantCount: 0},
		{name: "nil_auth", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{}}, wantCount: 0},
		{name: "empty_ttl", fs: mk(""), wantCount: 0},
		{name: "invalid_duration", fs: mk("abc"), wantCount: 0},
		{name: "15m_ok", fs: mk("15m"), wantCount: 0},
		{name: "30m_ok", fs: mk("30m"), wantCount: 0},
		{name: "1h_warning", fs: mk("1h"), wantCount: 1},
		{name: "45m_warning", fs: mk("45m"), wantCount: 1},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runSec402AccessTTLUpperBound(t, c)
		})
	}
}
