//ff:func feature=agent type=test control=sequence
//ff:what TestCallOpenAICompat — API 키 부재 조기반환 + httptest 로 성공/비200/파싱오류/빈선택/요청오류 분기 검증
package agent

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCallOpenAICompatNon200(t *testing.T) {
	t.Setenv("XAI_API_KEY", "fake-key")
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`bad request`))
	}))
	defer srv.Close()

	out, err := callOpenAICompat(srv.URL, "xai", "grok", "sys", "user")
	if err == nil {
		t.Fatal("expected error on non-200 status")
	}
	if out != "" {
		t.Errorf("expected empty output, got %q", out)
	}
	if !strings.Contains(err.Error(), "400") {
		t.Errorf("expected status code in error, got: %v", err)
	}
}
