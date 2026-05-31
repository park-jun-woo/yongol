//ff:func feature=migration type=test control=sequence
//ff:what build_and_mig001_unit_test — buildAlterColumnNullable/Type + mig001CheckRename(Columns/Tables) + removeLegacyBaseline 단위 테스트
package migration

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRemoveLegacyBaseline(t *testing.T) {
	dir := t.TempDir()
	// no file -> no-op, no panic
	removeLegacyBaseline(dir)

	legacy := filepath.Join(dir, LegacySnapshotFileName)
	if err := os.WriteFile(legacy, []byte("stale"), 0644); err != nil {
		t.Fatalf("write: %v", err)
	}
	removeLegacyBaseline(dir)
	if _, err := os.Stat(legacy); !os.IsNotExist(err) {
		t.Errorf("legacy baseline should have been removed, stat err = %v", err)
	}
}
