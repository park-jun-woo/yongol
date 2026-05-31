//ff:func feature=validate type=test control=iteration topic=sqlc
//ff:what TestXQS15InputKeyInitialism — @call 입력 키의 Go initialism 위반 검출

package ssac_sqlc

import (
	"strings"
	"testing"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXQS15InputKeyInitialism_ZeroCov(t *testing.T) {
	fs := &yongol.Fullstack{
		ServiceFuncs: []ssacparser.ServiceFunc{{
			Name:     "DoThing",
			FileName: "do_thing.ssac",
			Sequences: []ssacparser.Sequence{
				// non-call sequence → skipped
				{Type: "get", Model: "A.B", Args: []ssacparser.Arg{{Source: "request", Field: "OrgId"}}, Line: 2},
				// call with a violating key (OrgId → OrgID) appearing twice → dedup to one diag
				{Type: "call", Model: "svc.Do", Args: []ssacparser.Arg{{Source: "request", Field: "OrgId"}}, Line: 5},
				{Type: "call", Model: "svc.Do2", Args: []ssacparser.Arg{{Source: "request", Field: "OrgId"}}, Line: 8},
				// call with a clean key → no diag
				{Type: "call", Model: "svc.Ok", Args: []ssacparser.Arg{{Source: "request", Field: "OrgID"}}, Line: 11},
			},
		}},
	}

	diags := xqs15InputKeyInitialism(fs)
	if len(diags) != 1 {
		t.Fatalf("expected 1 deduped diag, got %d: %v", len(diags), diags)
	}
	d := diags[0]
	if d.File != "do_thing.ssac" || d.Line != 5 {
		t.Errorf("diag loc = %q:%d, want do_thing.ssac:5", d.File, d.Line)
	}
	if !strings.Contains(d.Message, "[XQS-15]") || !strings.Contains(d.Message, "OrgID") {
		t.Errorf("message missing parts: %q", d.Message)
	}
}

func TestXQS15InputKeyInitialism_NoFuncs_ZeroCov(t *testing.T) {
	if diags := xqs15InputKeyInitialism(&yongol.Fullstack{}); len(diags) != 0 {
		t.Errorf("expected 0 diags for empty fullstack, got %d", len(diags))
	}
}
