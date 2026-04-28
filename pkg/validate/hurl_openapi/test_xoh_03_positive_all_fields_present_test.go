//ff:func feature=validate type=test control=sequence topic=hurl-openapi
//ff:what TestXoh03_Positive_AllFieldsPresent — 모든 body 필드 존재 시 진단 없음

package hurl_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXoh03_Positive_AllFieldsPresent(t *testing.T) {
	op := withRequestBody("Register", map[string]*openapi3.Schema{
		"email":    {Type: &openapi3.Types{"string"}},
		"password": {Type: &openapi3.Types{"string"}},
	})
	fs := &yongol.Fullstack{
		OpenAPIDoc: newDoc(map[string]map[string]*openapi3.Operation{
			"/auth/register": {"POST": op},
		}),
		HurlEntries: []hurl.HurlEntry{
			{
				Method: "POST", Path: "/auth/register", StatusCode: "201",
				File: "t.hurl", Line: 1,
				BodyFields: []string{"email", "password"},
			},
		},
	}
	if diags := xoh03RequestFieldInSchema(fs); len(diags) != 0 {
		t.Fatalf("want 0 diags, got %d: %+v", len(diags), diags)
	}
}
