//ff:func feature=gen-gogin type=test control=iteration dimension=2
//ff:what methodGen.assignOp 단위 테스트 (FirstErr + DeclaredVars 추적으로 := / = 선택)

package ssac

import "testing"

func TestMethodGenAssignOp(t *testing.T) {
	t.Run("err-only first := then =", func(t *testing.T) {
		g := &methodGen{FirstErr: true, DeclaredVars: map[string]bool{}}
		if got := g.assignOp(false, ""); got != ":=" {
			t.Errorf("first err-only = %q, want :=", got)
		}
		if got := g.assignOp(false, ""); got != "=" {
			t.Errorf("second err-only = %q, want =", got)
		}
	})
	t.Run("new var always :=", func(t *testing.T) {
		g := &methodGen{FirstErr: false, DeclaredVars: map[string]bool{}}
		if got := g.assignOp(true, "row"); got != ":=" {
			t.Errorf("new var = %q, want :=", got)
		}
	})
	t.Run("re-declared var becomes =", func(t *testing.T) {
		g := &methodGen{FirstErr: true, DeclaredVars: map[string]bool{}}
		if got := g.assignOp(true, "row"); got != ":=" {
			t.Errorf("first row = %q, want :=", got)
		}
		if got := g.assignOp(true, "row"); got != "=" {
			t.Errorf("repeat row = %q, want =", got)
		}
	})
	t.Run("blank identifier not tracked", func(t *testing.T) {
		g := &methodGen{FirstErr: true, DeclaredVars: map[string]bool{}}
		if got := g.assignOp(true, "_"); got != ":=" {
			t.Errorf("blank with new var = %q, want :=", got)
		}
	})
}
