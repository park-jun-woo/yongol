//ff:func feature=migration type=test control=iteration dimension=1
//ff:what rename hints 적용 전후의 Diff 동작 검증
package migration

import "testing"

func TestDiff_RenameColumnHint(t *testing.T) {
	prev := mustAST(t, `CREATE TABLE users (id BIGSERIAL PRIMARY KEY, email VARCHAR(255) NOT NULL);`)
	curr := mustAST(t, `CREATE TABLE users (id BIGSERIAL PRIMARY KEY, email_address VARCHAR(255) NOT NULL);`)
	// Without hints: expect drop + add
	opsNoHint := Diff(prev, curr, nil)
	hasDrop, hasAdd := false, false
	for _, op := range opsNoHint {
		if _, ok := op.(DropColumn); ok {
			hasDrop = true
		}
		if _, ok := op.(AddColumn); ok {
			hasAdd = true
		}
	}
	if !hasDrop || !hasAdd {
		t.Errorf("without hints expected Drop + Add; got %+v", opsNoHint)
	}

	// With rename hint: expect single RenameColumn
	hints := &Hints{
		RenameColumns: []RenameColumnHint{{Table: "users", From: "email", To: "email_address"}},
	}
	opsWithHint := Diff(prev, curr, hints)
	foundRename := false
	for _, op := range opsWithHint {
		if r, ok := op.(RenameColumn); ok {
			foundRename = true
			if r.From != "email" || r.To != "email_address" {
				t.Errorf("rename wrong: %+v", r)
			}
		}
		if _, ok := op.(DropColumn); ok {
			t.Errorf("unexpected DropColumn with rename hint: %+v", opsWithHint)
		}
	}
	if !foundRename {
		t.Errorf("RenameColumn missing: %+v", opsWithHint)
	}
}

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
