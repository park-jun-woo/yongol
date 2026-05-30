//ff:func feature=gen-gogin type=test control=branch topic=depth-report
//ff:what TestIfStmtDepth — else 없음 / else가 더 깊음 / body가 더 깊음 분기 검증

package qcheck

import "testing"

func TestIfStmtDepth(t *testing.T) {
	t.Run("NoElse", func(t *testing.T) {
		// if { _ } -> body block adds one level: depth(0)+1 = 1.
		ifs := firstIf(t, "if c { _ = 1 }")
		if got := ifStmtDepth(ifs, 0); got != 1 {
			t.Errorf("NoElse depth = %d, want 1", got)
		}
	})

	t.Run("ElseDeeper", func(t *testing.T) {
		// body is shallow; else contains a nested if -> else deeper.
		ifs := firstIf(t, "if c { _ = 1 } else { if d { _ = 2 } }")
		if got := ifStmtDepth(ifs, 0); got != 2 {
			t.Errorf("ElseDeeper depth = %d, want 2", got)
		}
	})

	t.Run("BodyDeeper", func(t *testing.T) {
		// body contains a nested if; else shallow -> body deeper.
		ifs := firstIf(t, "if c { if d { _ = 1 } } else { _ = 2 }")
		if got := ifStmtDepth(ifs, 0); got != 2 {
			t.Errorf("BodyDeeper depth = %d, want 2", got)
		}
	})
}
