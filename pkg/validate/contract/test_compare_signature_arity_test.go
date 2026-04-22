//ff:func feature=validate-contract type=test control=iteration dimension=1
//ff:what TestCompareSignatureArity — 파라미터 개수 차이 diff 감지 검증

package contract

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/contract"
)

func TestCompareSignatureArity(t *testing.T) {
	a := contract.FuncSignature{Name: "X", Params: []contract.FuncParam{{Type: "int"}}}
	b := contract.FuncSignature{Name: "X"}
	diffs := compareSignature(a, b)
	if len(diffs) == 0 {
		t.Fatalf("expected arity diff")
	}
	found := false
	for _, d := range diffs {
		if strings.Contains(d, "parameter arity") {
			found = true
		}
	}
	if !found {
		t.Errorf("expected parameter arity in diffs, got %v", diffs)
	}
}
