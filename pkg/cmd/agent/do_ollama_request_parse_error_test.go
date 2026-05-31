//ff:func feature=agent type=test control=sequence
//ff:what TestDoOllamaRequest — 직렬화 불가 body 및 연결 불가 URL에서 에러 반환 검증
package agent

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestDoOllamaRequestParseError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`not json`))
	}))
	defer srv.Close()

	_, err := doOllamaRequest(srv.URL, map[string]any{"model": "x"})
	if err == nil {
		t.Fatal("expected parse error on malformed JSON")
	}
	if !strings.Contains(err.Error(), "parse ollama response") {
		t.Errorf("expected parse error, got: %v", err)
	}
}
