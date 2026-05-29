//ff:func feature=orchestrator type=test control=sequence
//ff:what ParseFile — `-- name: UserFindByID :one` 한 줄로 QuerySpec 추출

package sqlc

import "testing"

func TestParseFile_ValidSingleQuery(t *testing.T) {
	tmp := t.TempDir()
	path := writeSQL(t, tmp, "users.sql", `-- name: UserFindByID :one
SELECT * FROM users WHERE id = @id;
`)
	specs, diags := ParseFile(path)
	if len(diags) > 0 {
		t.Fatalf("unexpected diags: %v", diags)
	}
	if len(specs) != 1 {
		t.Fatalf("want 1 spec, got %d", len(specs))
	}
	s := specs[0]
	if s.Name != "UserFindByID" || s.Cardinality != "one" || s.Model != "User" {
		t.Errorf("spec mismatch: %+v", s)
	}
	if s.Method != "FindByID" {
		t.Errorf("Method = %q, want %q", s.Method, "FindByID")
	}
	if len(s.Params) != 1 || s.Params[0] != "ID" {
		t.Errorf("Params = %v, want [ID]", s.Params)
	}
	if s.Line != 1 {
		t.Errorf("Line = %d, want 1", s.Line)
	}
}
