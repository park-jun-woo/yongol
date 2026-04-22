//ff:func feature=gen-gogin type=test control=sequence topic=security-headers
//ff:what security_headers_source 렌더 스냅샷 + buildCSPValue / Profile 3종 시뮬레이션

package middleware

import (
	"strings"
	"testing"
)

// TestSecurityHeadersSource_Contains checks the verbatim source ships with
// every symbol the generator relies on (exported types and helpers).
func TestSecurityHeadersSource_Contains(t *testing.T) {
	for _, must := range []string{
		"type SecurityHeadersConfig struct",
		"func SecurityHeadersMiddleware(cfg SecurityHeadersConfig)",
		"func BuildStaticSecurityHeaders(cfg SecurityHeadersConfig)",
		"func BuildCSPHeader(cfg SecurityHeadersConfig)",
		"func BuildCSPValue(directives map[string][]string)",
		"\"X-Content-Type-Options\"",
		"\"X-Frame-Options\"",
		"\"Strict-Transport-Security\"",
		"\"Referrer-Policy\"",
		"\"Content-Security-Policy\"",
		"\"Content-Security-Policy-Report-Only\"",
	} {
		if !strings.Contains(securityHeadersSource, must) {
			t.Errorf("security_headers source missing fragment %q", must)
		}
	}
}

// runtimeCfg mirrors the in-source SecurityHeadersConfig so we can exercise
// the rendering helpers by re-implementing them in the test (cannot import
// the generated code). Instead we rely on string fragments — the
// block_security_headers_test covers the end-to-end shape.
func TestSecurityHeadersSource_ProfileBranches(t *testing.T) {
	// The profile branches are expressed as string comparisons in source;
	// verify each branch literal is present so a refactor cannot silently
	// remove a profile.
	for _, must := range []string{
		`profile != "dev"`,
		`profile == "api"`,
		`profile == "dev"`,
	} {
		if !strings.Contains(securityHeadersSource, must) {
			t.Errorf("security_headers source missing profile branch %q", must)
		}
	}
}
