//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestCheckSqlcOutPath — sqlc.yaml out 경로 검증의 read/parse/empty/mismatch/match 분기
package gogin

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestCheckSqlcOutPath(t *testing.T) {
	t.Run("ReadError", func(t *testing.T) {
		specs := t.TempDir()
		err := checkSqlcOutPath(specs, t.TempDir())
		if err == nil || !strings.Contains(err.Error(), "read ") {
			t.Errorf("expected read error, got: %v", err)
		}
	})

	t.Run("ParseError", func(t *testing.T) {
		specs := t.TempDir()
		writeSqlc(t, specs, "sql: [::::bad yaml")
		err := checkSqlcOutPath(specs, t.TempDir())
		if err == nil || !strings.Contains(err.Error(), "parse ") {
			t.Errorf("expected parse error, got: %v", err)
		}
	})

	t.Run("NoSQLEntries", func(t *testing.T) {
		specs := t.TempDir()
		writeSqlc(t, specs, "version: \"2\"\n")
		err := checkSqlcOutPath(specs, t.TempDir())
		if err == nil || !strings.Contains(err.Error(), "no sql entries") {
			t.Errorf("expected no sql entries error, got: %v", err)
		}
	})

	t.Run("EmptyOut", func(t *testing.T) {
		specs := t.TempDir()
		writeSqlc(t, specs, "sql:\n  - gen:\n      go:\n          out: \"\"\n")
		err := checkSqlcOutPath(specs, t.TempDir())
		if err == nil || !strings.Contains(err.Error(), "is empty") {
			t.Errorf("expected empty out error, got: %v", err)
		}
	})

	t.Run("Mismatch", func(t *testing.T) {
		specs := t.TempDir()
		writeSqlc(t, specs, "sql:\n  - gen:\n      go:\n          out: wrong/place\n")
		err := checkSqlcOutPath(specs, t.TempDir())
		if err == nil || !strings.Contains(err.Error(), "expected") {
			t.Errorf("expected mismatch error, got: %v", err)
		}
	})

	t.Run("MatchAbsolute", func(t *testing.T) {
		specs := t.TempDir()
		artifacts := t.TempDir()
		absArtifacts, _ := filepath.Abs(artifacts)
		expected := filepath.Join(absArtifacts, "backend", "internal", "db")
		writeSqlc(t, specs, "sql:\n  - gen:\n      go:\n          out: "+expected+"\n")
		if err := checkSqlcOutPath(specs, artifacts); err != nil {
			t.Errorf("expected match (nil error), got: %v", err)
		}
	})

	t.Run("MatchRelative", func(t *testing.T) {
		specs := t.TempDir()
		artifacts := t.TempDir()
		absArtifacts, _ := filepath.Abs(artifacts)
		absDB, _ := filepath.Abs(filepath.Join(specs, "db"))
		expected := filepath.Join(absArtifacts, "backend", "internal", "db")
		// Relative out that, joined against specs/db, resolves to expected.
		rel, err := filepath.Rel(absDB, expected)
		if err != nil {
			t.Fatalf("setup rel: %v", err)
		}
		writeSqlc(t, specs, "sql:\n  - gen:\n      go:\n          out: "+rel+"\n")
		if err := checkSqlcOutPath(specs, artifacts); err != nil {
			t.Errorf("expected relative match (nil error), got: %v", err)
		}
	})
}
