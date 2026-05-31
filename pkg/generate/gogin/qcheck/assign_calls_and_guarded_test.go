//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestAssignCallsAndGuarded — guarded shape true + 각 거부 분기 검증
package qcheck

import (
	"go/ast"
	"testing"
)

func TestAssignCallsAndGuarded(t *testing.T) {
	t.Run("Guarded", func(t *testing.T) {
		list := blockStmts(t, "err := pkg.Func()\nif err != nil { return }")
		assign := list[0].(*ast.AssignStmt)
		if !assignCallsAndGuarded(assign, "pkg", "Func", list, 0) {
			t.Errorf("expected true for guarded shape")
		}
	})

	t.Run("MultiRHS", func(t *testing.T) {
		list := blockStmts(t, "a, err := 1, pkg.Func()\nif err != nil { return }")
		// Two LHS but single RHS? Use truly multi RHS:
		list2 := blockStmts(t, "a, b := pkg.Func(), pkg.Func()\nif err != nil { return }")
		assign := list2[0].(*ast.AssignStmt)
		if assignCallsAndGuarded(assign, "pkg", "Func", list2, 0) {
			t.Errorf("expected false for multi-RHS")
		}
		_ = list
	})

	t.Run("NonCallRHS", func(t *testing.T) {
		list := blockStmts(t, "err := 5\nif err != nil { return }")
		assign := list[0].(*ast.AssignStmt)
		if assignCallsAndGuarded(assign, "pkg", "Func", list, 0) {
			t.Errorf("expected false for non-call RHS")
		}
	})

	t.Run("WrongSelector", func(t *testing.T) {
		list := blockStmts(t, "err := other.Other()\nif err != nil { return }")
		assign := list[0].(*ast.AssignStmt)
		if assignCallsAndGuarded(assign, "pkg", "Func", list, 0) {
			t.Errorf("expected false for wrong selector")
		}
	})

	t.Run("NoErrLHS", func(t *testing.T) {
		list := blockStmts(t, "x := pkg.Func()\nif err != nil { return }")
		assign := list[0].(*ast.AssignStmt)
		if assignCallsAndGuarded(assign, "pkg", "Func", list, 0) {
			t.Errorf("expected false when LHS has no err")
		}
	})

	t.Run("NoFollowingStmt", func(t *testing.T) {
		list := blockStmts(t, "err := pkg.Func()")
		assign := list[0].(*ast.AssignStmt)
		if assignCallsAndGuarded(assign, "pkg", "Func", list, 0) {
			t.Errorf("expected false when no following guard stmt")
		}
	})

	t.Run("FollowingNotGuard", func(t *testing.T) {
		list := blockStmts(t, "err := pkg.Func()\n_ = err")
		assign := list[0].(*ast.AssignStmt)
		if assignCallsAndGuarded(assign, "pkg", "Func", list, 0) {
			t.Errorf("expected false when following stmt is not an err guard")
		}
	})
}
