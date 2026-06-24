//ff:func feature=validate type=test control=iteration dimension=1 topic=manifest-structural
//ff:what C-14 테스트 — route_prefix 중복 시 ERROR, 고유/빈 값은 무진단

package manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestC14DomainRoutePrefixUnique(t *testing.T) {
	cases := []domainRuleCase{
		{"nil fs", nil, 0},
		{"nil manifest", &yongol.Fullstack{}, 0},
		{"no domains", fsWithDomains(nil), 0},
		{"unique", fsWithDomains(map[string]pmanifest.DomainConfig{
			"public": {RoutePrefix: "/api"},
			"admin":  {RoutePrefix: "/api/admin"},
		}), 0},
		{"empty ignored", fsWithDomains(map[string]pmanifest.DomainConfig{
			"public": {RoutePrefix: ""},
			"admin":  {RoutePrefix: ""},
		}), 0},
		{"duplicate pair", fsWithDomains(map[string]pmanifest.DomainConfig{
			"public": {RoutePrefix: "/api"},
			"admin":  {RoutePrefix: "/api"},
		}), 1},
		{"two duplicate groups", fsWithDomains(map[string]pmanifest.DomainConfig{
			"public": {RoutePrefix: "/api"},
			"admin":  {RoutePrefix: "/api"},
			"intA":   {RoutePrefix: "/x"},
			"intB":   {RoutePrefix: "/x"},
		}), 2},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runDomainRuleCase(t, c14DomainRoutePrefixUnique, c, diagnostic.LevelError, "[C-14]")
		})
	}
}
