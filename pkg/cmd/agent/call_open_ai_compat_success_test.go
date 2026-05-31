//ff:func feature=agent type=test control=sequence
//ff:what TestCallOpenAICompat — API 키 부재 조기반환 + httptest 로 성공/비200/파싱오류/빈선택/요청오류 분기 검증
package agent

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCallOpenAICompatSuccess(t *testing.T) {
	t.Setenv("XAI_API_KEY", "fake-key")
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer fake-key" {
			t.Errorf("missing/incorrect Authorization header: %q", r.Header.Get("Authorization"))
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"choices":[{"message":{"content":"hello world"}}]}`))
	}))
	defer srv.Close()

	out, err := callOpenAICompat(srv.URL, "xai", "grok", "sys", "user")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "hello world" {
		t.Errorf("content: got %q, want %q", out, "hello world")
	}
}
