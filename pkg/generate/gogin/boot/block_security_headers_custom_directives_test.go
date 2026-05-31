//ff:func feature=gen-gogin type=test control=sequence topic=security-headers
//ff:what TestBlockSecurityHeaders_CustomDirectives — manifest CSP directives 가 기본값을 대체

package boot

import (
	"strings"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestBlockSecurityHeaders_CustomDirectives ensures CSP directives from the
// manifest override the built-in defaults.
func TestBlockSecurityHeaders_CustomDirectives(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Module: "example.com/zenflow",
				SecurityHeaders: &pmanifest.SecurityHeadersConfig{
					CSP: &pmanifest.CSPConfig{
						Directives: map[string][]string{
							"default-src": {"'self'"},
							"script-src":  {"'self'", "cdn.example.com"},
						},
					},
				},
			},
		},
	}
	block := blockSecurityHeaders(fs, "example.com/zenflow")
	combined := strings.Join(block.Lines, "\n") + "\n" + strings.Join(block.Funcs, "\n")
	if !strings.Contains(combined, `"script-src": []string{"'self'", "cdn.example.com"}`) {
		t.Fatalf("custom script-src not emitted; got:\n%s", combined)
	}
	// custom directives map replaces defaults entirely — frame-ancestors
	// should NOT appear since the custom map omitted it.
	if strings.Contains(combined, `"frame-ancestors"`) {
		t.Fatalf("custom directives should replace defaults; frame-ancestors leaked:\n%s", combined)
	}
}
