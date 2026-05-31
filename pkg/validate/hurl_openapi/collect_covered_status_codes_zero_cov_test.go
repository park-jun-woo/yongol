//ff:func feature=validate type=test control=sequence topic=hurl-openapi
//ff:what zz_zerocov_test — hurl_openapi 0% 함수 (collectCoveredStatusCodes / xoh12CheckRoute / xoh13CheckFunc / xoh13CheckGuard / xoh13CheckHappyPath) 단위 테스트
package hurl_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
)

func TestCollectCoveredStatusCodes_ZeroCov(t *testing.T) {
	routes := []apiRoute{{
		Path:     "/users",
		Method:   "GET",
		Segments: []string{"users"},
		Op:       &openapi3.Operation{OperationID: "ListUsers"},
	}}
	entries := []hurl.HurlEntry{
		{Method: "GET", Path: "/users", StatusCode: "200"},
		{Method: "GET", Path: "/users", StatusCode: "401"},
		{Method: "GET", Path: "/users", StatusCode: ""},      // skipped: no code
		{Method: "GET", Path: "/unknown", StatusCode: "200"}, // skipped: no route
	}
	covered := collectCoveredStatusCodes(entries, routes)
	set := covered["ListUsers"]
	if !set["200"] || !set["401"] {
		t.Fatalf("expected 200 & 401 covered, got %v", set)
	}
	if len(set) != 2 {
		t.Errorf("expected exactly 2 codes, got %v", set)
	}
}
