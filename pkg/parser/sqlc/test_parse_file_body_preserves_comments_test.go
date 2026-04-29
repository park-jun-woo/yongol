//ff:func feature=orchestrator type=test control=sequence
//ff:what ParseFile — Body 보관: 인라인 `-- ...` 주석을 지우지 않고 verbatim 보존

package sqlc

import (
	"strings"
	"testing"
)

// TestParseFile_Body_PreservesComments checks that inline `--` comments inside
// the SQL body survive into Body verbatim. The parser stays neutral; downstream
// consumers may strip them (e.g. extractReturningClause in XQS-20).
func TestParseFile_Body_PreservesComments(t *testing.T) {
	tmp := t.TempDir()
	path := writeSQL(t, tmp, "users.sql", `-- name: UserCreate :one
-- create one user
INSERT INTO users (email) VALUES (@email)
RETURNING id, email; -- partial RETURNING
`)
	specs, diags := ParseFile(path)
	if len(diags) > 0 {
		t.Fatalf("unexpected diags: %v", diags)
	}
	if len(specs) != 1 {
		t.Fatalf("want 1 spec, got %d", len(specs))
	}
	body := specs[0].Body
	if !strings.Contains(body, "-- create one user") {
		t.Errorf("body missing leading inline comment: %q", body)
	}
	if !strings.Contains(body, "-- partial RETURNING") {
		t.Errorf("body missing trailing inline comment: %q", body)
	}
	if !strings.Contains(body, "RETURNING id, email") {
		t.Errorf("body missing RETURNING clause: %q", body)
	}
}
