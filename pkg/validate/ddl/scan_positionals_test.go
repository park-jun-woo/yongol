//ff:func feature=validate type=test control=sequence topic=ddl-structural
//ff:what TestDDLHelpers — unit tests for the pure DDL validate helper functions
package ddl

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScanPositionals(t *testing.T) {
	tmp := t.TempDir()
	path := filepath.Join(tmp, "q.sql")
	content := "" +
		"-- name: GetUser :one\n" +
		"SELECT * FROM users\n" +
		"WHERE id = $1\n" +
		"  AND org = $2;\n" +
		"-- name: ListPosts :many\n" +
		"SELECT * FROM posts WHERE owner = $1;\n"
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
	// Query starts at line 1 (the first -- name: header). FindString returns
	// the first $N per line, so $1 (line 3) and $2 (line 4) are collected;
	// scanning stops at the next -- name: header (line 5).
	hits := scanPositionals(path, 1)
	if len(hits) != 2 {
		t.Fatalf("got %d hits, want 2 (stops at next -- name:)", len(hits))
	}
	if hits[0].param != "$1" || hits[1].param != "$2" {
		t.Errorf("params = %v", hits)
	}
	if hits[0].line != 3 {
		t.Errorf("first hit line = %d, want 3", hits[0].line)
	}

	// Nonexistent file → nil.
	if got := scanPositionals(filepath.Join(tmp, "missing.sql"), 1); got != nil {
		t.Errorf("expected nil for missing file, got %v", got)
	}
}
