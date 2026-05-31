//ff:func feature=gen-gogin type=test control=sequence topic=security-headers
//ff:what TestBlockSecurityHeaders_HSTSOverride — HSTS max_age/preload/subs override 확인

package boot

import (
	"strings"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestBlockSecurityHeaders_HSTSOverride ensures HSTS sub-fields are honored.
func TestBlockSecurityHeaders_HSTSOverride(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Module: "example.com/zenflow",
				SecurityHeaders: &pmanifest.SecurityHeadersConfig{
					HSTS: &pmanifest.HSTSConfig{
						MaxAge:            63072000,
						IncludeSubDomains: false,
						Preload:           true,
					},
				},
			},
		},
	}
	block := blockSecurityHeaders(fs, "example.com/zenflow")
	combined := strings.Join(block.Lines, "\n") + "\n" + strings.Join(block.Funcs, "\n")
	if !strings.Contains(combined, `63072000`) {
		t.Fatalf("custom HSTS max_age not emitted; got:\n%s", combined)
	}
	if !strings.Contains(combined, `HSTSPreload:       true`) {
		t.Fatalf("HSTS preload=true not emitted; got:\n%s", combined)
	}
	if !strings.Contains(combined, `HSTSIncludeSubs:   false`) {
		t.Fatalf("HSTS include_subdomains=false not emitted; got:\n%s", combined)
	}
}
