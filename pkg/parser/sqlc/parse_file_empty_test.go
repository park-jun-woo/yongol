//ff:func feature=orchestrator type=test control=sequence
//ff:what ParseFile — 빈 파일은 빈 결과 반환

package sqlc

import "testing"

func TestParseFile_Empty(t *testing.T) {
	tmp := t.TempDir()
	path := writeSQL(t, tmp, "users.sql", "")
	specs, diags := ParseFile(path)
	if len(diags) > 0 {
		t.Fatalf("unexpected diags: %v", diags)
	}
	if len(specs) != 0 {
		t.Fatalf("want 0 specs, got %d", len(specs))
	}
}
