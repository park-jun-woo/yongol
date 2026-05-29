//ff:func feature=validate type=test control=sequence topic=hurl-openapi
//ff:what TestXoh01_ExternalVarSkip — 외부 {{var}} 엔트리는 XOH-01 skip, 절대 URL 은 유지 (BUG-092)

package hurl_openapi

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestXoh01_ExternalVarSkip verifies BUG-092's XOH-01 change: an entry
// whose URL uses a non-host {{var}} prefix (URLVar="authurl", an external
// service) is skipped and raises no "path not declared" diagnostic, while
// an absolute http(s):// entry (URLVar="") to an unknown path STILL
// raises XOH-01 (protecting the existing behavior for the user's own API).
func TestXoh01_ExternalVarSkip(t *testing.T) {
	fs := &yongol.Fullstack{
		OpenAPIDoc: newDoc(map[string]map[string]*openapi3.Operation{
			"/workflows": {"GET": {OperationID: "ListWorkflows"}},
		}),
		HurlEntries: []hurl.HurlEntry{
			// External service ({{authurl}}) — must be skipped.
			{Method: "POST", Path: "/auth/v1/token", URLVar: "authurl", File: "ext.hurl", Line: 1},
			// Absolute URL to an unknown path — must still raise XOH-01.
			{Method: "GET", Path: "/unknown", URLVar: "", File: "own.hurl", Line: 2},
		},
	}

	diags := xoh01URLMethod(fs)
	if len(diags) != 1 {
		t.Fatalf("want exactly 1 diag (only the absolute-URL entry), got %d: %+v", len(diags), diags)
	}
	if diags[0].File != "own.hurl" {
		t.Fatalf("diag should belong to the absolute-URL entry, got file %q", diags[0].File)
	}
	if !strings.Contains(diags[0].Message, "[XOH-01]") || !strings.Contains(diags[0].Message, "path not declared") {
		t.Fatalf("unexpected msg: %q", diags[0].Message)
	}
}
