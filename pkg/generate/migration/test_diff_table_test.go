//ff:func feature=migration type=test control=iteration dimension=1
//ff:what Diff 테이블 단위 — 신규 / 삭제 / 변경 매트릭스
package migration

import "testing"

func TestDiff_CreateTable(t *testing.T) {
	prev := NewSchema()
	curr := mustAST(t, `CREATE TABLE users (id BIGSERIAL PRIMARY KEY);`)
	ops := Diff(prev, curr, nil)
	if len(ops) == 0 {
		t.Fatalf("expected CreateTable op, got none")
	}
	if _, ok := ops[0].(CreateTable); !ok {
		t.Errorf("op[0] should be CreateTable, got %T", ops[0])
	}
}

func TestDiff_DropTable(t *testing.T) {
	prev := mustAST(t, `CREATE TABLE users (id BIGSERIAL PRIMARY KEY);`)
	curr := NewSchema()
	ops := Diff(prev, curr, nil)
	foundDrop := false
	for _, op := range ops {
		if _, ok := op.(DropTable); ok {
			foundDrop = true
		}
	}
	if !foundDrop {
		t.Errorf("DropTable missing from ops: %+v", ops)
	}
}

func TestDiff_Initial_AllCreates(t *testing.T) {
	curr := mustAST(t, `
CREATE TABLE orgs (id BIGSERIAL PRIMARY KEY);
CREATE TABLE users (id BIGSERIAL PRIMARY KEY, org_id BIGINT REFERENCES orgs(id));
`)
	ops := Diff(NewSchema(), curr, nil)
	if len(ops) < 3 {
		t.Fatalf("expected >=3 ops (2 CreateTable + 1 AddFK), got %d: %+v", len(ops), ops)
	}
	// CreateTable phase (7) must precede AddForeignKey phase (12).
	createOrgIdx := -1
	addFKIdx := -1
	for i, op := range ops {
		if ct, ok := op.(CreateTable); ok && ct.Table.Name == "orgs" {
			createOrgIdx = i
		}
		if _, ok := op.(AddForeignKey); ok {
			addFKIdx = i
		}
	}
	if createOrgIdx < 0 || addFKIdx < 0 {
		t.Fatalf("missing ops: createOrgIdx=%d addFKIdx=%d", createOrgIdx, addFKIdx)
	}
	if createOrgIdx >= addFKIdx {
		t.Errorf("CreateTable must precede AddForeignKey")
	}
}
