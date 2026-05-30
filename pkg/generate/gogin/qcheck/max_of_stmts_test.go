//ff:func feature=gen-gogin type=test control=iteration topic=depth-report
//ff:what TestMaxOfStmts — 빈 리스트/더 깊은 stmt/얕은 stmt에서 최대 depth 반환 검증

package qcheck

import "testing"

func TestMaxOfStmts(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		if got := maxOfStmts(nil, 3); got != 3 {
			t.Errorf("empty list = %d, want 3 (base depth)", got)
		}
	})

	t.Run("DeeperStmt", func(t *testing.T) {
		// One flat stmt + one with a nested for -> for adds a level.
		list := bodyBlock(t, "_ = 1\nfor { break }").List
		if got := maxOfStmts(list, 0); got != 1 {
			t.Errorf("DeeperStmt = %d, want 1", got)
		}
	})

	t.Run("AllFlat", func(t *testing.T) {
		list := bodyBlock(t, "_ = 1\n_ = 2").List
		if got := maxOfStmts(list, 2); got != 2 {
			t.Errorf("AllFlat = %d, want 2 (base depth unchanged)", got)
		}
	})
}
