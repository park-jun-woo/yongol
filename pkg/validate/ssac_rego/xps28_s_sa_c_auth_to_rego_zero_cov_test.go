//ff:func feature=validate type=test control=sequence topic=ssac-rego
//ff:what zz_zerocov_test — ssac_rego.Run 0% 커버리지 단위 테스트
package ssac_rego

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

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
