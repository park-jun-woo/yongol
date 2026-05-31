//ff:func feature=agent type=test control=sequence
//ff:what TestDoOllamaRequest — 직렬화 불가 body 및 연결 불가 URL에서 에러 반환 검증
package agent

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestDoOllamaRequestNon200(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`boom`))
	}))
	defer srv.Close()

	_, err := doOllamaRequest(srv.URL, map[string]any{"model": "x"})
	if err == nil {
		t.Fatal("expected error on non-200 status")
	}
	if !strings.Contains(err.Error(), "500") {
		t.Errorf("expected status code in error, got: %v", err)
	}
}
