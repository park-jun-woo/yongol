//ff:func feature=cli type=test control=sequence
//ff:what printMigrationSnapshotInfo test — snapshot 상태 출력 검증

package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/migration"
)

func TestPrintMigrationSnapshotInfo(t *testing.T) {
	t.Run("Absent", func(t *testing.T) {
		snapPath := filepath.Join(t.TempDir(), "nonexistent.sql")
		var buf bytes.Buffer
		ok := printMigrationSnapshotInfo(&buf, snapPath)
		if !ok {
			t.Fatal("expected ok=true for absent snapshot")
		}
		if !strings.Contains(buf.String(), "absent") {
			t.Errorf("expected absent message, got: %q", buf.String())
		}
	})

	t.Run("ValidSnapshot", func(t *testing.T) {
		body := "CREATE TABLE foo (id BIGINT);\n"
		sum := sha256.Sum256([]byte(body))
		hash := hex.EncodeToString(sum[:])
		content := fmt.Sprintf("%s%s\n%s", migration.SnapshotHashHeaderPrefix, hash, body)

		snapPath := filepath.Join(t.TempDir(), "snap.sql")
		if err := os.WriteFile(snapPath, []byte(content), 0o644); err != nil {
			t.Fatal(err)
		}

		var buf bytes.Buffer
		ok := printMigrationSnapshotInfo(&buf, snapPath)
		if !ok {
			t.Fatal("expected ok=true")
		}
		out := buf.String()
		if !strings.Contains(out, "ok") {
			t.Errorf("expected ok status, got: %q", out)
		}
		if !strings.Contains(out, hash[:8]) {
			t.Errorf("expected hash prefix, got: %q", out)
		}
	})

	t.Run("DriftedSnapshot", func(t *testing.T) {
		content := fmt.Sprintf("%s%s\n%s",
			migration.SnapshotHashHeaderPrefix, "badhash0", "CREATE TABLE bar;\n")

		snapPath := filepath.Join(t.TempDir(), "snap.sql")
		if err := os.WriteFile(snapPath, []byte(content), 0o644); err != nil {
			t.Fatal(err)
		}

		var buf bytes.Buffer
		ok := printMigrationSnapshotInfo(&buf, snapPath)
		if !ok {
			t.Fatal("expected ok=true")
		}
		if !strings.Contains(buf.String(), "DRIFT") {
			t.Errorf("expected DRIFT status, got: %q", buf.String())
		}
	})

	t.Run("ReadError", func(t *testing.T) {
		// Create a directory where a file is expected.
		dir := t.TempDir()
		snapPath := filepath.Join(dir, "dir_not_file")
		if err := os.MkdirAll(snapPath, 0o755); err != nil {
			t.Fatal(err)
		}

		var buf bytes.Buffer
		ok := printMigrationSnapshotInfo(&buf, snapPath)
		if ok {
			t.Error("expected ok=false for read error")
		}
		if !strings.Contains(buf.String(), "ERROR") {
			t.Errorf("expected ERROR message, got: %q", buf.String())
		}
	})
}
