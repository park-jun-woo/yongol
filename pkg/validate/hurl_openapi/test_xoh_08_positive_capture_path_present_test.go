//ff:func feature=validate type=test control=sequence topic=hurl-openapi
//ff:what TestXoh08_Positive_CapturePathPresent — schema 에 존재하는 capture jsonpath 는 진단 없음

package hurl_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXoh08_Positive_CapturePathPresent(t *testing.T) {
	fs := &yongol.Fullstack{
		OpenAPIDoc: newDoc(map[string]map[string]*openapi3.Operation{
			"/auth/login": {"POST": withJSONResponse("Login", "200", map[string]*openapi3.Schema{
				"access_token": {Type: &openapi3.Types{"string"}},
			})},
		}),
		HurlEntries: []hurl.HurlEntry{
			{
				Method: "POST", Path: "/auth/login", StatusCode: "200",
				File: "t.hurl", Line: 1,
				Captures: []hurl.HurlCapture{{
					Var:      "token",
					Source:   "jsonpath",
					JSONPath: "$.access_token",
					Line:     10,
				}},
			},
		},
	}
	if diags := xoh08CapturePathInSchema(fs); len(diags) != 0 {
		t.Fatalf("want 0 diags, got %+v", diags)
	}
}
