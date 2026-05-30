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

// TestCallGeminiRequestPath drives the body-marshalling and http.Post path by
// supplying a fake API key. The endpoint is hardcoded to googleapis.com, so in
// a hermetic test environment the POST fails (DNS/network), exercising the
// "gemini request: %w" error branch after the request payload is built.
func TestCallGeminiRequestPath(t *testing.T) {
	t.Setenv("GEMINI_API_KEY", "fake-key-for-test")

	out, err := callGemini("gemini-2.0-flash", "system prompt", "user prompt")
	if err == nil {
		t.Skip("network reachable; gemini endpoint responded — request branch not exercised")
	}
	if out != "" {
		t.Errorf("expected empty output on error, got %q", out)
	}
	if !strings.Contains(err.Error(), "gemini") {
		t.Errorf("expected gemini-related error, got: %v", err)
	}
}
