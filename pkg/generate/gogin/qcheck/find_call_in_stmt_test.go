//ff:func feature=gen-gogin type=test control=branch topic=err-guard
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

func TestFindCallInStmt_NotFound(t *testing.T) {
	list := blockStmts(t, "_ = json.Marshal(v)")
	if got := findCallInStmt(list[0], "json", "Unmarshal"); got != nil {
		t.Errorf("expected nil for non-matching call, got %+v", got)
	}
}

func TestFindCallInStmt_FirstMatchOnly(t *testing.T) {
	// Two matching calls; first one short-circuits the inspection.
	list := blockStmts(t, "_, _ = json.Unmarshal(a, b), json.Unmarshal(c, d)")
	call := findCallInStmt(list[0], "json", "Unmarshal")
	if call == nil {
		t.Fatalf("expected a match")
	}
}

func TestFindCallInStmt_NoCall(t *testing.T) {
	list := blockStmts(t, "x := 1")
	if got := findCallInStmt(list[0], "json", "Unmarshal"); got != nil {
		t.Errorf("expected nil when no call present, got %+v", got)
	}
}
