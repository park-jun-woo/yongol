//ff:func feature=contract type=test control=iteration dimension=1
//ff:what test: TestCollectFieldSelector — exported 필드 접근만 수집, 호출 selector·비-exported 제외 분기 검증

package contract

import (
	"go/ast"
	"testing"
)

func TestCollectFieldSelector(t *testing.T) {
	fset, body := bodyFromFunc(t,
		"_ = user.Name\n"+ // exported field -> collected
			"_ = user.age\n"+ // non-exported -> skipped
			"obj.Method()\n") // call selector -> skipped (recorded in callSelectors)

	// First pass records call selectors so the field pass can skip them.
	callSelectors := map[*ast.SelectorExpr]struct{}{}
	queries := map[string]struct{}{}
	calls := map[string]struct{}{}
	ast.Inspect(body, func(n ast.Node) bool {
		return collectCallSelector(n, queries, calls, callSelectors)
	})

	fields := map[string]struct{}{}
	ast.Inspect(body, func(n ast.Node) bool {
		return collectFieldSelector(fset, n, callSelectors, fields)
	})

	if _, ok := fields["user.Name"]; !ok {
		t.Errorf("expected user.Name in fields, got %v", fields)
	}
	if _, ok := fields["user.age"]; ok {
		t.Errorf("non-exported user.age should be skipped, got %v", fields)
	}
	if _, ok := fields["obj.Method"]; ok {
		t.Errorf("call selector obj.Method should be skipped, got %v", fields)
	}
	if len(fields) != 1 {
		t.Errorf("fields: got %v want exactly {user.Name}", fields)
	}
}

// TestCollectFieldSelector_NestedReceiver covers a non-Ident receiver routed
// through go/printer (renderRecv printer path) producing a chained field key.
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

// TestCollectFieldSelector_NonSelectorNode covers the early return for nodes
// that are not SelectorExpr.
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
