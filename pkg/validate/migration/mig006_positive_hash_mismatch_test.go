//ff:func feature=validate type=test control=sequence topic=migration-snapshot
//ff:what MIG-006 positive — 해시 불일치 시 ERROR

package migration

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	migration "github.com/park-jun-woo/yongol/pkg/generate/migration"
)

func TestMIG006_Positive_HashMismatch(t *testing.T) {
	tmp := t.TempDir()
	baselineDir := filepath.Join(tmp, migration.BaselineSubdir)
	if err := os.MkdirAll(baselineDir, 0755); err != nil {
		t.Fatal(err)
	}
	content := migration.SnapshotHashHeaderPrefix + "deadbeef\nCREATE TABLE t (id INTEGER);\n"
	if err := os.WriteFile(filepath.Join(baselineDir, migration.SnapshotFileName), []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
	diags := Mig006SnapshotDrift(tmp)
	if len(diags) != 1 {
		t.Fatalf("expected 1 MIG-006 diag, got %+v", diags)
	}
	if diags[0].Level != diagnostic.LevelError {
		t.Errorf("expected ERROR, got %v", diags[0].Level)
	}
}
