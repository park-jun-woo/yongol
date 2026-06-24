//ff:func feature=gen-gogin type=test control=sequence
//ff:what domainWireMode — override / backend 상속 / bearer fallback 분기 검증

package boot

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestDomainWireMode(t *testing.T) {
	mkFS := func(auth *manifest.Auth, mode string) *yongol.Fullstack {
		return &yongol.Fullstack{Manifest: &manifest.ProjectConfig{
			Backend: manifest.Backend{Auth: auth},
			Domains: map[string]manifest.DomainConfig{
				"d": {RoutePrefix: "/api", AuthMode: mode},
			},
		}}
	}
	if got := domainWireMode(mkFS(&manifest.Auth{Mode: "cookie"}, "bearer"), "d"); got != "bearer" {
		t.Errorf("override → bearer, got %q", got)
	}
	if got := domainWireMode(mkFS(&manifest.Auth{Mode: "cookie"}, ""), "d"); got != "cookie" {
		t.Errorf("inherit backend → cookie, got %q", got)
	}
	if got := domainWireMode(mkFS(nil, ""), "d"); got != "bearer" {
		t.Errorf("no auth block → bearer fallback, got %q", got)
	}
}
