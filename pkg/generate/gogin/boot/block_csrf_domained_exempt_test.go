//ff:func feature=gen-gogin type=test control=sequence topic=csrf
//ff:what blockCsrf(domained) — bearer 도메인 prefix 가 ExemptPaths 에 추가되는지 검증

package boot

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBlockCsrf_DomainedExemptsBearer(t *testing.T) {
	fs := &yongol.Fullstack{Manifest: &manifest.ProjectConfig{
		Backend: manifest.Backend{Module: "example.com/app", Auth: &manifest.Auth{Mode: "cookie"}},
		Domains: map[string]manifest.DomainConfig{
			"public": {RoutePrefix: "/api"},                           // cookie → CSRF checked
			"admin":  {RoutePrefix: "/api/admin", AuthMode: "bearer"}, // bearer → exempt
		},
	}}
	a := prepared.Auth{Present: true, Mode: "cookie", Raw: &manifest.Auth{Mode: "cookie"}}
	block := blockCsrf(fs, a, "example.com/app")
	body := strings.Join(block.Lines, "\n")
	if !strings.Contains(body, `ExemptPaths:      []string{"/api/admin"}`) {
		t.Errorf("bearer domain prefix must be exempt:\n%s", body)
	}
}
