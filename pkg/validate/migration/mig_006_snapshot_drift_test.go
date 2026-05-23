//ff:func feature=validate type=test control=sequence topic=migration-snapshot
//ff:what Mig006SnapshotDrift — 파일 부재/헤더 누락/해시 불일치/정상 검증

package migration

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"path/filepath"
	"strings"
	"testing"

	gmig "github.com/park-jun-woo/yongol/pkg/generate/migration"
)

func TestMig006SnapshotDrift(t *testing.T) {
	t.Run("absent snapshot returns nil", func(t *testing.T) {
		tmp := t.TempDir()
		diags := Mig006SnapshotDrift(tmp)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("no hash header raises error", func(t *testing.T) {
		tmp := t.TempDir()
		dir := filepath.Join(tmp, gmig.BaselineSubdir)
		os.MkdirAll(dir, 0755)
		os.WriteFile(filepath.Join(dir, gmig.SnapshotFileName), []byte("CREATE TABLE t (id INT);\n"), 0644)

		diags := Mig006SnapshotDrift(tmp)
		if len(diags) != 1 {
			t.Fatalf("expected 1 diagnostic, got %d", len(diags))
		}
		if !strings.Contains(diags[0].Message, "MIG-006") {
			t.Errorf("Message missing MIG-006: %s", diags[0].Message)
		}
		if !strings.Contains(diags[0].Message, "no YONGOL_SCHEMA_HASH") {
			t.Errorf("Message missing detail: %s", diags[0].Message)
		}
	})

	t.Run("hash mismatch raises error", func(t *testing.T) {
		tmp := t.TempDir()
		dir := filepath.Join(tmp, gmig.BaselineSubdir)
		os.MkdirAll(dir, 0755)
		content := gmig.SnapshotHashHeaderPrefix + "deadbeef\nCREATE TABLE t (id INT);\n"
		os.WriteFile(filepath.Join(dir, gmig.SnapshotFileName), []byte(content), 0644)

		diags := Mig006SnapshotDrift(tmp)
		if len(diags) != 1 {
			t.Fatalf("expected 1 diagnostic, got %d", len(diags))
		}
		if !strings.Contains(diags[0].Message, "drift") {
			t.Errorf("Message missing drift: %s", diags[0].Message)
		}
	})

	t.Run("correct hash returns nil", func(t *testing.T) {
		tmp := t.TempDir()
		dir := filepath.Join(tmp, gmig.BaselineSubdir)
		os.MkdirAll(dir, 0755)

		body := "CREATE TABLE t (id INTEGER);\n"
		sum := sha256.Sum256([]byte(body))
		hashStr := hex.EncodeToString(sum[:])
		content := gmig.SnapshotHashHeaderPrefix + hashStr + "\n" + body
		os.WriteFile(filepath.Join(dir, gmig.SnapshotFileName), []byte(content), 0644)

		diags := Mig006SnapshotDrift(tmp)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})
}
