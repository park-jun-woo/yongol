//ff:func feature=agent type=test control=sequence
//ff:what TestCallOpenAICompat — API 키 부재 조기반환 + httptest 로 성공/비200/파싱오류/빈선택/요청오류 분기 검증

package agent

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCallOpenAICompat(t *testing.T) {
	// No XAI_API_KEY and an empty XDG dir forces loadAPIKey to fail, so the
	// function returns before constructing or sending any HTTP request.
	t.Setenv("XAI_API_KEY", "")
	t.Setenv("XDG_CONFIG_HOME", t.TempDir())

	out, err := callOpenAICompat("https://api.x.ai/v1/chat/completions", "xai", "grok", "sys", "user")
	if err == nil {
		t.Fatal("expected error when API key is unavailable")
	}
	if out != "" {
		t.Errorf("expected empty output on error, got %q", out)
	}
}

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

func TestCallOpenAICompatParseError(t *testing.T) {
	t.Setenv("XAI_API_KEY", "fake-key")
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`not json`))
	}))
	defer srv.Close()

	_, err := callOpenAICompat(srv.URL, "xai", "grok", "sys", "user")
	if err == nil {
		t.Fatal("expected parse error on malformed JSON")
	}
	if !strings.Contains(err.Error(), "parse") {
		t.Errorf("expected parse error, got: %v", err)
	}
}

func TestCallOpenAICompatEmptyChoices(t *testing.T) {
	t.Setenv("XAI_API_KEY", "fake-key")
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"choices":[]}`))
	}))
	defer srv.Close()

	_, err := callOpenAICompat(srv.URL, "xai", "grok", "sys", "user")
	if err == nil {
		t.Fatal("expected error on empty choices")
	}
	if !strings.Contains(err.Error(), "empty choices") {
		t.Errorf("expected empty-choices error, got: %v", err)
	}
}

func TestCallOpenAICompatRequestError(t *testing.T) {
	t.Setenv("XAI_API_KEY", "fake-key")
	// A syntactically valid but unroutable scheme triggers http.DefaultClient.Do
	// to fail, exercising the "request: %w" branch.
	_, err := callOpenAICompat("http://127.0.0.1:0/x", "xai", "grok", "sys", "user")
	if err == nil {
		t.Fatal("expected request error for unreachable endpoint")
	}
}

func TestCallOpenAICompatNewRequestError(t *testing.T) {
	t.Setenv("XAI_API_KEY", "fake-key")
	// A control character in the URL makes http.NewRequest fail, exercising the
	// "create request: %w" branch before any network activity.
	_, err := callOpenAICompat("http://invalid\x7f.example/", "xai", "grok", "sys", "user")
	if err == nil {
		t.Fatal("expected create-request error for malformed URL")
	}
	if !strings.Contains(err.Error(), "create") {
		t.Errorf("expected create-request error, got: %v", err)
	}
}
