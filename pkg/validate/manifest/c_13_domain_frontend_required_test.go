//ff:func feature=validate type=test control=iteration dimension=1 topic=manifest-structural
//ff:what C-13 테스트 — 도메인 frontend 누락 시 도메인별 ERROR, 모두 존재/미선언 시 무진단

package manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestC13DomainFrontendRequired(t *testing.T) {
	cases := []domainRuleCase{
		{"nil fs", nil, 0},
		{"nil manifest", &yongol.Fullstack{}, 0},
		{"no domains", fsWithDomains(nil), 0},
		{"all present", fsWithDomains(map[string]pmanifest.DomainConfig{
			"public": {Frontend: "frontend/public"},
			"admin":  {Frontend: "frontend/admin"},
		}), 0},
		{"one missing", fsWithDomains(map[string]pmanifest.DomainConfig{
			"public": {Frontend: "frontend/public"},
			"admin":  {Frontend: ""},
		}), 1},
		{"both missing", fsWithDomains(map[string]pmanifest.DomainConfig{
			"public": {},
			"admin":  {},
		}), 2},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runDomainRuleCase(t, c13DomainFrontendRequired, c, diagnostic.LevelError, "[C-13]")
		})
	}
}
