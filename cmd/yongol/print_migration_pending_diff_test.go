//ff:func feature=cli type=test control=sequence
//ff:what printMigrationPendingDiff test — pending diff 요약 출력 검증

package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestPrintMigrationPendingDiff(t *testing.T) {
	t.Run("EmptyDDLDir", func(t *testing.T) {
		ddlDir := t.TempDir()
		snapshotPath := filepath.Join(t.TempDir(), "snap.sql")

		var buf bytes.Buffer
		ok := printMigrationPendingDiff(&buf, ddlDir, snapshotPath)
		if !ok {
			t.Fatal("expected ok=true")
		}
		if !strings.Contains(buf.String(), "pending: 0 change(s)") {
			t.Errorf("expected 0 changes, got: %q", buf.String())
		}
	})

	t.Run("WithOneTable", func(t *testing.T) {
		ddlDir := t.TempDir()
		sql := "CREATE TABLE users (\n  id BIGINT NOT NULL PRIMARY KEY\n);\n"
		if err := os.WriteFile(filepath.Join(ddlDir, "users.sql"), []byte(sql), 0o644); err != nil {
			t.Fatal(err)
		}
		snapshotPath := filepath.Join(t.TempDir(), "snap.sql")

		var buf bytes.Buffer
		ok := printMigrationPendingDiff(&buf, ddlDir, snapshotPath)
		if !ok {
			t.Fatal("expected ok=true")
		}
		out := buf.String()
		// Should detect at least 1 change (the new table).
		if strings.Contains(out, "pending: 0 change(s)") {
			t.Errorf("expected at least 1 change, got: %q", out)
		}
	})

	t.Run("ParseError", func(t *testing.T) {
		// Use a file path (not directory) as ddlDir to trigger read error.
		fakeFile := filepath.Join(t.TempDir(), "not_a_dir.sql")
		if err := os.WriteFile(fakeFile, []byte("x"), 0o644); err != nil {
			t.Fatal(err)
		}
		snapshotPath := filepath.Join(t.TempDir(), "snap.sql")

		var buf bytes.Buffer
		ok := printMigrationPendingDiff(&buf, fakeFile, snapshotPath)
		if ok {
			t.Error("expected ok=false for parse error")
		}
		if !strings.Contains(buf.String(), "parse error") {
			t.Errorf("expected parse error message, got: %q", buf.String())
		}
	})

	t.Run("MoreThan10Changes", func(t *testing.T) {
		ddlDir := t.TempDir()
		// Create 12 tables so diff produces >10 operations.
		for i := 0; i < 12; i++ {
			sql := fmt.Sprintf("CREATE TABLE t%02d (\n  id BIGINT NOT NULL PRIMARY KEY\n);\n", i)
			if err := os.WriteFile(filepath.Join(ddlDir, fmt.Sprintf("t%02d.sql", i)), []byte(sql), 0o644); err != nil {
				t.Fatal(err)
			}
		}
		snapshotPath := filepath.Join(t.TempDir(), "snap.sql")

		var buf bytes.Buffer
		ok := printMigrationPendingDiff(&buf, ddlDir, snapshotPath)
		if !ok {
			t.Fatal("expected ok=true")
		}
		out := buf.String()
		if !strings.Contains(out, "... and") {
			t.Errorf("expected truncation message, got: %q", out)
		}
	})
}
