//ff:func feature=orchestrator type=test control=iteration dimension=1
//ff:what ParseFile — Body 보관: 다중 쿼리, 각 쿼리가 자기 본문만 가짐 (-- name: 미포함)

package sqlc

import (
	"strings"
	"testing"
)

// TestParseFile_Body_MultipleQueries verifies each query gets its own Body
// scoped between `-- name:` markers and that the marker line itself is not
// included in any Body.
func TestParseFile_Body_MultipleQueries(t *testing.T) {
	tmp := t.TempDir()
	path := writeSQL(t, tmp, "users.sql", `-- name: UserCreate :one
INSERT INTO users (email) VALUES (@email) RETURNING *;

-- name: UserFindByID :one
SELECT * FROM users WHERE id = @id;

-- name: UserList :many
SELECT * FROM users;
`)
	specs, diags := ParseFile(path)
	if len(diags) > 0 {
		t.Fatalf("unexpected diags: %v", diags)
	}
	if len(specs) != 3 {
		t.Fatalf("want 3 specs, got %d", len(specs))
	}
	checks := []struct {
		mustHave string
		mustOmit string
	}{
		{"INSERT INTO users", "SELECT *"},
		{"SELECT * FROM users WHERE id", "INSERT"},
		{"SELECT * FROM users;", "INSERT"},
	}
	for i, c := range checks {
		body := specs[i].Body
		if !strings.Contains(body, c.mustHave) {
			t.Errorf("specs[%d].Body missing %q: %q", i, c.mustHave, body)
		}
		if c.mustOmit != "" && strings.Contains(body, c.mustOmit) {
			t.Errorf("specs[%d].Body unexpectedly contains %q: %q", i, c.mustOmit, body)
		}
		if strings.Contains(body, "-- name:") {
			t.Errorf("specs[%d].Body must not contain `-- name:` marker: %q", i, body)
		}
	}
}
