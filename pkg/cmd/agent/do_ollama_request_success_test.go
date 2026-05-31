//ff:func feature=agent type=test control=sequence
//ff:what TestDoOllamaRequest — 직렬화 불가 body 및 연결 불가 URL에서 에러 반환 검증
package agent

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDoOllamaRequestSuccess(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"message":{"content":"ollama reply"}}`))
	}))
	defer srv.Close()

	out, err := doOllamaRequest(srv.URL, map[string]any{"model": "x"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "ollama reply" {
		t.Errorf("content: got %q, want %q", out, "ollama reply")
	}
}
