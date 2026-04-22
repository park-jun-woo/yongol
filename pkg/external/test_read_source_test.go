//ff:func feature=external type=loader control=sequence
//ff:what TestReadSource: readSource 의 happy-path / 에러 경로(HTTP status, body 스니펫, 줄바꿈 붕괴, 대용량 자름, 네트워크 실패, 파일 경로, 파일 미존재)
package external

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestReadSource_HTTPHappyPath(t *testing.T) {
	body := `{"openapi":"3.0.0"}`
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(body))
	}))
	defer srv.Close()

	data, err := readSource(srv.URL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if string(data) != body {
		t.Errorf("body mismatch: got %q want %q", string(data), body)
	}
}

func TestReadSource_HTTPStatus404WithBody(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"error":"not found"}`))
	}))
	defer srv.Close()

	_, err := readSource(srv.URL)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	msg := err.Error()
	if !strings.Contains(msg, "status 404") {
		t.Errorf("expected 'status 404' in error, got: %s", msg)
	}
	if !strings.Contains(msg, "not found") {
		t.Errorf("expected 'not found' in error, got: %s", msg)
	}
	if !strings.Contains(msg, srv.URL) {
		t.Errorf("expected URL %q in error, got: %s", srv.URL, msg)
	}
}

func TestReadSource_HTTPStatus500NoBody(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer srv.Close()

	_, err := readSource(srv.URL)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	msg := err.Error()
	if !strings.Contains(msg, "status 500") {
		t.Errorf("expected 'status 500' in error, got: %s", msg)
	}
	// Expected format: "fetch URL <url>: status 500" — no trailing ": <snippet>"
	// URL itself may contain a colon (e.g., "http://127.0.0.1:PORT").
	// After trimming "fetch URL " and the URL, the remainder must be exactly ": status 500".
	prefix := "fetch URL " + srv.URL
	if !strings.HasPrefix(msg, prefix) {
		t.Fatalf("expected prefix %q, got: %s", prefix, msg)
	}
	tail := strings.TrimPrefix(msg, prefix)
	if tail != ": status 500" {
		t.Errorf("expected tail %q (no snippet section), got: %q", ": status 500", tail)
	}
}

func TestReadSource_HTTPStatus400LargeBodyTruncated(t *testing.T) {
	// 2 KiB body should be trimmed to 1 KiB + "..."
	bigBody := strings.Repeat("a", 2048)
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(bigBody))
	}))
	defer srv.Close()

	_, err := readSource(srv.URL)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	msg := err.Error()
	if !strings.Contains(msg, "status 400") {
		t.Errorf("expected 'status 400' in error, got: %s", msg)
	}
	if !strings.HasSuffix(msg, "...") {
		t.Errorf("expected message to end with '...' for truncated body, got: %s", msg)
	}
	// Ensure body was indeed limited: snippet of 1024 'a's + "..." appears.
	expectedSnippet := strings.Repeat("a", 1024) + "..."
	if !strings.Contains(msg, expectedSnippet) {
		t.Errorf("expected truncated snippet of 1024 'a's + '...' in error, got length=%d", len(msg))
	}
	// And the full 2048-char sequence must NOT be present.
	if strings.Contains(msg, strings.Repeat("a", 2048)) {
		t.Errorf("error message should not contain full 2048-byte body")
	}
}

func TestReadSource_HTTPStatus400MultilineBodyCollapsed(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("line1\nline2"))
	}))
	defer srv.Close()

	_, err := readSource(srv.URL)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	msg := err.Error()
	if strings.Contains(msg, "\n") {
		t.Errorf("expected no newline in error message, got: %q", msg)
	}
	if !strings.Contains(msg, "line1 line2") {
		t.Errorf("expected 'line1 line2' (newline collapsed to space) in error, got: %s", msg)
	}
}

func TestReadSource_HTTPNetworkError(t *testing.T) {
	// Start and immediately close — URL is still valid string but server is gone.
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	url := srv.URL
	srv.Close()

	_, err := readSource(url)
	if err == nil {
		t.Fatal("expected network error, got nil")
	}
	msg := err.Error()
	if !strings.Contains(msg, "fetch URL") {
		t.Errorf("expected 'fetch URL' wrap in error chain, got: %s", msg)
	}
	if !strings.Contains(msg, url) {
		t.Errorf("expected URL %q in error, got: %s", url, msg)
	}
}

func TestReadSource_FilePath(t *testing.T) {
	tmp := t.TempDir()
	path := filepath.Join(tmp, "spec.yaml")
	content := []byte("openapi: 3.0.0\n")
	if err := os.WriteFile(path, content, 0644); err != nil {
		t.Fatalf("setup write: %v", err)
	}

	data, err := readSource(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if string(data) != string(content) {
		t.Errorf("content mismatch: got %q want %q", string(data), string(content))
	}
}

func TestReadSource_FileNotFound(t *testing.T) {
	path := filepath.Join(t.TempDir(), "does-not-exist.yaml")
	_, err := readSource(path)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !os.IsNotExist(err) {
		t.Errorf("expected os.IsNotExist(err) = true, got: %v", err)
	}
}
