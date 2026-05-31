//ff:func feature=migration type=test control=sequence
//ff:what apply*Hint 단위 테스트 — 함수명 컨벤션에 맞춘 직접 커버 (전 분기)
package migration

import (
	"testing"
)

func TestApplyRenameHints(t *testing.T) {
	prev := mustAST(t, `CREATE TABLE old_users (id BIGINT PRIMARY KEY, name TEXT NOT NULL);`)
	h := newEmptyHints()
	h.RenameTables = []RenameTableHint{{From: "old_users", To: "users"}}
	h.RenameColumns = []RenameColumnHint{{Table: "users", From: "name", To: "full_name"}}
	out := applyRenameHints(prev, h)
	if _, ok := out.Tables["users"]; !ok {
		t.Errorf("table not renamed to users: %v", out.Tables)
	}
	// nil hints → unchanged.
	if applyRenameHints(prev, nil) != prev {
		t.Errorf("nil hints should return prev unchanged")
	}
	// no rename rules → unchanged.
	if applyRenameHints(prev, newEmptyHints()) != prev {
		t.Errorf("empty rules should return prev unchanged")
	}
}
