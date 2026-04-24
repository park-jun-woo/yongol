//ff:func feature=migration type=test control=sequence
//ff:what TestEmit_SentinelBeforeForeignKey — @sentinel INSERT 가 ALTER TABLE ADD FK 보다 앞에 배치되는지 확인
package migration

import (
	"strings"
	"testing"
)

func TestEmit_SentinelBeforeForeignKey(t *testing.T) {
	s := NewSchema()
	// organizations has a sentinel row; users has a FK -> organizations.
	sql := `
CREATE TABLE organizations (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

-- @sentinel
INSERT INTO organizations (id, name) OVERRIDING SYSTEM VALUE VALUES (0, 'system') ON CONFLICT DO NOTHING;

CREATE TABLE users (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    org_id BIGINT NOT NULL DEFAULT 0 REFERENCES organizations(id)
);
`
	if err := BuildASTFromSQL(s, sql); err != nil {
		t.Fatalf("BuildASTFromSQL: %v", err)
	}
	ops := Diff(NewSchema(), s, nil)
	out := EmitSQL(ops, EmitOptions{})
	insertIdx := strings.Index(out, "INSERT INTO organizations")
	fkIdx := strings.Index(out, "ADD CONSTRAINT")
	if insertIdx < 0 {
		t.Fatalf("sentinel INSERT missing:\n%s", out)
	}
	if fkIdx < 0 {
		t.Fatalf("ALTER ADD CONSTRAINT missing:\n%s", out)
	}
	if insertIdx >= fkIdx {
		t.Errorf("expected sentinel INSERT before ADD CONSTRAINT; insertIdx=%d fkIdx=%d\n%s",
			insertIdx, fkIdx, out)
	}
	// And: last CREATE TABLE must precede the first sentinel INSERT.
	createUsersIdx := strings.Index(out, "CREATE TABLE users")
	if createUsersIdx < 0 {
		t.Fatalf("CREATE TABLE users missing:\n%s", out)
	}
	if createUsersIdx >= insertIdx {
		t.Errorf("expected all CREATE TABLE before sentinel INSERT; createUsersIdx=%d insertIdx=%d\n%s",
			createUsersIdx, insertIdx, out)
	}
}
