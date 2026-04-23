//ff:func feature=migration type=test-helper control=iteration dimension=1
//ff:what assertRenameColumnWithHints — rename 힌트 적용 시 단일 RenameColumn 생성 검증
package migration

import "testing"

func assertRenameColumnWithHints(t *testing.T, prev, curr *Schema, hints *Hints) {
	t.Helper()
	ops := Diff(prev, curr, hints)
	foundRename := false
	for _, op := range ops {
		foundRename = checkRenameColumnOp(t, op, foundRename, ops)
	}
	if !foundRename {
		t.Errorf("RenameColumn missing: %+v", ops)
	}
}
