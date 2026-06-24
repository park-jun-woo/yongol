//ff:func feature=validate type=test control=iteration dimension=1 topic=manifest-structural
//ff:what C-17 테스트 — domains 1개 선언 시 ERROR, 2개 이상/미선언은 무진단

package manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestC17DomainMinimumTwo(t *testing.T) {
	cases := []domainRuleCase{
		{"nil fs", nil, 0},
		{"nil manifest", &yongol.Fullstack{}, 0},
		{"no domains", fsWithDomains(nil), 0},
		{"single domain", fsWithDomains(map[string]pmanifest.DomainConfig{
			"public": {OpenAPI: "api/public.yaml"},
		}), 1},
		{"two domains", fsWithDomains(map[string]pmanifest.DomainConfig{
			"public": {OpenAPI: "api/public.yaml"},
			"admin":  {OpenAPI: "api/admin.yaml"},
		}), 0},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runDomainRuleCase(t, c17DomainMinimumTwo, c, diagnostic.LevelError, "[C-17]")
		})
	}
}
