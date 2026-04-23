//ff:func feature=validate type=test control=sequence topic=migration-snapshot
//ff:what MIG-006 negative — 해시가 정상이면 진단 없음

package migration

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"path/filepath"
	"testing"

	migration "github.com/park-jun-woo/yongol/pkg/generate/migration"
)

func TestMIG006_Negative_ValidHash(t *testing.T) {
	tmp := t.TempDir()
	dbDir := filepath.Join(tmp, "db")
	if err := os.MkdirAll(dbDir, 0755); err != nil {
		t.Fatal(err)
	}
	body := "CREATE TABLE t (id INTEGER);\n"
	sum := sha256.Sum256([]byte(body))
	content := migration.SnapshotHashHeaderPrefix + hex.EncodeToString(sum[:]) + "\n" + body
	if err := os.WriteFile(filepath.Join(dbDir, migration.SnapshotFileName), []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
	diags := Mig006SnapshotDrift(tmp)
	if len(diags) != 0 {
		t.Errorf("valid hash should not diag, got %+v", diags)
	}
}
