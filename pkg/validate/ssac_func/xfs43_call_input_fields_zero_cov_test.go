//ff:func feature=validate type=test control=sequence topic=ssac-func
//ff:what zz_zerocov_test — ssac_func 0% 규칙 (Run / xfs42 / xfs43 / xfs45 / xsf46) 단위 테스트
package ssac_func

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXfs43CallInputFields_ZeroCov(t *testing.T) {
	// input key not in request fields → fires.
	seq := parsessac.Sequence{
		Type:   "call",
		Model:  "billing.CheckCredits",
		Line:   4,
		Inputs: map[string]string{"Unknown": "x"},
	}
	spec := funcspec.FuncSpec{
		Package:       "billing",
		Name:          "checkCredits",
		RequestFields: []funcspec.Field{{Name: "OrgID", Type: "string"}},
	}
	fs := buildFSZeroCov(seq, spec)
	if got := xfs43CallInputFields(fs); len(got) != 1 {
		t.Fatalf("expected 1 diag, got %d: %+v", len(got), got)
	}
	if got := xfs43CallInputFields(&yongol.Fullstack{}); got != nil {
		t.Errorf("nil ground should return nil, got %v", got)
	}
}
