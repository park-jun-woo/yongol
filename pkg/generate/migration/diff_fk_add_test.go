//ff:func feature=migration type=test control=sequence
//ff:what TestDiff_FK_Add — 새 FK 추가 시 AddForeignKey 생성
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
