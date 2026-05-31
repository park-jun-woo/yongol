//ff:func feature=contract type=test control=sequence
//ff:what test: TestCollectCallSelector — Queries 메서드/패키지 호출/denylist/비-exported 분기를 queries·calls 맵에 분류 검증
package contract

import (
	"go/ast"
	"testing"
)

func TestCollectCallSelector_NonIdentReceiver(t *testing.T) {
	_, body := bodyFromFunc(t, "outer.inner.Method(ctx)\n")

	queries := map[string]struct{}{}
	calls := map[string]struct{}{}
	callSelectors := map[*ast.SelectorExpr]struct{}{}

	ast.Inspect(body, func(n ast.Node) bool {
		return collectCallSelector(n, queries, calls, callSelectors)
	})

	if len(queries) != 0 {
		t.Errorf("expected no queries, got %v", queries)
	}
	if len(calls) != 0 {
		t.Errorf("expected no calls, got %v", calls)
	}
	if len(callSelectors) == 0 {
		t.Errorf("expected the outer selector to be tagged")
	}
}
