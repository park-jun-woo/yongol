//ff:func feature=validate type=test control=sequence topic=hurl-openapi
//ff:what TestXoh11_MissingOpsRaises — 미커버 operationId → [XOH-11] ERROR + missing 목록

package hurl_openapi

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXoh11_MissingOpsRaises(t *testing.T) {
	fs := &yongol.Fullstack{
		OpenAPIDoc: newDoc(map[string]map[string]*openapi3.Operation{
			"/gigs":       {"GET": {OperationID: "ListGigs"}, "POST": {OperationID: "CreateGig"}},
			"/gigs/{id}":  {"GET": {OperationID: "GetGig"}},
			"/auth/login": {"POST": {OperationID: "Login"}},
		}),
		HurlFiles: []string{"specs/tests/smoke.hurl"},
		HurlEntries: []hurl.HurlEntry{
			{Method: "GET", Path: "/gigs", File: "specs/tests/smoke.hurl", Line: 1},
			{Method: "POST", Path: "/auth/login", File: "specs/tests/smoke.hurl", Line: 5},
		},
	}
	diags := xoh11SmokeCoverage(fs)
	if len(diags) != 1 {
		t.Fatalf("want 1 diag, got %d", len(diags))
	}
	msg := diags[0].Message
	if !strings.Contains(msg, "[XOH-11]") {
		t.Fatalf("want [XOH-11], got: %q", msg)
	}
	if !strings.Contains(msg, "2/4") {
		t.Fatalf("want 2/4, got: %q", msg)
	}
	if !strings.Contains(msg, "CreateGig") {
		t.Fatalf("want CreateGig in missing, got: %q", msg)
	}
	if !strings.Contains(msg, "GetGig") {
		t.Fatalf("want GetGig in missing, got: %q", msg)
	}
}
