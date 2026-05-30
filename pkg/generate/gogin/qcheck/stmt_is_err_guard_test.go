//ff:func feature=gen-gogin type=test control=branch topic=err-guard
//ff:what TestStmtIsErrGuard — if err != nil 가드 true + 비-if/비-binary/비-NEQ false 분기 검증

package qcheck

import "testing"

func TestStmtIsErrGuard(t *testing.T) {
	t.Run("ErrGuard", func(t *testing.T) {
		list := blockStmts(t, "if err != nil { return }")
		if !stmtIsErrGuard(list[0]) {
			t.Errorf("expected true for `if err != nil`")
		}
	})

	t.Run("NotIfStmt", func(t *testing.T) {
		list := blockStmts(t, "x := 1")
		if stmtIsErrGuard(list[0]) {
			t.Errorf("expected false for non-if stmt")
		}
	})

	t.Run("CondNotBinary", func(t *testing.T) {
		list := blockStmts(t, "if cond { return }")
		if stmtIsErrGuard(list[0]) {
			t.Errorf("expected false for non-binary cond")
		}
	})

	t.Run("WrongOperator", func(t *testing.T) {
		list := blockStmts(t, "if err == nil { return }")
		if stmtIsErrGuard(list[0]) {
			t.Errorf("expected false for == operator")
		}
	})

	t.Run("NotErrNilCheck", func(t *testing.T) {
		list := blockStmts(t, "if x != nil { return }")
		if stmtIsErrGuard(list[0]) {
			t.Errorf("expected false for non-err identifier")
		}
	})
}
