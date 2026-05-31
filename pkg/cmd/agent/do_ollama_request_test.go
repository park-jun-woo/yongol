//ff:func feature=agent type=test control=sequence
//ff:what TestDoOllamaRequest — 직렬화 불가 body 및 연결 불가 URL에서 에러 반환 검증
package agent

import (
	"strings"
	"testing"
)

func TestDoOllamaRequest(t *testing.T) {
	// A channel cannot be JSON-marshalled, exercising the marshal-error branch
	// before any network call.
	if _, err := doOllamaRequest("http://127.0.0.1:1/x", make(chan int)); err == nil {
		t.Fatal("expected marshal error for non-serialisable body")
	} else if !strings.Contains(err.Error(), "marshal ollama request") {
		t.Errorf("err = %v, want marshal error", err)
	}

	// Valid body but an unroutable port fails the HTTP POST (connection refused).
	out, err := doOllamaRequest("http://127.0.0.1:1/api/chat", map[string]any{"model": "x"})
	if err == nil {
		t.Fatal("expected request error against unreachable server")
	}
	if out != "" {
		t.Errorf("expected empty output on error, got %q", out)
	}
}
