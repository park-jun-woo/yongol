//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestFindCallInStmt — 매칭 호출 발견 / 비매칭 nil / 첫 매칭 단락 검증
package qcheck

import (
	"go/ast"
	"testing"
)

func TestFindCallInStmt_Found(t *testing.T) {
	list := blockStmts(t, "_ = json.Unmarshal(b, v)")
	call := findCallInStmt(list[0], "json", "Unmarshal")
	if call == nil {
		t.Fatalf("expected to find json.Unmarshal call")
	}
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok || sel.Sel.Name != "Unmarshal" {
		t.Errorf("unexpected call: %+v", call)
	}
}
