//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestScanCallInStmt — 미가드 .Scan() DF-02 + if-init/assign-guard/무호출 스킵 분기 검증
package qcheck

import (
	"go/token"
	"testing"
)

func TestScanCallInStmt(t *testing.T) {
	fset := token.NewFileSet()

	t.Run("UnguardedFinding", func(t *testing.T) {
		list := blockStmts(t, "_ = r.Scan(x)")
		got := scanCallInStmt(list[0], list, 0, fset)
		if len(got) != 1 || got[0].Category != "DF-02" || got[0].Detail != "r.Scan" {
			t.Fatalf("expected one DF-02 r.Scan finding, got %+v", got)
		}
	})

	t.Run("IfInitGuard", func(t *testing.T) {
		list := blockStmts(t, "if err := r.Scan(x); err != nil { return }")
		if got := scanCallInStmt(list[0], list, 0, fset); got != nil {
			t.Errorf("expected nil for if-init guarded Scan, got %+v", got)
		}
	})

	t.Run("AssignGuard", func(t *testing.T) {
		list := blockStmts(t, "err := r.Scan(x)\nif err != nil { return }")
		if got := scanCallInStmt(list[0], list, 0, fset); got != nil {
			t.Errorf("expected nil for assign-guarded Scan, got %+v", got)
		}
	})

	t.Run("NoScanCall", func(t *testing.T) {
		list := blockStmts(t, "_ = r.Read(x)")
		if got := scanCallInStmt(list[0], list, 0, fset); got != nil {
			t.Errorf("expected nil when no Scan call present, got %+v", got)
		}
	})
}
