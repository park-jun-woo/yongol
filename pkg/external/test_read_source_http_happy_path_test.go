//ff:func feature=external type=test control=sequence
//ff:what readSource — HTTP 200 정상 응답은 body bytes 반환

package external

import (
	"net/http"
	"net/http/httptest"
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
