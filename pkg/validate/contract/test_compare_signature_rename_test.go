//ff:func feature=validate-contract type=test control=sequence
//ff:what TestCompareSignatureRename — 함수명 불일치 diff 감지 검증

package contract

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/contract"
)

func TestCompareSignatureRename(t *testing.T) {
	a := contract.FuncSignature{Name: "Foo"}
	b := contract.FuncSignature{Name: "Bar"}
	diffs := compareSignature(a, b)
	if len(diffs) == 0 || !strings.Contains(diffs[0], "function name") {
		t.Errorf("expected function-name diff, got %v", diffs)
	}
}
