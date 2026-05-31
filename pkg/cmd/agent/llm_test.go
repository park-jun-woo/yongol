//ff:func feature=agent type=test control=sequence
//ff:what TestLLMCall — 미지원 backend는 즉시 에러, 네트워크 호출 backend는 조기 실패 경로 검증
package agent

import (
	"strings"
	"testing"
)

func TestLLMCall(t *testing.T) {
	// Unsupported backend returns an error before any network call.
	out, err := llmCall("unknown-backend", "m", "sys", "user")
	if err == nil {
		t.Fatal("expected error for unsupported backend")
	}
	if out != "" {
		t.Errorf("expected empty output, got %q", out)
	}
	if !strings.Contains(err.Error(), "unsupported backend") {
		t.Errorf("error = %v, want mention of unsupported backend", err)
	}

	// gemini path fails fast without credentials (no network call made).
	t.Setenv("GEMINI_API_KEY", "")
	t.Setenv("XDG_CONFIG_HOME", t.TempDir())
	if _, err := llmCall("gemini", "gemini-2.0-flash", "sys", "user"); err == nil {
		t.Error("expected error when gemini credentials are unavailable")
	}

	// xai path fails fast without credentials (loadAPIKey error before network).
	t.Setenv("XAI_API_KEY", "")
	if _, err := llmCall("xai", "grok", "sys", "user"); err == nil {
		t.Error("expected error when xai credentials are unavailable")
	}

	// ollama path routes to callOllama; with no local server reachable the
	// request fails. Either an error or (rarely) a live response is acceptable —
	// the goal is to exercise the ollama switch arm.
	_, _ = llmCall("ollama", "llama3", "sys", "user")
}
