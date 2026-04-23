//ff:func feature=orchestrator type=parser control=iteration dimension=1
//ff:what ParseFile — QuerySpec.RowType 이 :one/:many 에 채워지고 :exec/:execresult 는 빈 값인지 회귀

package sqlc

import (
	"os"
	"path/filepath"
	"testing"
)

// TestParseFile_RowTypePerCardinality exercises the four sqlc cardinality
// macros in a single file and asserts RowType follows the "<Name>Row" rule
// for :one/:many and is empty for :exec/:execresult.
func TestParseFile_RowTypePerCardinality(t *testing.T) {
	tmp := t.TempDir()
	body := `-- name: UserFindByID :one
SELECT * FROM users WHERE id = @id;

-- name: UserList :many
SELECT * FROM users;

-- name: UserDelete :exec
DELETE FROM users WHERE id = @id;

-- name: UserUpdate :execresult
UPDATE users SET email = @email WHERE id = @id;
`
	path := filepath.Join(tmp, "users.sql")
	if err := os.WriteFile(path, []byte(body), 0644); err != nil {
		t.Fatalf("WriteFile: %v", err)
	}

	specs, diags := ParseFile(path)
	if len(diags) > 0 {
		t.Fatalf("unexpected diags: %v", diags)
	}
	if len(specs) != 4 {
		t.Fatalf("want 4 specs, got %d (%+v)", len(specs), specs)
	}

	byName := make(map[string]QuerySpec, len(specs))
	for _, s := range specs {
		byName[s.Name] = s
	}

	cases := []struct {
		name        string
		wantRowType string
	}{
		{"UserFindByID", "UserFindByIDRow"},
		{"UserList", "UserListRow"},
		{"UserDelete", ""},
		{"UserUpdate", ""},
	}
	for _, tc := range cases {
		got, ok := byName[tc.name]
		if !ok {
			t.Fatalf("spec %q missing", tc.name)
		}
		if got.RowType != tc.wantRowType {
			t.Errorf("%s: RowType = %q, want %q", tc.name, got.RowType, tc.wantRowType)
		}
	}
}
