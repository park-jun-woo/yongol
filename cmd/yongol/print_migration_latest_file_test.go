//ff:func feature=cli type=test control=sequence
//ff:what printMigrationLatestFile test — 최근 마이그레이션 파일명 출력 검증

package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/migration"
)

func TestPrintMigrationLatestFile(t *testing.T) {
	t.Run("Normal", func(t *testing.T) {
		arts := t.TempDir()
		migDir := filepath.Join(arts, migration.MigrationsSubdir)
		if err := os.MkdirAll(migDir, 0o755); err != nil {
			t.Fatal(err)
		}
		mustWriteEmpty(t, filepath.Join(migDir, "0001_init.sql"))
		mustWriteEmpty(t, filepath.Join(migDir, "0002_add_users.sql"))

		var buf bytes.Buffer
		printMigrationLatestFile(&buf, arts)
		out := buf.String()
		if !strings.Contains(out, "0002_add_users.sql") {
			t.Errorf("expected latest file, got: %q", out)
		}
	})

	t.Run("NoSQLFiles", func(t *testing.T) {
		arts := t.TempDir()
		migDir := filepath.Join(arts, migration.MigrationsSubdir)
		if err := os.MkdirAll(migDir, 0o755); err != nil {
			t.Fatal(err)
		}
		mustWriteEmpty(t, filepath.Join(migDir, "readme.md"))

		var buf bytes.Buffer
		printMigrationLatestFile(&buf, arts)
		if buf.Len() != 0 {
			t.Errorf("expected no output when no .sql files, got: %q", buf.String())
		}
	})

	t.Run("MissingDir", func(t *testing.T) {
		var buf bytes.Buffer
		printMigrationLatestFile(&buf, "/tmp/no-such-yongol-dir")
		if buf.Len() != 0 {
			t.Errorf("expected no output for missing dir, got: %q", buf.String())
		}
	})

	t.Run("EmptyArtsDir", func(t *testing.T) {
		var buf bytes.Buffer
		printMigrationLatestFile(&buf, "")
		if buf.Len() != 0 {
			t.Errorf("expected no output for empty artsDir, got: %q", buf.String())
		}
	})
}
