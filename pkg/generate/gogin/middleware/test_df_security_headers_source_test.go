//ff:func feature=gen-gogin type=test control=iteration dimension=1 topic=security-headers
//ff:what TestSecurityHeadersSource_Contains — 수기 템플릿 상수 내 주요 심볼 존재 확인

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
