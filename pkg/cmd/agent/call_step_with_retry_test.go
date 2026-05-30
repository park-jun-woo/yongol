//ff:func feature=agent type=test control=sequence
//ff:what TestCallStepWithRetry — 미지원 backend 로 첫 호출+재시도 모두 실패하는 에러 경로 검증

package agent

import (
	"strings"
	"testing"
)

// TestCallStepWithRetry drives the failure path: an unsupported backend makes
// llmCall (and thus llmCallWithNumCtx) return an error on both the initial call
// and the single retry, so callStepWithRetry surfaces the retry error.
//
// The success and empty-content branches require a live LLM endpoint and are
// not deterministically reachable in a hermetic unit test.
func TestCallStepWithRetry(t *testing.T) {
	cfg := Config{Backend: "unsupported-backend", Model: "none"}

	out, err := callStepWithRetry(cfg, "do something")
	if err == nil {
		t.Fatal("expected error when backend is unsupported")
	}
	if out != "" {
		t.Errorf("expected empty output on error, got %q", out)
	}
	if !strings.Contains(err.Error(), "unsupported backend") {
		t.Errorf("expected unsupported-backend error, got: %v", err)
	}
}
