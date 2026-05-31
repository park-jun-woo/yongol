//ff:func feature=agent type=test control=sequence
//ff:what TestCallGemini — API 키 미존재 시 HTTP 호출 전 조기 에러 반환 검증
package agent

import (
	"strings"
	"testing"
)

func TestCallGemini(t *testing.T) {
	// No GEMINI_API_KEY and an XDG dir without credentials.yaml forces the
	// loadAPIKey failure so callGemini returns before any network call.
	t.Setenv("GEMINI_API_KEY", "")
	t.Setenv("XDG_CONFIG_HOME", t.TempDir())

	out, err := callGemini("gemini-2.0-flash", "system", "user")
	if err == nil {
		t.Fatal("expected error when API key is unavailable")
	}
	if out != "" {
		t.Errorf("expected empty output on error, got %q", out)
	}
	if !strings.Contains(err.Error(), "gemini") {
		t.Errorf("expected gemini-related key error, got: %v", err)
	}
}
