//ff:func feature=validate type=test control=iteration dimension=1 topic=hurl-openapi
//ff:what xoh03RequestFieldInSchema — hurl request body JSON field가 OpenAPI schema에 존재하는지 검증

package hurl_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXoh03RequestFieldInSchema(t *testing.T) {
	doc := &openapi3.T{Paths: &openapi3.Paths{}}
	op := &openapi3.Operation{
		Responses: &openapi3.Responses{},
		RequestBody: &openapi3.RequestBodyRef{
			Value: openapi3.NewRequestBody().WithJSONSchema(&openapi3.Schema{
				Properties: openapi3.Schemas{
					"name": &openapi3.SchemaRef{Value: &openapi3.Schema{}},
				},
			}),
		},
	}
	doc.Paths.Set("/users", &openapi3.PathItem{Post: op})

	cases := []struct {
		name      string
		fs        *yongol.Fullstack
		wantCount int
	}{
		{name: "nil_fs", fs: nil, wantCount: 0},
		{name: "nil_doc", fs: &yongol.Fullstack{}, wantCount: 0},
		{
			name: "valid_field_no_diag",
			fs: &yongol.Fullstack{
				OpenAPIDoc:  doc,
				HurlEntries: []hurl.HurlEntry{{Method: "POST", Path: "/users", BodyFields: []string{"name"}}},
			},
			wantCount: 0,
		},
		{
			name: "invalid_field_produces_diag",
			fs: &yongol.Fullstack{
				OpenAPIDoc:  doc,
				HurlEntries: []hurl.HurlEntry{{Method: "POST", Path: "/users", BodyFields: []string{"phone"}, File: "t.hurl"}},
			},
			wantCount: 1,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runDiagCodeCase(t, xoh03RequestFieldInSchema(c.fs), c.wantCount, "[XOH-03]")
		})
	}
}
