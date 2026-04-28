//ff:func feature=validate type=test control=sequence topic=hurl-openapi
//ff:what TestXoh04_Negative_AssertPathMissing — response schema 에 없는 jsonpath → [XOH-04]

package hurl_openapi

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXoh04_Negative_AssertPathMissing(t *testing.T) {
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
				Asserts: []hurl.HurlAssert{{JSONPath: "$.nonexistent", Line: 5}},
			},
		},
	}
	diags := xoh04AssertPathInSchema(fs)
	if len(diags) != 1 || !strings.Contains(diags[0].Message, "[XOH-04]") {
		t.Fatalf("want 1 XOH-04 diag, got %+v", diags)
	}
}
