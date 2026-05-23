//ff:func feature=validate type=test control=iteration dimension=1 topic=hurl-openapi
//ff:what xoh02StatusDeclared — hurl HTTP status가 OpenAPI responses에 선언됐는지 검증

package hurl_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXoh02StatusDeclared(t *testing.T) {
	doc := &openapi3.T{Paths: &openapi3.Paths{}}
	op := &openapi3.Operation{Responses: &openapi3.Responses{}}
	op.Responses.Set("200", &openapi3.ResponseRef{Value: openapi3.NewResponse()})
	doc.Paths.Set("/users", &openapi3.PathItem{Get: op})

	cases := []TestXoh02StatusDeclaredCase{
		{name: "nil_fs", fs: nil, wantCount: 0},
		{name: "nil_doc", fs: &yongol.Fullstack{}, wantCount: 0},
		{
			name: "declared_status_no_diag",
			fs: &yongol.Fullstack{
				OpenAPIDoc:  doc,
				HurlEntries: []hurl.HurlEntry{{Method: "GET", Path: "/users", StatusCode: "200"}},
			},
			wantCount: 0,
		},
		{
			name: "undeclared_status_produces_diag",
			fs: &yongol.Fullstack{
				OpenAPIDoc:  doc,
				HurlEntries: []hurl.HurlEntry{{Method: "GET", Path: "/users", StatusCode: "404", File: "t.hurl", Line: 1}},
			},
			wantCount: 1,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runXoh02StatusDeclared(t, c)
		})
	}
}
