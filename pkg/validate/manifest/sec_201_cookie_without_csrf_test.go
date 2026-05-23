//ff:func feature=validate type=test control=iteration dimension=1 topic=manifest-auth
//ff:what sec201CookieWithoutCsrf — cookie/hybrid mode에서 csrf.enabled=false 금지 검증

package manifest

import (
	"testing"

	pm "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestSec201CookieWithoutCsrf(t *testing.T) {
	cases := []TestSec201CookieWithoutCsrfCase{
		{name: "nil_manifest", fs: &yongol.Fullstack{}, wantCount: 0},
		{name: "nil_auth", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{}}, wantCount: 0},
		{name: "bearer_mode", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{Backend: pm.Backend{Auth: &pm.Auth{Mode: "bearer"}}}}, wantCount: 0},
		{name: "cookie_nil_csrf", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{Backend: pm.Backend{Auth: &pm.Auth{Mode: "cookie"}}}}, wantCount: 0},
		{name: "cookie_csrf_enabled", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{Backend: pm.Backend{Auth: &pm.Auth{Mode: "cookie", Csrf: &pm.CsrfConfig{Enabled: true}}}}}, wantCount: 0},
		{name: "cookie_csrf_disabled", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{Backend: pm.Backend{Auth: &pm.Auth{Mode: "cookie", Csrf: &pm.CsrfConfig{Enabled: false}}}}}, wantCount: 1},
		{name: "hybrid_csrf_disabled", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{Backend: pm.Backend{Auth: &pm.Auth{Mode: "hybrid", Csrf: &pm.CsrfConfig{Enabled: false}}}}}, wantCount: 1},
		{name: "default_mode_csrf_disabled", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{Backend: pm.Backend{Auth: &pm.Auth{Csrf: &pm.CsrfConfig{Enabled: false}}}}}, wantCount: 1},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runSec201CookieWithoutCsrf(t, c)
		})
	}
}
