//ff:func feature=agent type=test control=sequence
//ff:what TestWriteSQLcQueryContext — sqlc 쿼리 이름 목록 기록, 부재 시 무기록 검증

package agent

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestWriteSQLcQueryContext(t *testing.T) {
	dir := t.TempDir()
	qdir := filepath.Join(dir, "db", "queries")
	if err := os.MkdirAll(qdir, 0o755); err != nil {
		t.Fatal(err)
	}
	content := "-- name: GetUser :one\nSELECT 1;\n-- name: ListUsers :many\nSELECT 2;\n"
	if err := os.WriteFile(filepath.Join(qdir, "users.sql"), []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}

	var b strings.Builder
	writeSQLcQueryContext(&b, dir, "users")
	out := b.String()
	if !strings.Contains(out, "sqlc queries:") || !strings.Contains(out, "-- name: GetUser :one") {
		t.Errorf("queries → %q", out)
	}

	// No query file: nothing.
	var b2 strings.Builder
	writeSQLcQueryContext(&b2, dir, "missing")
	if b2.Len() != 0 {
		t.Errorf("missing wrote %q", b2.String())
	}
}
