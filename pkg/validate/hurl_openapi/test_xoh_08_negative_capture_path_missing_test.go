//ff:func feature=validate type=test control=sequence topic=hurl-openapi
//ff:what TestXoh08_Negative_CapturePathMissing — schema 에 없는 capture jsonpath → [XOH-08]

package hurl_openapi

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXoh08_Negative_CapturePathMissing(t *testing.T) {
	fs := &yongol.Fullstack{
		OpenAPIDoc: newDoc(map[string]map[string]*openapi3.Operation{
			"/auth/register": {"POST": withJSONResponse("Register", "201", map[string]*openapi3.Schema{
				"user": {Type: &openapi3.Types{"object"}},
			})},
		}),
		HurlEntries: []hurl.HurlEntry{
			{
				Method: "POST", Path: "/auth/register", StatusCode: "201",
				File: "t.hurl", Line: 1,
				Captures: []hurl.HurlCapture{{
					Var:      "token",
					Source:   "jsonpath",
					JSONPath: "$.access_token", // BUG-031 — not in response
					Line:     10,
				}},
			},
		},
	}
	diags := xoh08CapturePathInSchema(fs)
	if len(diags) != 1 || !strings.Contains(diags[0].Message, "[XOH-08]") {
		t.Fatalf("want 1 XOH-08 diag, got %+v", diags)
	}
}
