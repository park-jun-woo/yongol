//ff:func feature=gen-gogin type=test control=sequence topic=csrf
//ff:what TestCsrfSourceTemplate_BearerDefaultGatesOff — bearer 빌드는 csrfAuthMode 폴백 "bearer" 주입 (BUG-116)

package middleware

import (
	"fmt"
	"strings"
	"testing"
)

// TestCsrfSourceTemplate_BearerDefaultGatesOff verifies the build-time
// default mode is injected so that, with BACKEND_AUTH_MODE unset, a bearer
// build's csrfAuthMode() falls back to "bearer" (CSRF no-op).
func TestCsrfSourceTemplate_BearerDefaultGatesOff(t *testing.T) {
	out := fmt.Sprintf(csrfSourceTemplate, "bearer")
	if !strings.Contains(out, `return "bearer"`) {
		t.Errorf("bearer build must inject \"bearer\" as csrfAuthMode fallback:\n%s", out)
	}
}
