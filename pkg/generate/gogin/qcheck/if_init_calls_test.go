//ff:func feature=gen-gogin type=test control=branch topic=err-guard
//ff:what TestIfInitCalls — if-init이 pkg.Func 호출인지 각 분기 검증

package qcheck

import (
	"go/ast"
	"testing"
)

// firstIf returns the first *ast.IfStmt in the parsed func body.
func firstIf(t *testing.T, body string) *ast.IfStmt {
	t.Helper()
	for _, s := range blockStmts(t, body) {
		if ifs, ok := s.(*ast.IfStmt); ok {
			return ifs
		}
	}
	t.Fatalf("no if statement found in %q", body)
	return nil
}

func TestIfInitCalls(t *testing.T) {
	t.Run("Match", func(t *testing.T) {
		ifs := firstIf(t, "if err := json.Unmarshal(b, v); err != nil { return }")
		if !ifInitCalls(ifs, "json", "Unmarshal") {
			t.Errorf("expected true for matching if-init call")
		}
	})

	t.Run("NilInit", func(t *testing.T) {
		ifs := firstIf(t, "if cond { return }")
		if ifInitCalls(ifs, "json", "Unmarshal") {
			t.Errorf("expected false for nil init")
		}
	})

	t.Run("NonAssignInit", func(t *testing.T) {
		ifs := firstIf(t, "if x++; x > 0 { return }")
		if ifInitCalls(ifs, "json", "Unmarshal") {
			t.Errorf("expected false for non-assign init")
		}
	})

	t.Run("MultiRHS", func(t *testing.T) {
		ifs := firstIf(t, "if a, b := json.Unmarshal(x, y), 1; b > 0 { return }")
		if ifInitCalls(ifs, "json", "Unmarshal") {
			t.Errorf("expected false for multi-RHS init")
		}
	})

	t.Run("NonCallRHS", func(t *testing.T) {
		ifs := firstIf(t, "if v := 5; v > 0 { return }")
		if ifInitCalls(ifs, "json", "Unmarshal") {
			t.Errorf("expected false for non-call RHS")
		}
	})

	t.Run("WrongSelector", func(t *testing.T) {
		ifs := firstIf(t, "if err := json.Marshal(v); err != nil { return }")
		if ifInitCalls(ifs, "json", "Unmarshal") {
			t.Errorf("expected false for wrong func name")
		}
	})

	t.Run("NilStmt", func(t *testing.T) {
		if ifInitCalls(nil, "json", "Unmarshal") {
			t.Errorf("expected false for nil if-stmt")
		}
	})
}
