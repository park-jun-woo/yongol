//ff:func feature=validate type=test control=iteration dimension=1 topic=manifest-structural
//ff:what C-16 테스트 — 도메인 frontend 가 단일 사이트 STML 루트와 충돌 시 WARNING

package manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestC16DomainFrontendConflict(t *testing.T) {
	cases := []domainRuleCase{
		{"nil fs", nil, 0},
		{"nil manifest", &yongol.Fullstack{}, 0},
		{"no domains", fsWithDomains(nil), 0},
		{"dedicated dirs", fsWithDomains(map[string]pmanifest.DomainConfig{
			"public": {Frontend: "frontend/public"},
			"admin":  {Frontend: "frontend/admin"},
		}), 0},
		{"empty skipped", fsWithDomains(map[string]pmanifest.DomainConfig{
			"public": {Frontend: ""},
			"admin":  {Frontend: "frontend/admin"},
		}), 0},
		{"one collides", fsWithDomains(map[string]pmanifest.DomainConfig{
			"public": {Frontend: "frontend"},
			"admin":  {Frontend: "frontend/admin"},
		}), 1},
		{"unclean path collides", fsWithDomains(map[string]pmanifest.DomainConfig{
			"public": {Frontend: "./frontend/"},
			"admin":  {Frontend: "frontend/admin"},
		}), 1},
		{"both collide", fsWithDomains(map[string]pmanifest.DomainConfig{
			"public": {Frontend: "frontend"},
			"admin":  {Frontend: "frontend"},
		}), 2},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runDomainRuleCase(t, c16DomainFrontendConflict, c, diagnostic.LevelWarning, "[C-16]")
		})
	}
}
