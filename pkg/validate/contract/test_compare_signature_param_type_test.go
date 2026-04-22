//ff:func feature=validate-contract type=test control=sequence
//ff:what TestCompareSignatureParamType — 파라미터 타입 차이 diff 감지 검증

package contract

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/contract"
)

func TestCompareSignatureParamType(t *testing.T) {
	a := contract.FuncSignature{Name: "X", Params: []contract.FuncParam{{Type: "int"}}}
	b := contract.FuncSignature{Name: "X", Params: []contract.FuncParam{{Type: "string"}}}
	diffs := compareSignature(a, b)
	if len(diffs) != 1 || !strings.Contains(diffs[0], "param #1 type") {
		t.Errorf("expected param type diff, got %v", diffs)
	}
}
