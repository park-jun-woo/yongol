//ff:func feature=migration type=test control=iteration dimension=1
//ff:what TestDiff_RenameTableHint — rename 힌트 있을 때 RenameTable Op 생성 + drop/create 없음
package migration

import "testing"

func TestDiff_RenameTableHint(t *testing.T) {
	prev := mustAST(t, `CREATE TABLE members (id BIGSERIAL PRIMARY KEY);`)
	curr := mustAST(t, `CREATE TABLE users (id BIGSERIAL PRIMARY KEY);`)
	hints := &Hints{
		RenameTables: []RenameTableHint{{From: "members", To: "users"}},
	}
	ops := Diff(prev, curr, hints)
	foundRename := false
	for _, op := range ops {
		if _, ok := op.(RenameTable); ok {
			foundRename = true
		}
		if _, ok := op.(DropTable); ok {
			t.Errorf("unexpected DropTable with rename hint: %+v", ops)
		}
		if _, ok := op.(CreateTable); ok {
			t.Errorf("unexpected CreateTable with rename hint: %+v", ops)
		}
	}
	if !foundRename {
		t.Errorf("RenameTable missing: %+v", ops)
	}
}
