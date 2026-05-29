//ff:func feature=agent type=test control=sequence
//ff:what TestCallOpenAICompat — API 키 미존재 시 HTTP 호출 전 조기 에러 반환 검증

package agent

import "testing"

func TestCallOpenAICompat(t *testing.T) {
	// No XAI_API_KEY and an empty XDG dir forces loadAPIKey to fail, so the
	// function returns before constructing or sending any HTTP request.
	t.Setenv("XAI_API_KEY", "")
	t.Setenv("XDG_CONFIG_HOME", t.TempDir())

	out, err := callOpenAICompat("https://api.x.ai/v1/chat/completions", "xai", "grok", "sys", "user")
	if err == nil {
		t.Fatal("expected error when API key is unavailable")
	}
	if out != "" {
		t.Errorf("expected empty output on error, got %q", out)
	}
}
