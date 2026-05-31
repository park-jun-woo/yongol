//ff:func feature=external type=test control=sequence
//ff:what external 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용

package external

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestExtractPathMethods_ZeroCov(t *testing.T) {
	doc := sampleDoc()
	pi := doc.Paths.Map()["/items/{item_id}"]
	methods := extractPathMethods(pi, "/items/{item_id}")
	if len(methods) != 1 || methods[0].HTTPMethod != "GET" {
		t.Errorf("expected one GET method, got %#v", methods)
	}
}

func TestReadHTTPSource_ZeroCov(t *testing.T) {
	body := `{"openapi":"3.0.0"}`
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(body))
	}))
	defer srv.Close()
	data, err := readHTTPSource(srv.URL)
	if err != nil || string(data) != body {
		t.Fatalf("happy path: data=%q err=%v", string(data), err)
	}

	// 404 with body -> error
	srv404 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("not found"))
	}))
	defer srv404.Close()
	if _, err := readHTTPSource(srv404.URL); err == nil {
		t.Errorf("expected error for 404")
	}
}
