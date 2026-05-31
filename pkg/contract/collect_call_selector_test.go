//ff:func feature=contract type=test control=sequence
//ff:what test: TestCollectCallSelector — Queries 메서드/패키지 호출/denylist/비-exported 분기를 queries·calls 맵에 분류 검증
package contract

import (
	"go/ast"
	"testing"
)

func TestCollectCallSelector(t *testing.T) {
	_, body := bodyFromFunc(t,
		"server.Queries.GetUserByID(ctx, 1)\n"+ // queries
			"billing.HoldEscrow(ctx, 100)\n"+ // exported pkg call
			"ctx.Done()\n"+ // denylisted receiver -> skipped
			"helper.privateFn()\n") // non-exported selector -> skipped

	queries := map[string]struct{}{}
	calls := map[string]struct{}{}
	callSelectors := map[*ast.SelectorExpr]struct{}{}

	ast.Inspect(body, func(n ast.Node) bool {
		return collectCallSelector(n, queries, calls, callSelectors)
	})

	if _, ok := queries["GetUserByID"]; !ok {
		t.Errorf("expected GetUserByID in queries, got %v", queries)
	}
	if len(queries) != 1 {
		t.Errorf("queries: got %v want exactly {GetUserByID}", queries)
	}
	if _, ok := calls["billing.HoldEscrow"]; !ok {
		t.Errorf("expected billing.HoldEscrow in calls, got %v", calls)
	}
	if len(calls) != 1 {
		t.Errorf("calls: got %v want exactly {billing.HoldEscrow}", calls)
	}
	if len(callSelectors) == 0 {
		t.Errorf("expected callSelectors to be populated")
	}
}
