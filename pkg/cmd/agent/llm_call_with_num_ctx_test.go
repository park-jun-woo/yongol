//ff:func feature=agent type=test control=sequence
//ff:what TestLLMCallWithNumCtx — 비-ollama backend는 표준 호출로 위임(미지원 backend 에러)되는지 검증

package agent

import (
	"strings"
	"testing"
)

func TestLLMCallWithNumCtx(t *testing.T) {
	// Non-ollama backend ignores numCtx and delegates to llmCall, which
	// rejects an unsupported backend before any network activity.
	_, err := llmCallWithNumCtx("unknown-backend", "m", "sys", "user", 4096)
	if err == nil || !strings.Contains(err.Error(), "unsupported backend") {
		t.Fatalf("err = %v, want unsupported backend error", err)
	}

	// numCtx <= 0 also delegates to the standard llmCall path.
	if _, err := llmCallWithNumCtx("unknown-backend", "m", "sys", "user", 0); err == nil {
		t.Error("expected error for unsupported backend with numCtx=0")
	}

	// ollama with a positive numCtx routes to callOllamaWithCtx. With no local
	// Ollama server reachable the call fails — the branch is still exercised.
	_, _ = llmCallWithNumCtx("ollama", "llama3", "sys", "user", 4096)
}
