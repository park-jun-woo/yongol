//ff:func feature=validate type=test control=sequence topic=hurl-openapi
//ff:what TestXoh01_Positive_ExactMatch — path + method 일치 시 진단 없음

package hurl_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXoh01_Positive_ExactMatch(t *testing.T) {
	fs := &yongol.Fullstack{
		OpenAPIDoc: newDoc(map[string]map[string]*openapi3.Operation{
			"/gigs/{id}": {"GET": {OperationID: "GetGig"}},
		}),
		HurlEntries: []hurl.HurlEntry{
			{Method: "GET", Path: "/gigs/{{id}}", File: "t.hurl", Line: 1},
		},
	}
	if diags := xoh01URLMethod(fs); len(diags) != 0 {
		t.Fatalf("want 0 diags, got %d: %+v", len(diags), diags)
	}
}
