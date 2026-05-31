//ff:func feature=gen-gogin type=test control=sequence topic=security-headers
//ff:what resolveSecurityHeaders — manifest.security_headers + production 기본값 병합
package boot

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestResolveSecurityHeaders_DirectivesAndPermissions(t *testing.T) {
	cspOn := true
	fs := &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{
		Backend: pmanifest.Backend{SecurityHeaders: &pmanifest.SecurityHeadersConfig{
			CSP: &pmanifest.CSPConfig{
				Enabled:    &cspOn,
				Directives: map[string][]string{"script-src": {"'self'", "cdn.example.com"}},
			},
			PermissionsPolicy: map[string][]string{"camera": {"'self'"}},
		}},
	}}
	out := resolveSecurityHeaders(fs)
	if got := out.CSPDirectives["script-src"]; len(got) != 2 || got[0] != "'self'" {
		t.Errorf("csp directives override wrong: %+v", out.CSPDirectives)
	}
	// custom directives replace the defaults entirely.
	if _, ok := out.CSPDirectives["default-src"]; ok {
		t.Errorf("custom directives should replace defaults, got: %+v", out.CSPDirectives)
	}
	if got := out.PermissionsPolicy["camera"]; len(got) != 1 || got[0] != "'self'" {
		t.Errorf("permissions policy override wrong: %+v", out.PermissionsPolicy)
	}
}
