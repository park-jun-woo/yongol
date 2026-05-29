//ff:func feature=migration type=test control=iteration dimension=1
//ff:what TestDiff_FK_OnDeleteChange — ON DELETE 변경 시 DropFK + AddFK 생성 순서
package migration

import "testing"

func TestDiff_FK_OnDeleteChange(t *testing.T) {
	prev := mustAST(t, `
CREATE TABLE orgs (id BIGSERIAL PRIMARY KEY);
CREATE TABLE users (org_id BIGINT REFERENCES orgs(id) ON DELETE RESTRICT);`)
	curr := mustAST(t, `
CREATE TABLE orgs (id BIGSERIAL PRIMARY KEY);
CREATE TABLE users (org_id BIGINT REFERENCES orgs(id) ON DELETE CASCADE);`)
	ops := Diff(prev, curr, nil)
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
