//ff:func feature=validate-contract type=test control=sequence
//ff:what TestCompareSignatureHasErr — HasErr 차이 diff 감지 검증

package contract

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/contract"
)

func TestCompareSignatureHasErr(t *testing.T) {
	a := contract.FuncSignature{Name: "X", HasErr: true, Returns: []string{"error"}}
	b := contract.FuncSignature{Name: "X", HasErr: false}
	diffs := compareSignature(a, b)
	if len(diffs) == 0 {
		t.Fatalf("expected diff")
	}
}
