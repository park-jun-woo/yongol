//ff:func feature=agent type=test control=sequence
//ff:what TestCallGemini — API 키 미존재 시 HTTP 호출 전 조기 에러 반환 검증
package agent

import (
	"strings"
	"testing"
)

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
