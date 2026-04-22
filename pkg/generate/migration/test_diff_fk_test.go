//ff:func feature=migration type=test control=iteration dimension=1
//ff:what Diff FK — 추가/삭제/ON DELETE 변경 + 테이블 drop 전 FK drop 순서
package migration

import "testing"

func TestDiff_FK_Add(t *testing.T) {
	prev := mustAST(t, `
CREATE TABLE orgs (id BIGSERIAL PRIMARY KEY);
CREATE TABLE users (id BIGSERIAL PRIMARY KEY, org_id BIGINT NOT NULL);`)
	curr := mustAST(t, `
CREATE TABLE orgs (id BIGSERIAL PRIMARY KEY);
CREATE TABLE users (id BIGSERIAL PRIMARY KEY, org_id BIGINT NOT NULL REFERENCES orgs(id));`)
	ops := Diff(prev, curr, nil)
	if len(ops) != 1 {
		t.Fatalf("expected 1 op, got %d: %+v", len(ops), ops)
	}
	if _, ok := ops[0].(AddForeignKey); !ok {
		t.Errorf("expected AddForeignKey, got %T", ops[0])
	}
}

func TestDiff_FK_OnDeleteChange(t *testing.T) {
	prev := mustAST(t, `
CREATE TABLE orgs (id BIGSERIAL PRIMARY KEY);
CREATE TABLE users (org_id BIGINT REFERENCES orgs(id) ON DELETE RESTRICT);`)
	curr := mustAST(t, `
CREATE TABLE orgs (id BIGSERIAL PRIMARY KEY);
CREATE TABLE users (org_id BIGINT REFERENCES orgs(id) ON DELETE CASCADE);`)
	ops := Diff(prev, curr, nil)
	// Expect DropFK then AddFK.
	if len(ops) < 2 {
		t.Fatalf("expected >=2 ops, got %d: %+v", len(ops), ops)
	}
	if _, ok := ops[0].(DropForeignKey); !ok {
		t.Errorf("op[0] should be DropForeignKey, got %T", ops[0])
	}
	foundAdd := false
	for _, op := range ops {
		if _, ok := op.(AddForeignKey); ok {
			foundAdd = true
		}
	}
	if !foundAdd {
		t.Errorf("AddForeignKey not found: %+v", ops)
	}
}

func TestDiff_TableDrop_FKFirst(t *testing.T) {
	prev := mustAST(t, `
CREATE TABLE orgs (id BIGSERIAL PRIMARY KEY);
CREATE TABLE users (id BIGSERIAL PRIMARY KEY, org_id BIGINT REFERENCES orgs(id));`)
	curr := mustAST(t, `
CREATE TABLE orgs (id BIGSERIAL PRIMARY KEY);`)
	ops := Diff(prev, curr, nil)
	// Ensure DropForeignKey precedes DropTable in the ordered list.
	dropFKIdx, dropTableIdx := -1, -1
	for i, op := range ops {
		if _, ok := op.(DropForeignKey); ok && dropFKIdx < 0 {
			dropFKIdx = i
		}
		if _, ok := op.(DropTable); ok && dropTableIdx < 0 {
			dropTableIdx = i
		}
	}
	if dropFKIdx < 0 {
		t.Fatalf("DropForeignKey missing: %+v", ops)
	}
	if dropTableIdx < 0 {
		t.Fatalf("DropTable missing: %+v", ops)
	}
	if dropFKIdx > dropTableIdx {
		t.Errorf("DropFK must come before DropTable, got FK=%d table=%d", dropFKIdx, dropTableIdx)
	}
}
