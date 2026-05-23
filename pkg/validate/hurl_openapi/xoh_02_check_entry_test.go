//ff:func feature=validate type=test control=iteration dimension=1 topic=hurl-openapi
//ff:what xoh02CheckEntry — hurl entry의 status code가 OpenAPI에 선언됐는지 검사 검증

package hurl_openapi

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
)

func TestXoh02CheckEntry(t *testing.T) {
	routes := []apiRoute{
		{
			Path: "/users", Method: "GET", Segments: []string{"users"},
			Responses: map[string]bool{"200": true, "404": true},
		},
	}

	cases := []struct {
		name     string
		entry    hurl.HurlEntry
		wantDiag bool
	}{
		{
			name:     "empty_status_skipped",
			entry:    hurl.HurlEntry{Method: "GET", Path: "/users"},
			wantDiag: false,
		},
		{
			name:     "no_route_match_skipped",
			entry:    hurl.HurlEntry{Method: "GET", Path: "/orders", StatusCode: "200"},
			wantDiag: false,
		},
		{
			name:     "declared_status_no_diag",
			entry:    hurl.HurlEntry{Method: "GET", Path: "/users", StatusCode: "200"},
			wantDiag: false,
		},
		{
			name:     "undeclared_status_produces_diag",
			entry:    hurl.HurlEntry{Method: "GET", Path: "/users", StatusCode: "500", File: "t.hurl", Line: 5},
			wantDiag: true,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			diag, ok := xoh02CheckEntry(c.entry, routes)
			if ok != c.wantDiag {
				t.Fatalf("ok = %v, want %v; diag=%v", ok, c.wantDiag, diag)
			}
			if ok && !strings.Contains(diag.Message, "[XOH-02]") {
				t.Errorf("expected [XOH-02] in message, got %q", diag.Message)
			}
		})
	}
}
