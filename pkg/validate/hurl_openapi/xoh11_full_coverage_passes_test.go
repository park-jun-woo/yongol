//ff:func feature=validate type=test control=sequence topic=hurl-openapi
//ff:what TestXoh11_FullCoveragePasses — 전수 커버 → 진단 없음

package hurl_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXoh11_FullCoveragePasses(t *testing.T) {
	fs := &yongol.Fullstack{
		OpenAPIDoc: newDoc(map[string]map[string]*openapi3.Operation{
			"/gigs":       {"GET": {OperationID: "ListGigs"}, "POST": {OperationID: "CreateGig"}},
			"/gigs/{id}":  {"GET": {OperationID: "GetGig"}},
		}),
		HurlFiles: []string{"specs/tests/smoke.hurl"},
		HurlEntries: []hurl.HurlEntry{
			{Method: "GET", Path: "/gigs", File: "specs/tests/smoke.hurl", Line: 1},
			{Method: "POST", Path: "/gigs", File: "specs/tests/smoke.hurl", Line: 5},
			{Method: "GET", Path: "/gigs/{{id}}", File: "specs/tests/smoke.hurl", Line: 9},
		},
	}
	if diags := xoh11SmokeCoverage(fs); len(diags) != 0 {
		t.Fatalf("want 0 diags, got %d: %+v", len(diags), diags)
	}
}
