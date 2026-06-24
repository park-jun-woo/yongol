//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what generateDomainAuth — bearer/cookie 도메인 파일명·함수명·토큰추출 우회(§3c) + 공유 헬퍼 미방출 검증

package auth

import (
	"go/format"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestGenerateDomainAuth(t *testing.T) {
	dir := t.TempDir()
	fs := &yongol.Fullstack{Manifest: &manifest.ProjectConfig{
		Backend: manifest.Backend{Module: "example.com/app", Auth: &manifest.Auth{Mode: "cookie"}},
		Domains: map[string]manifest.DomainConfig{
			"public": {RoutePrefix: "/api"},                           // cookie (inherits)
			"admin":  {RoutePrefix: "/api/admin", AuthMode: "bearer"}, // bearer
		},
	}}
	p := prepared.State{Auth: prepared.Auth{Present: true, Mode: "cookie"}}
	if err := generateDomainAuth(fs, p, dir, "example.com/app"); err != nil {
		t.Fatalf("generateDomainAuth: %v", err)
	}
	mwDir := filepath.Join(dir, "backend", "internal", "middleware")

	read := func(name string) string {
		b, err := os.ReadFile(filepath.Join(mwDir, name))
		if err != nil {
			t.Fatalf("read %s: %v", name, err)
		}
		return string(b)
	}

	bearer := read("bearerauth_admin.go")
	if !strings.Contains(bearer, "func BearerAuthStrictAdmin(publicOps map[string]bool) api_admin.StrictMiddlewareFunc") {
		t.Errorf("bearerauth_admin.go missing suffixed func / api_admin type:\n%s", bearer)
	}
	if strings.Contains(bearer, "extractToken(") || strings.Contains(bearer, "authMode(") {
		t.Errorf("bearer domain must bypass global extractToken/authMode (§3c):\n%s", bearer)
	}
	if !strings.Contains(bearer, `ctx.GetHeader("Authorization")`) {
		t.Errorf("bearer domain must inline Authorization header parse:\n%s", bearer)
	}
	if !strings.Contains(bearer, `"example.com/app/internal/api_admin"`) {
		t.Errorf("bearer domain must import api_admin:\n%s", bearer)
	}

	cookie := read("cookieauth_public.go")
	if !strings.Contains(cookie, "func CookieAuthStrictPublic(publicOps map[string]bool) api_public.StrictMiddlewareFunc") {
		t.Errorf("cookieauth_public.go missing suffixed func / api_public type:\n%s", cookie)
	}
	if !strings.Contains(cookie, "auth.ExtractAccessFromCookie(ctx)") {
		t.Errorf("cookie domain must read cookie directly (§3c):\n%s", cookie)
	}
	if strings.Contains(cookie, "extractToken(") {
		t.Errorf("cookie domain must not call extractToken (§3c):\n%s", cookie)
	}

	// Shared helpers must NOT be emitted in domain mode (§3d).
	for _, gone := range []string{"auth_mode.go", "extract_token.go", "bearerauth.go"} {
		if _, err := os.Stat(filepath.Join(mwDir, gone)); err == nil {
			t.Errorf("domain mode must not emit shared/fixed file %q", gone)
		}
	}

	// Emitted sources must be gofmt-clean (template + substitution well-formed).
	for _, name := range []string{"bearerauth_admin.go", "cookieauth_public.go"} {
		if _, err := format.Source([]byte(read(name))); err != nil {
			t.Errorf("%s not gofmt-parseable: %v", name, err)
		}
	}
}
