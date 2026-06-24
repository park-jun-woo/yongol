//ff:func feature=validate type=test control=iteration dimension=1 topic=manifest-auth
//ff:what C-15 테스트 — 도메인 auth_mode enum(validAuthModes 재사용) 검증

package manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestC15DomainAuthModeEnum(t *testing.T) {
	cases := []domainRuleCase{
		{"nil fs", nil, 0},
		{"nil manifest", &yongol.Fullstack{}, 0},
		{"no domains", fsWithDomains(nil), 0},
		{"valid modes", fsWithDomains(map[string]pmanifest.DomainConfig{
			"public": {AuthMode: "cookie"},
			"admin":  {AuthMode: "bearer"},
			"hyb":    {AuthMode: "hybrid"},
		}), 0},
		{"empty inherits", fsWithDomains(map[string]pmanifest.DomainConfig{
			"public": {AuthMode: ""},
			"admin":  {AuthMode: "bearer"},
		}), 0},
		{"one unknown", fsWithDomains(map[string]pmanifest.DomainConfig{
			"public": {AuthMode: "cookie"},
			"admin":  {AuthMode: "jwt"},
		}), 1},
		{"both unknown", fsWithDomains(map[string]pmanifest.DomainConfig{
			"public": {AuthMode: "cookies"},
			"admin":  {AuthMode: "session"},
		}), 2},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runDomainRuleCase(t, c15DomainAuthModeEnum, c, diagnostic.LevelError, "[C-15]")
		})
	}
}
