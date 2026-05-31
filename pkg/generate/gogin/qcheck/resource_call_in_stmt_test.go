//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestResourceCallInStmt — 미닫힘 리소스 DF-06 + 닫힘/비-assign/discard 등 스킵 분기 검증
package qcheck

import (
	"go/token"
	"testing"
)

func TestResourceCallInStmt(t *testing.T) {
	fset := token.NewFileSet()

	t.Run("UnclosedFinding", func(t *testing.T) {
		list := blockStmts(t, "f, err := os.Open(p)\n_ = err\n_ = f")
		got := resourceCallInStmt(list[0], list, 0, fset)
		if len(got) != 1 || got[0].Category != "DF-06" || got[0].Detail != "os.Open" {
			t.Fatalf("expected one DF-06 os.Open finding, got %+v", got)
		}
	})

	t.Run("ClosedNoFinding", func(t *testing.T) {
		list := blockStmts(t, "f, err := os.Open(p)\ndefer f.Close()\n_ = err")
		if got := resourceCallInStmt(list[0], list, 0, fset); got != nil {
			t.Errorf("expected nil when defer Close present, got %+v", got)
		}
	})

	t.Run("NotAssign", func(t *testing.T) {
		list := blockStmts(t, "os.Open(p)")
		if got := resourceCallInStmt(list[0], list, 0, fset); got != nil {
			t.Errorf("expected nil for non-assign stmt, got %+v", got)
		}
	})

	t.Run("RHSNotCall", func(t *testing.T) {
		list := blockStmts(t, "f := 5")
		if got := resourceCallInStmt(list[0], list, 0, fset); got != nil {
			t.Errorf("expected nil for non-call RHS, got %+v", got)
		}
	})

	t.Run("NotResourceCall", func(t *testing.T) {
		list := blockStmts(t, "x := compute()")
		if got := resourceCallInStmt(list[0], list, 0, fset); got != nil {
			t.Errorf("expected nil for non-resource call, got %+v", got)
		}
	})

	t.Run("DiscardReceiver", func(t *testing.T) {
		list := blockStmts(t, "_ = os.Open(p)")
		if got := resourceCallInStmt(list[0], list, 0, fset); got != nil {
			t.Errorf("expected nil when receiver is _, got %+v", got)
		}
	})

	t.Run("NonIdentReceiver", func(t *testing.T) {
		list := blockStmts(t, "obj.f = os.Open(p)")
		if got := resourceCallInStmt(list[0], list, 0, fset); got != nil {
			t.Errorf("expected nil for non-ident receiver, got %+v", got)
		}
	})
}
