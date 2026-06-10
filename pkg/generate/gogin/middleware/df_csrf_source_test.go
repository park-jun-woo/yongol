//ff:func feature=gen-gogin type=test control=iteration dimension=1 topic=csrf
//ff:what csrf_source render snapshot — double-submit cookie skeleton + BUG-116 런타임 게이트 검증

package middleware

import (
	"fmt"
	"strings"
	"testing"
)

func TestCsrfSourceTemplate_ContainsDoubleSubmitFragments(t *testing.T) {
	out := fmt.Sprintf(csrfSourceTemplate, "cookie")
	for _, must := range []string{
		"func Csrf(cfg CsrfConfig) gin.HandlerFunc",
		"type CsrfConfig struct",
		"type CsrfEnvelope struct",
		"func generateCsrfToken()",
		"base64.RawURLEncoding.EncodeToString",
		"func isSafeMethod(",
		"func isExemptPath(",
		"func constantTimeEqual(",
		"subtle.ConstantTimeCompare",
		"http.MethodGet, http.MethodHead, http.MethodOptions",
		"csrf_token_invalid",
		"CSRF token is invalid",
		"HybridBearerSkip",
		"hasBearerHeader",
		`"Authorization"`,
		// BUG-116 / Phase-B1 — runtime auth-mode gate.
		"func csrfAuthMode() string",
		"func csrfRuntimeActive() bool",
		"BACKEND_AUTH_MODE",
		"if !csrfRuntimeActive() {",
		`return "cookie"`, // build-time default injected via %q
	} {
		if !strings.Contains(out, must) {
			t.Errorf("csrfSourceTemplate missing fragment %q", must)
		}
	}
	// No placeholders should leak, and the single %q verb must be consumed.
	if strings.Contains(out, "__MODULE__") || strings.Contains(out, "%!q") || strings.Contains(out, "%q") {
		t.Errorf("csrfSourceTemplate should not carry unresolved placeholders:\n%s", out)
	}
}
