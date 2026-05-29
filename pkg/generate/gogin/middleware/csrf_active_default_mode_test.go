//ff:func feature=gen-gogin type=test control=sequence topic=csrf
//ff:what TestCsrfActive_DefaultMode_BUG009 — auth.mode 생략 시 기본 cookie 해석으로 csrf 미들웨어 활성

package middleware

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestCsrfActive_DefaultMode_BUG009 pins the BUG-009 regression: a
// manifest that omits backend.auth.mode must still emit the csrf
// middleware file, because prepared.Auth.Mode defaults to "cookie".
func TestCsrfActive_DefaultMode_BUG009(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Auth: &pmanifest.Auth{}, // Mode intentionally blank — defaults to "cookie"
			},
		},
	}
	p := prepared.New(fs)
	if p.Auth.Mode != "cookie" {
		t.Fatalf("prepared.Auth.Mode=%q; expected %q", p.Auth.Mode, "cookie")
	}
	if !csrfActive(p.Auth) {
		t.Fatalf("csrfActive must return true when Mode defaults to cookie (BUG-009 regression)")
	}
}
