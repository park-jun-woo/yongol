//ff:func feature=gen-gogin type=test control=sequence topic=csrf
//ff:what csrf_source 렌더 스냅샷 — double-submit cookie 골격 존재 확인

package middleware

import (
	"strings"
	"testing"
)

func TestCsrfSourceTemplate_ContainsDoubleSubmitFragments(t *testing.T) {
	out := csrfSourceTemplate
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
		"CSRF 토큰이 유효하지 않습니다",
		"HybridBearerSkip",
		"hasBearerHeader",
		`"Authorization"`,
	} {
		if !strings.Contains(out, must) {
			t.Errorf("csrfSourceTemplate missing fragment %q", must)
		}
	}
	// No placeholders should leak.
	if strings.Contains(out, "__MODULE__") {
		t.Errorf("csrfSourceTemplate should not carry unresolved placeholders")
	}
}
