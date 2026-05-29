//ff:func feature=validate-contract type=test control=sequence
//ff:what TestCompareSignatureMatch — 동일 시그니처 비교 시 diff 없음 검증

package contract

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/contract"
)

func TestCompareSignatureMatch(t *testing.T) {
	sig := contract.FuncSignature{
		Name:    "X",
		Params:  []contract.FuncParam{{Name: "a", Type: "int"}},
		Returns: []string{"error"},
		HasErr:  true,
	}
	if diffs := compareSignature(sig, sig); len(diffs) != 0 {
		t.Errorf("identical sigs should match, got diffs: %v", diffs)
	}
}
