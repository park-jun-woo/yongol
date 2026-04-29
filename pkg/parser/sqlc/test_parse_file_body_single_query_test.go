//ff:func feature=orchestrator type=test control=sequence
//ff:what ParseFile — Body 보관: 단일 쿼리 + 한 줄 본문

package sqlc

import "testing"

// TestParseFile_Body_SingleQuery covers the simplest case: one query with one
// SQL line. Body must hold that line, no leading/trailing whitespace.
func TestParseFile_Body_SingleQuery(t *testing.T) {
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
	want := "SELECT * FROM users WHERE id = @id;"
	if specs[0].Body != want {
		t.Errorf("Body = %q, want %q", specs[0].Body, want)
	}
}
