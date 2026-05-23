//ff:func feature=validate type=test control=iteration dimension=1 topic=hurl-openapi
//ff:what xoh01URLMethod — hurl 요청 URL path + method가 OpenAPI에 선언되었는지 검증

package hurl_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXoh01URLMethod(t *testing.T) {
	doc := &openapi3.T{Paths: &openapi3.Paths{}}
	doc.Paths.Set("/users", &openapi3.PathItem{
		Get:  &openapi3.Operation{Responses: &openapi3.Responses{}},
		Post: &openapi3.Operation{Responses: &openapi3.Responses{}},
	})

	cases := []TestXoh01URLMethodCase{
		{name: "nil_fs", fs: nil, wantCount: 0},
		{name: "nil_doc", fs: &yongol.Fullstack{}, wantCount: 0},
		{
			name: "matching_route_no_diag",
			fs: &yongol.Fullstack{
				OpenAPIDoc: doc,
				HurlEntries: []hurl.HurlEntry{
					{Method: "GET", Path: "/users", File: "t.hurl", Line: 1},
				},
			},
			wantCount: 0,
		},
		{
			name: "unmatched_path_produces_diag",
			fs: &yongol.Fullstack{
				OpenAPIDoc: doc,
				HurlEntries: []hurl.HurlEntry{
					{Method: "GET", Path: "/orders", File: "t.hurl", Line: 1},
				},
			},
			wantCount: 1,
		},
		{
			name: "unmatched_method_produces_diag",
			fs: &yongol.Fullstack{
				OpenAPIDoc: doc,
				HurlEntries: []hurl.HurlEntry{
					{Method: "DELETE", Path: "/users", File: "t.hurl", Line: 1},
				},
			},
			wantCount: 1,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runXoh01URLMethod(t, c)
		})
	}
}
