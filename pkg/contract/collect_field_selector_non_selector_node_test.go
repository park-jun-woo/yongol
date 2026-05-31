//ff:func feature=contract type=test control=sequence
//ff:what test: TestCollectFieldSelector — exported 필드 접근만 수집, 호출 selector·비-exported 제외 분기 검증
package contract

import (
	"go/ast"
	"testing"
)

func TestCollectFieldSelector_NonSelectorNode(t *testing.T) {
	fset, body := bodyFromFunc(t, "x := 1\n_ = x\n")
	callSelectors := map[*ast.SelectorExpr]struct{}{}
	fields := map[string]struct{}{}
	ast.Inspect(body, func(n ast.Node) bool {
		return collectFieldSelector(fset, n, callSelectors, fields)
	})
	if len(fields) != 0 {
		t.Errorf("expected no fields, got %v", fields)
	}
}
