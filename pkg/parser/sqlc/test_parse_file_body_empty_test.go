//ff:func feature=orchestrator type=test control=sequence
//ff:what ParseFile — Body 보관: -- name: 직후 다른 -- name: 면 Body == ""

package sqlc

import (
	"strings"
	"testing"
)

// TestParseFile_Body_EmptyBody covers the edge case where a `-- name:` is
// declared but no SQL body follows (immediately followed by another marker
// or EOF). Body must be the empty string, not panic.
func TestParseFile_Body_EmptyBody(t *testing.T) {
	tmp := t.TempDir()
	path := writeSQL(t, tmp, "users.sql", `-- name: UserNoOp :exec
-- name: UserFindByID :one
SELECT * FROM users WHERE id = @id;
`)
	specs, diags := ParseFile(path)
	if len(diags) > 0 {
		t.Fatalf("unexpected diags: %v", diags)
	}
	if len(specs) != 2 {
		t.Fatalf("want 2 specs, got %d", len(specs))
	}
	if specs[0].Body != "" {
		t.Errorf("specs[0].Body = %q, want empty", specs[0].Body)
	}
	if !strings.Contains(specs[1].Body, "SELECT * FROM users") {
		t.Errorf("specs[1].Body missing SELECT: %q", specs[1].Body)
	}
}
