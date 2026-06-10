//ff:func feature=validate type=test control=iteration dimension=1 topic=manifest-auth
//ff:what sec202RuntimeModeCsrf — bearer mode + csrf.enabled=false 런타임 CSRF 게이트 상실 경고 검증

package manifest

import (
	"testing"

	pm "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestSec202RuntimeModeCsrf(t *testing.T) {
	cases := []TestSec202RuntimeModeCsrfCase{
		{name: "nil_fs", fs: nil, wantCount: 0},
		{name: "nil_manifest", fs: &yongol.Fullstack{}, wantCount: 0},
		{name: "nil_auth", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{}}, wantCount: 0},
		// bearer with the runtime-gated CSRF middleware emitted (Phase-B1
		// default) — contract holds, no diagnostic.
		{name: "bearer_nil_csrf", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{Backend: pm.Backend{Auth: &pm.Auth{Mode: "bearer"}}}}, wantCount: 0},
		{name: "bearer_csrf_enabled", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{Backend: pm.Backend{Auth: &pm.Auth{Mode: "bearer", Csrf: &pm.CsrfConfig{Enabled: true}}}}}, wantCount: 0},
		// bearer + explicit opt-out — the only config that loses the
		// runtime gate: BACKEND_AUTH_MODE=cookie|hybrid would run cookie
		// auth without CSRF.
		{name: "bearer_csrf_disabled", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{Backend: pm.Backend{Auth: &pm.Auth{Mode: "bearer", Csrf: &pm.CsrfConfig{Enabled: false}}}}}, wantCount: 1},
		// cookie/hybrid (and the unspecified-mode cookie default, incl.
		// jwt-typed) with csrf disabled are SEC-201's ERROR domain.
		{name: "cookie_csrf_disabled_sec201_domain", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{Backend: pm.Backend{Auth: &pm.Auth{Mode: "cookie", Csrf: &pm.CsrfConfig{Enabled: false}}}}}, wantCount: 0},
		{name: "hybrid_csrf_disabled_sec201_domain", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{Backend: pm.Backend{Auth: &pm.Auth{Mode: "hybrid", Csrf: &pm.CsrfConfig{Enabled: false}}}}}, wantCount: 0},
		{name: "default_mode_csrf_disabled_sec201_domain", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{Backend: pm.Backend{Auth: &pm.Auth{Csrf: &pm.CsrfConfig{Enabled: false}}}}}, wantCount: 0},
		{name: "jwt_type_default_mode_csrf_disabled_sec201_domain", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{Backend: pm.Backend{Auth: &pm.Auth{Type: "jwt", Csrf: &pm.CsrfConfig{Enabled: false}}}}}, wantCount: 0},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runSec202RuntimeModeCsrf(t, c)
		})
	}
}
