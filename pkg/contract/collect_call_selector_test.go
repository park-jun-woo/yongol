//ff:func feature=contract type=test control=iteration dimension=1
//ff:what test: TestCollectCallSelector — Queries 메서드/패키지 호출/denylist/비-exported 분기를 queries·calls 맵에 분류 검증

package contract

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

// bodyFromFunc parses a single function and returns its body block.
func bodyFromFunc(t *testing.T, body string) (*token.FileSet, *ast.BlockStmt) {
	t.Helper()
	src := "package p\nfunc F() {\n" + body + "\n}\n"
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "x.go", src, parser.ParseComments)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}
	fd := f.Decls[0].(*ast.FuncDecl)
	return fset, fd.Body
}

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

// TestCollectCallSelector_NonIdentReceiver covers the branch where the selector
// receiver is itself a SelectorExpr (not an *ast.Ident) and not a Queries
// receiver, so it is tagged but not routed into queries or calls.
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

// TestCollectCallSelector_NonCallNode exercises the early returns for nodes that
// are neither CallExpr nor have a SelectorExpr Fun.
func TestCollectCallSelector_NonCallNode(t *testing.T) {
	_, body := bodyFromFunc(t, "plainFunc()\n") // Fun is *ast.Ident, not SelectorExpr

	queries := map[string]struct{}{}
	calls := map[string]struct{}{}
	callSelectors := map[*ast.SelectorExpr]struct{}{}

	ast.Inspect(body, func(n ast.Node) bool {
		return collectCallSelector(n, queries, calls, callSelectors)
	})

	if len(callSelectors) != 0 {
		t.Errorf("expected no selectors tagged, got %v", callSelectors)
	}
}
