//ff:func feature=external type=test control=sequence
//ff:what readSource — HTTP 404 + body 는 "status 404" + body snippet 포함 에러

package external

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

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
