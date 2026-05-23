//ff:func feature=validate type=test control=iteration dimension=1 topic=hurl-openapi
//ff:what collectCoveredOps — smoke.hurl entries에서 커버된 operationId 수집 검증

package hurl_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
)

func TestCollectCoveredOps(t *testing.T) {
	routes := []apiRoute{
		{Path: "/users", Method: "GET", Segments: []string{"users"},
			Op: &openapi3.Operation{OperationID: "getUsers"}},
		{Path: "/users", Method: "POST", Segments: []string{"users"},
			Op: &openapi3.Operation{OperationID: "createUser"}},
	}

	cases := []struct {
		name    string
		entries []hurl.HurlEntry
		want    map[string]bool
	}{
		{name: "nil_entries", entries: nil, want: map[string]bool{}},
		{
			name: "matching_entry",
			entries: []hurl.HurlEntry{
				{Method: "GET", Path: "/users"},
			},
			want: map[string]bool{"getUsers": true},
		},
		{
			name: "non_matching_entry",
			entries: []hurl.HurlEntry{
				{Method: "GET", Path: "/orders"},
			},
			want: map[string]bool{},
		},
		{
			name: "multiple_entries",
			entries: []hurl.HurlEntry{
				{Method: "GET", Path: "/users"},
				{Method: "POST", Path: "/users"},
			},
			want: map[string]bool{"getUsers": true, "createUser": true},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runBoolMapCase(t, collectCoveredOps(c.entries, routes), c.want)
		})
	}
}
