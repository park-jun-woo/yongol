//ff:func feature=migration type=test control=iteration dimension=1
//ff:what TestDiff_Initial_AllCreates — 초기 모드 (CreateTable + AddForeignKey 순서 보장)
package migration

import "testing"

func TestDiff_Initial_AllCreates(t *testing.T) {
	curr := mustAST(t, `
CREATE TABLE orgs (id BIGSERIAL PRIMARY KEY);
CREATE TABLE users (id BIGSERIAL PRIMARY KEY, org_id BIGINT REFERENCES orgs(id));
`)
	ops := Diff(NewSchema(), curr, nil)
	if len(ops) < 3 {
		t.Fatalf("expected >=3 ops (2 CreateTable + 1 AddFK), got %d: %+v", len(ops), ops)
	}
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
