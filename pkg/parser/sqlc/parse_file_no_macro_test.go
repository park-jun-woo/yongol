//ff:func feature=orchestrator type=test control=sequence
//ff:what ParseFile — `-- name:` 패턴이 없으면 빈 결과 반환

package sqlc

import "testing"

func TestParseFile_NoMacro(t *testing.T) {
	tmp := t.TempDir()
	path := writeSQL(t, tmp, "users.sql", `-- just a comment
SELECT 1;
`)
	specs, diags := ParseFile(path)
	if len(diags) > 0 {
		t.Fatalf("unexpected diags: %v", diags)
	}
	if len(specs) != 0 {
		t.Fatalf("want 0 specs, got %d", len(specs))
	}
}
