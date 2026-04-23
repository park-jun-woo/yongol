//ff:func feature=gen-gogin type=test control=sequence topic=security-headers
//ff:what TestBlockSecurityHeaders_HSTSOverride — HSTS max_age/preload/subs override 확인

package boot

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
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
	body := strings.Join(blockSecurityHeaders(fs, "example.com/zenflow").Lines, "\n")
	if !strings.Contains(body, `63072000`) {
		t.Fatalf("custom HSTS max_age not emitted; got:\n%s", body)
	}
	if !strings.Contains(body, `HSTSPreload:       true`) {
		t.Fatalf("HSTS preload=true not emitted; got:\n%s", body)
	}
	if !strings.Contains(body, `HSTSIncludeSubs:   false`) {
		t.Fatalf("HSTS include_subdomains=false not emitted; got:\n%s", body)
	}
}
