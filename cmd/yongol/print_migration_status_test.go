//ff:func feature=cli type=test control=sequence
//ff:what printMigrationStatus test — Migration Status 섹션 출력 검증

package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/migration"
)

func TestPrintMigrationStatus(t *testing.T) {
	t.Run("EmptyProject", func(t *testing.T) {
		specsDir := t.TempDir()
		artsDir := t.TempDir()
		// Create the expected ddl subdirectory (even though empty).
		if err := os.MkdirAll(filepath.Join(specsDir, migration.DDLSubdir), 0o755); err != nil {
			t.Fatal(err)
		}

		var buf bytes.Buffer
		printMigrationStatus(&buf, specsDir, artsDir)
		out := buf.String()

		if !strings.Contains(out, "Migration Status") {
			t.Errorf("expected header, got: %q", out)
		}
		if !strings.Contains(out, "absent") {
			t.Errorf("expected absent snapshot, got: %q", out)
		}
		if !strings.Contains(out, "pending: 0 change(s)") {
			t.Errorf("expected 0 pending, got: %q", out)
		}
	})

	t.Run("DDLParseError", func(t *testing.T) {
		specsDir := t.TempDir()
		artsDir := t.TempDir()
		// DDL dir points to a file (not directory) so BuildASTFromDir fails.
		ddlParent := filepath.Join(specsDir, migration.DDLSubdir)
		if err := os.MkdirAll(filepath.Dir(ddlParent), 0o755); err != nil {
			t.Fatal(err)
		}
		// Write a file where the DDL dir should be.
		if err := os.WriteFile(ddlParent, []byte("x"), 0o644); err != nil {
			t.Fatal(err)
		}

		var buf bytes.Buffer
		printMigrationStatus(&buf, specsDir, artsDir)
		out := buf.String()

		if !strings.Contains(out, "parse error") {
			t.Errorf("expected parse error, got: %q", out)
		}
		// Should not print latest after parse error.
		if strings.Contains(out, "latest:") {
			t.Errorf("should not print latest after parse error, got: %q", out)
		}
	})

	t.Run("SnapshotReadError", func(t *testing.T) {
		specsDir := t.TempDir()
		artsDir := t.TempDir()
		// Create the snapshot path as a directory to trigger read error.
		snapDir := filepath.Join(artsDir, migration.BaselineSubdir, migration.SnapshotFileName)
		if err := os.MkdirAll(snapDir, 0o755); err != nil {
			t.Fatal(err)
		}

		var buf bytes.Buffer
		printMigrationStatus(&buf, specsDir, artsDir)
		out := buf.String()

		if !strings.Contains(out, "ERROR") {
			t.Errorf("expected snapshot error, got: %q", out)
		}
		// Should not print pending/latest after error.
		if strings.Contains(out, "pending:") {
			t.Errorf("should not print pending after snapshot error, got: %q", out)
		}
	})
}
