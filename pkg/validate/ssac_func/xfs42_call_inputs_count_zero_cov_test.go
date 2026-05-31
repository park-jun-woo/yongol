//ff:func feature=validate type=test control=sequence topic=ssac-func
//ff:what zz_zerocov_test — ssac_func 0% 규칙 (Run / xfs42 / xfs43 / xfs45 / xsf46) 단위 테스트
package ssac_func

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXfs42CallInputsCount_ZeroCov(t *testing.T) {
	// 2 inputs vs 1 request field → fires.
	seq := parsessac.Sequence{
		Type:   "call",
		Model:  "billing.CheckCredits",
		Line:   3,
		Inputs: map[string]string{"OrgID": "x", "Extra": "y"},
	}
	spec := funcspec.FuncSpec{
		Package:       "billing",
		Name:          "checkCredits",
		RequestFields: []funcspec.Field{{Name: "OrgID", Type: "string"}},
	}
	fs := buildFSZeroCov(seq, spec)
	if got := xfs42CallInputsCount(fs); len(got) != 1 {
		t.Fatalf("expected 1 diag, got %d: %+v", len(got), got)
	}
	// nil Ground → nil.
	if got := xfs42CallInputsCount(&yongol.Fullstack{}); got != nil {
		t.Errorf("nil ground should return nil, got %v", got)
	}
}
