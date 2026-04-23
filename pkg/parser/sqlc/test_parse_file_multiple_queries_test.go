//ff:func feature=orchestrator type=test control=sequence
//ff:what ParseFile — 한 파일에 3개 쿼리 + Params 추출 회귀

package sqlc

import "testing"

func TestParseFile_MultipleQueries(t *testing.T) {
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
	if specs[0].Name != "UserCreate" || specs[0].Cardinality != "one" {
		t.Errorf("spec[0] mismatch: %+v", specs[0])
	}
	if len(specs[0].Params) != 1 || specs[0].Params[0] != "Email" {
		t.Errorf("spec[0].Params = %v, want [Email]", specs[0].Params)
	}
	if specs[1].Name != "UserFindByID" {
		t.Errorf("spec[1].Name = %q, want UserFindByID", specs[1].Name)
	}
	if len(specs[1].Params) != 1 || specs[1].Params[0] != "ID" {
		t.Errorf("spec[1].Params = %v, want [ID]", specs[1].Params)
	}
	if specs[2].Name != "UserList" || specs[2].Cardinality != "many" {
		t.Errorf("spec[2] mismatch: %+v", specs[2])
	}
	if len(specs[2].Params) != 0 {
		t.Errorf("spec[2].Params = %v, want []", specs[2].Params)
	}
}
