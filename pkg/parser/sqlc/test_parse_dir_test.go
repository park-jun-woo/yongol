//ff:func feature=orchestrator type=parser control=iteration dimension=1
//ff:what ParseDir 통합 테스트 — 복수 파일 합산 / 없는 디렉토리 / 파일 존재 확인
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

func TestParseDir_MissingDir(t *testing.T) {
	// non-existent directory is not an error (empty result)
	specs, diags := ParseDir("/nonexistent/dir/for/sqlc/parse/dir")
	if len(specs) != 0 {
		t.Errorf("want 0 specs, got %d", len(specs))
	}
	if len(diags) != 0 {
		t.Errorf("want 0 diags, got %d: %v", len(diags), diags)
	}
}
