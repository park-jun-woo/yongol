//ff:func feature=migration type=test control=sequence
//ff:what TestParseIdentityCaseInsensitive — 소문자/혼합 대소문자 IDENTITY 파싱
package migration

import "testing"

func TestParseIdentityCaseInsensitive(t *testing.T) {
	sql := `
CREATE TABLE users (
    id bigint generated always as identity PRIMARY KEY,
    email TEXT NOT NULL
);
`
	s := NewSchema()
	if err := BuildASTFromSQL(s, sql); err != nil {
		t.Fatalf("parse: %v", err)
	}
	tbl := s.Tables["users"]
	if tbl == nil {
		t.Fatalf("users table not found")
	}
	idCol := tbl.Columns[0]
	if idCol.Identity == nil {
		t.Fatalf("Identity should be set, got nil")
	}
	if !idCol.Identity.Always {
		t.Errorf("Identity.Always should be true")
	}
	if idCol.Nullable {
		t.Errorf("IDENTITY column must be NOT NULL")
	}
}
