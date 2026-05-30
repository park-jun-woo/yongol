//ff:func feature=validate type=test topic=ssac-rego
//ff:what zz_zerocov_test — ssac_rego.Run 0% 커버리지 단위 테스트
package ssac_rego

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestRun_ZeroCov(t *testing.T) {
	if diags := Run(&yongol.Fullstack{}); len(diags) != 0 {
		t.Fatalf("empty fullstack → 0 diags, got %d: %+v", len(diags), diags)
	}
}

func TestXps28SSaCAuthToRego_ZeroCov(t *testing.T) {
	// SSaC @auth pair with no Rego policy → XPS-28 fires.
	fs := &yongol.Fullstack{
		ServiceFuncs: []ssac.ServiceFunc{{
			Name:     "DeleteProject",
			FileName: "svc/del.ssac",
			Sequences: []ssac.Sequence{{
				Type:     "auth",
				Action:   "delete",
				Resource: "project",
				Line:     3,
			}},
		}},
	}
	if diags := xps28SSaCAuthToRego(fs); len(diags) != 1 {
		t.Fatalf("expected 1 XPS-28 diag, got %d: %+v", len(diags), diags)
	}
	// nil fs → nil.
	if got := xps28SSaCAuthToRego(nil); got != nil {
		t.Errorf("nil fs should return nil, got %v", got)
	}
}

func TestXsp29RegoAllowToSSaC_ZeroCov(t *testing.T) {
	// Empty fs → no rego pairs → no diags, but body executes.
	if diags := xsp29RegoAllowToSSaC(&yongol.Fullstack{}); len(diags) != 0 {
		t.Errorf("empty fs should yield 0 diags, got %+v", diags)
	}
	// nil fs → nil.
	if got := xsp29RegoAllowToSSaC(nil); got != nil {
		t.Errorf("nil fs should return nil, got %v", got)
	}
}
