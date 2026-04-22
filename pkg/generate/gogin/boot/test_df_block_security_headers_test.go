//ff:func feature=gen-gogin type=test control=sequence topic=security-headers
//ff:what blockSecurityHeaders 활성/비활성/profile 분기 + CSP directive 렌더 테스트

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
	for _, must := range []string{
		`middleware.SecurityHeadersMiddleware(secHeadersCfg)`,
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
		if !strings.Contains(body, must) {
			t.Errorf("default block missing fragment %q; got:\n%s", must, body)
		}
	}
}

// TestBlockSecurityHeaders_DisabledExplicitly ensures Enabled=false yields
// an inert block with an Active guard that keeps it out of main.go.
func TestBlockSecurityHeaders_DisabledExplicitly(t *testing.T) {
	disabled := false
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Module: "example.com/zenflow",
				SecurityHeaders: &pmanifest.SecurityHeadersConfig{
					Enabled: &disabled,
				},
			},
		},
	}
	block := blockSecurityHeaders(fs, "example.com/zenflow")
	if block.Active == nil || block.Active(fs) {
		t.Fatalf("disabled block must not be active")
	}
	if len(block.Lines) != 0 {
		t.Fatalf("disabled block should emit no lines, got %+v", block.Lines)
	}
}

// TestBlockSecurityHeaders_DevProfile verifies the dev profile is wired
// through unchanged — runtime middleware decides what to do with it.
func TestBlockSecurityHeaders_DevProfile(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Module: "example.com/zenflow",
				SecurityHeaders: &pmanifest.SecurityHeadersConfig{
					Profile: "dev",
				},
			},
		},
	}
	body := strings.Join(blockSecurityHeaders(fs, "example.com/zenflow").Lines, "\n")
	if !strings.Contains(body, `"dev"`) {
		t.Fatalf("dev profile not emitted; got:\n%s", body)
	}
}

// TestBlockSecurityHeaders_APIProfile — api profile just propagates the
// string, runtime drops CSP.
func TestBlockSecurityHeaders_APIProfile(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Module: "example.com/zenflow",
				SecurityHeaders: &pmanifest.SecurityHeadersConfig{
					Profile: "api",
				},
			},
		},
	}
	body := strings.Join(blockSecurityHeaders(fs, "example.com/zenflow").Lines, "\n")
	if !strings.Contains(body, `"api"`) {
		t.Fatalf("api profile not emitted; got:\n%s", body)
	}
}

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
	body := strings.Join(blockSecurityHeaders(fs, "example.com/zenflow").Lines, "\n")
	if !strings.Contains(body, `"script-src": []string{"'self'", "cdn.example.com"}`) {
		t.Fatalf("custom script-src not emitted; got:\n%s", body)
	}
	// custom directives map replaces defaults entirely — frame-ancestors
	// should NOT appear since the custom map omitted it.
	if strings.Contains(body, `"frame-ancestors"`) {
		t.Fatalf("custom directives should replace defaults; frame-ancestors leaked:\n%s", body)
	}
}

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
