//ff:func feature=migration type=test control=sequence
//ff:what io_helpers_unit_test — LoadSnapshot/WriteSnapshot/NextSequenceNumber/LoadDataMigrationSQL/listSQLFiles/loadPrevSnapshot 단위 테스트
package migration

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadPrevSnapshot_LoadError(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "bad.sql")
	// valid header prefix but wrong hash -> LoadSnapshot returns error
	if err := os.WriteFile(path, []byte(SnapshotHashHeaderPrefix+"deadbeef\nbody\n"), 0644); err != nil {
		t.Fatalf("write: %v", err)
	}
	_, _, diags := loadPrevSnapshot(path, dir)
	if len(diags) != 1 || !strings.Contains(diags[0].Message, "MIG-006") {
		t.Errorf("expected one MIG-006 load-failed diag, got %v", diags)
	}
}
