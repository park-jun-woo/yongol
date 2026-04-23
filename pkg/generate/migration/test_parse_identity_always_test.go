//ff:func feature=migration type=test control=sequence
//ff:what TestParseIdentityAlways — GENERATED ALWAYS AS IDENTITY 파싱 검증
package migration

import "testing"

func TestParseIdentityAlways(t *testing.T) {
	sql := `
CREATE TABLE users (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    email TEXT NOT NULL
);
`
	s := NewSchema()
	if err := BuildASTFromSQL(s, sql); err != nil {
		t.Fatalf("parse: %v", err)
	}
	tbl, ok := s.Tables["users"]
	if !ok {
		t.Fatalf("users table not found")
	}
	idCol := tbl.Columns[0]
	if idCol.Name != "id" || idCol.Type.Base != "BIGINT" {
		t.Fatalf("id column: %+v", idCol)
	}
	if idCol.Identity == nil {
		t.Fatalf("Identity should be set, got nil")
	}
	if !idCol.Identity.Always {
		t.Errorf("Identity.Always should be true, got false")
	}
	if idCol.Nullable {
		t.Errorf("IDENTITY column must be NOT NULL, got Nullable=true")
	}
	if idCol.Default != "" {
		t.Errorf("IDENTITY column should not have DEFAULT, got %q", idCol.Default)
	}
}
