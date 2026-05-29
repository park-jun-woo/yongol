//ff:func feature=migration type=test control=sequence
//ff:what TestBuildASTFromSQL_ForeignKey — FK inline 파싱 + ON DELETE CASCADE + 이름 규칙
package migration

import "testing"

func TestBuildASTFromSQL_ForeignKey(t *testing.T) {
	sql := `
CREATE TABLE organizations (id BIGSERIAL PRIMARY KEY);
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    org_id BIGINT NOT NULL REFERENCES organizations(id) ON DELETE CASCADE
);
`
	s := NewSchema()
	if err := BuildASTFromSQL(s, sql); err != nil {
		t.Fatal(err)
	}
	u := s.Tables["users"]
	if u == nil || len(u.ForeignKeys) != 1 {
		t.Fatalf("expected 1 FK on users, got %+v", u)
	}
	fk := u.ForeignKeys[0]
	if fk.RefTable != "organizations" || fk.RefColumns[0] != "id" ||
		fk.OnDelete != "CASCADE" {
		t.Errorf("FK: %+v", fk)
	}
	if fk.Name != "users_org_id_fkey" {
		t.Errorf("FK name: %q", fk.Name)
	}
}
