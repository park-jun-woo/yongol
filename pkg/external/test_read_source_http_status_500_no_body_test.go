//ff:func feature=external type=test control=sequence
//ff:what readSource — HTTP 500 + 빈 body 는 snippet 섹션 없이 "status 500" 만 에러

package external

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

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
