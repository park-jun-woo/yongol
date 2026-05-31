//ff:func feature=contract type=test control=sequence
//ff:what test: TestCollectFieldSelector — exported 필드 접근만 수집, 호출 selector·비-exported 제외 분기 검증
package contract

import (
	"go/ast"
	"testing"
)

func TestCollectFieldSelector_NestedReceiver(t *testing.T) {
	fset, body := bodyFromFunc(t, "_ = outer.Inner.Field\n")

	callSelectors := map[*ast.SelectorExpr]struct{}{}
	fields := map[string]struct{}{}
	ast.Inspect(body, func(n ast.Node) bool {
		return collectFieldSelector(fset, n, callSelectors, fields)
	})

	// outer.Inner (Inner exported) and outer.Inner.Field (Field exported) both
	// collected; the latter uses the printer path for its receiver.
	if _, ok := fields["outer.Inner.Field"]; !ok {
		t.Errorf("expected outer.Inner.Field, got %v", fields)
	}
}
