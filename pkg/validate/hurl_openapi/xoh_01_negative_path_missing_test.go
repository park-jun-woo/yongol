//ff:func feature=validate type=test control=sequence topic=hurl-openapi
//ff:what TestXoh01_Negative_PathMissing — OpenAPI 에 path 선언 없음 → [XOH-01]

package hurl_openapi

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXoh01_Negative_PathMissing(t *testing.T) {
	fs := &yongol.Fullstack{
		OpenAPIDoc: newDoc(map[string]map[string]*openapi3.Operation{
			"/auth/login": {"POST": {OperationID: "Login"}},
		}),
		HurlEntries: []hurl.HurlEntry{
			{Method: "GET", Path: "/gigs", File: "t.hurl", Line: 1},
		},
	}
	diags := xoh01URLMethod(fs)
	if len(diags) != 1 {
		t.Fatalf("want 1 diag, got %d", len(diags))
	}
	if !strings.Contains(diags[0].Message, "[XOH-01]") || !strings.Contains(diags[0].Message, "path not declared") {
		t.Fatalf("unexpected msg: %q", diags[0].Message)
	}
}
