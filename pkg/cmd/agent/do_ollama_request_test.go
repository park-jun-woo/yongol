//ff:func feature=agent type=test control=sequence
//ff:what TestDoOllamaRequest — 직렬화 불가 body 및 연결 불가 URL에서 에러 반환 검증

package agent

import (
	"net/http"
	"net/http/httptest"
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
