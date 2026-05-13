//ff:func feature=gen-gogin type=test control=iteration dimension=1 topic=security-headers
//ff:what TestBlockSecurityHeaders_DefaultActive — 기본 manifest 에서도 production preset 방출

package boot

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

// TestBlockSecurityHeaders_DefaultActive verifies that a manifest without
// any security_headers block still emits the production preset.
func TestBlockSecurityHeaders_DefaultActive(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{Module: "example.com/zenflow"},
		},
	}
	block := blockSecurityHeaders(fs, "example.com/zenflow")
	if block.Name != "security-headers" {
		t.Fatalf("unexpected name %q", block.Name)
	}
	body := strings.Join(block.Lines, "\n")
	funcs := strings.Join(block.Funcs, "\n")
	combined := body + "\n" + funcs
	for _, must := range []string{
		`middleware.SecurityHeadersMiddleware(buildSecurityHeadersCfg(`,
		`BACKEND_SECURITY_HEADERS_ENABLED`,
		`BACKEND_SECURITY_HEADERS_PROFILE`,
		`BACKEND_SECURITY_HEADERS_HSTS_MAX_AGE`,
		`BACKEND_SECURITY_HEADERS_CSP_REPORT_ONLY`,
		`"production"`,
		`31536000`,
		`"DENY"`,
		`"strict-origin-when-cross-origin"`,
		`"default-src": []string{"'self'"}`,
		`"frame-ancestors": []string{"'none'"}`,
		`"base-uri": []string{"'self'"}`,
		`"camera": []string{}`,
	} {
		if !strings.Contains(combined, must) {
			t.Errorf("default block missing fragment %q; got:\n%s", must, combined)
		}
	}
}
