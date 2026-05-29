//ff:func feature=orchestrator type=test control=sequence
//ff:what ParseDir — 복수 sql 파일 합산 결과 검증 (non-sql 파일은 무시)

package sqlc

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseDir_MultipleFiles(t *testing.T) {
	tmp := t.TempDir()

	if err := os.WriteFile(filepath.Join(tmp, "users.sql"), []byte(`-- name: UserFindByID :one
SELECT * FROM users WHERE id = @id;
`), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(tmp, "workflows.sql"), []byte(`-- name: WorkflowList :many
SELECT * FROM workflows WHERE org_id = @org_id;

-- name: WorkflowCreate :one
INSERT INTO workflows (name) VALUES (@name) RETURNING *;
`), 0644); err != nil {
		t.Fatal(err)
	}
	// non-sql file should be ignored
	if err := os.WriteFile(filepath.Join(tmp, "ignore.txt"), []byte("not sql"), 0644); err != nil {
		t.Fatal(err)
	}

	specs, diags := ParseDir(tmp)
	if len(diags) > 0 {
		t.Fatalf("unexpected diags: %v", diags)
	}
	if len(specs) != 3 {
		t.Fatalf("want 3 specs (1 users + 2 workflows), got %d: %+v", len(specs), specs)
	}
}
