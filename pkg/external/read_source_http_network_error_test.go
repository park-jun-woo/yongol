//ff:func feature=external type=test control=sequence
//ff:what readSource — 서버 다운 시 "fetch URL" 래핑된 네트워크 에러 반환

package external

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

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
