//ff:func feature=external type=test control=sequence
//ff:what readSource — 2KiB body 는 1KiB + "..." 로 truncate 된 snippet 반환

package external

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

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
