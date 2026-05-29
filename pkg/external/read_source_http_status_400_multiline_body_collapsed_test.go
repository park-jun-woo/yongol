//ff:func feature=external type=test control=sequence
//ff:what readSource — multiline body 의 newline 은 공백으로 치환

package external

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

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
