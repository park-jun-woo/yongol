//ff:func feature=validate type=test control=sequence topic=ssac-func
//ff:what zz_zerocov_test — ssac_func 0% 규칙 (Run / xfs42 / xfs43 / xfs45 / xsf46) 단위 테스트
package ssac_func

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestXsf46CallResultIgnored_ZeroCov(t *testing.T) {
	// no @result but func has Response fields → warning.
	seq := parsessac.Sequence{
		Type:   "call",
		Model:  "billing.CheckCredits",
		Line:   6,
		Result: nil,
	}
	spec := funcspec.FuncSpec{
		Package:        "billing",
		Name:           "checkCredits",
		ResponseFields: []funcspec.Field{{Name: "Balance", Type: "int"}},
	}
	fs := buildFSZeroCov(seq, spec)
	if got := xsf46CallResultIgnored(fs); len(got) != 1 {
		t.Fatalf("expected 1 diag, got %d: %+v", len(got), got)
	}
}
