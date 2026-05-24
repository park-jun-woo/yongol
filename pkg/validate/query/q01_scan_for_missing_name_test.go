//ff:func feature=validate type=test control=sequence topic=query-structural
//ff:what q01ScanForMissingName — name 누락 SQL 파일 감지 (정상/누락/빈 파일/주석만/파일 없음) 검증

package query

import (
	"os"
	"path/filepath"
	"testing"
)

func TestQ01ScanForMissingName(t *testing.T) {
	writeFile := func(t *testing.T, content string) string {
		t.Helper()
		dir := t.TempDir()
		p := filepath.Join(dir, "query.sql")
		if err := os.WriteFile(p, []byte(content), 0o644); err != nil {
			t.Fatal(err)
		}
		return p
	}

	t.Run("file not found returns false", func(t *testing.T) {
		_, found := q01ScanForMissingName("/nonexistent/file.sql")
		if found {
			t.Error("expected false for nonexistent file")
		}
	})

	t.Run("empty file returns false", func(t *testing.T) {
		p := writeFile(t, "")
		_, found := q01ScanForMissingName(p)
		if found {
			t.Error("expected false for empty file")
		}
	})

	t.Run("only comments returns false", func(t *testing.T) {
		p := writeFile(t, "-- just a comment\n-- another comment\n")
		_, found := q01ScanForMissingName(p)
		if found {
			t.Error("expected false for comments-only file")
		}
	})

	t.Run("proper name annotation returns false", func(t *testing.T) {
		p := writeFile(t, "-- name: GetUser :one\nSELECT * FROM users WHERE id = $1;\n")
		_, found := q01ScanForMissingName(p)
		if found {
			t.Error("expected false for properly annotated file")
		}
	})

	t.Run("SQL before name annotation does not fire", func(t *testing.T) {
		// Once a name annotation is encountered, sawName becomes true and
		// the final check (pendingSQL && !sawName) is false.
		p := writeFile(t, "SELECT 1;\n-- name: GetUser :one\nSELECT * FROM users;\n")
		_, found := q01ScanForMissingName(p)
		if found {
			t.Error("expected false when name annotation follows SQL")
		}
	})

	t.Run("SQL with leading blank lines before name does not fire", func(t *testing.T) {
		p := writeFile(t, "\n\nSELECT 1;\n-- name: GetUser :one\nSELECT * FROM users;\n")
		_, found := q01ScanForMissingName(p)
		if found {
			t.Error("expected false when name annotation follows SQL")
		}
	})

	t.Run("only SQL no name at all fires", func(t *testing.T) {
		p := writeFile(t, "SELECT * FROM users;\n")
		line, found := q01ScanForMissingName(p)
		if !found {
			t.Fatal("expected true for SQL without any name")
		}
		if line != 1 {
			t.Errorf("expected line 1, got %d", line)
		}
	})

	t.Run("comment then SQL without name fires", func(t *testing.T) {
		p := writeFile(t, "-- description\nSELECT 1;\n")
		line, found := q01ScanForMissingName(p)
		if !found {
			t.Fatal("expected true")
		}
		if line != 2 {
			t.Errorf("expected line 2, got %d", line)
		}
	})
}
