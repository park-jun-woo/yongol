//ff:func feature=migration type=test control=iteration dimension=1
//ff:what TestDiff_TableDrop_FKFirst — 테이블 DROP 전에 FK DROP 이 먼저 나와야 함
package migration

import "testing"

func TestDiff_TableDrop_FKFirst(t *testing.T) {
	prev := mustAST(t, `
CREATE TABLE orgs (id BIGSERIAL PRIMARY KEY);
CREATE TABLE users (id BIGSERIAL PRIMARY KEY, org_id BIGINT REFERENCES orgs(id));`)
	curr := mustAST(t, `
CREATE TABLE orgs (id BIGSERIAL PRIMARY KEY);`)
	ops := Diff(prev, curr, nil)
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
