//ff:func feature=orchestrator type=test control=sequence
//ff:what TestParseDDLIfPresent — DDL 미탐지(early return) + 탐지 시 results/tables/queries 파싱
package yongol

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseDDLIfPresent_Present(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "users.sql"),
		[]byte("CREATE TABLE users (id BIGSERIAL PRIMARY KEY);"), 0o644); err != nil {
		t.Fatal(err)
	}
	// queries subdir with a valid sqlc query.
	qdir := filepath.Join(dir, "queries")
	if err := os.MkdirAll(qdir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(qdir, "q.sql"),
		[]byte("-- name: GetUser :one\nSELECT id FROM users WHERE id = $1;\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	fs := &Fullstack{}
	has := map[SSOTKind]DetectedSSOT{
		KindDDL: {Kind: KindDDL, Path: dir, Presence: SSOTPopulated},
	}
	parseDDLIfPresent(fs, has)

	if len(fs.DDLResults) == 0 {
		t.Errorf("expected DDLResults populated")
	}
	if len(fs.DDLTables) == 0 {
		t.Errorf("expected DDLTables populated")
	}
}
