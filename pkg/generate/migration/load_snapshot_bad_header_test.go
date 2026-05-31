//ff:func feature=migration type=test control=sequence
//ff:what io_helpers_unit_test — LoadSnapshot/WriteSnapshot/NextSequenceNumber/LoadDataMigrationSQL/listSQLFiles/loadPrevSnapshot 단위 테스트
package migration

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadSnapshot_BadHeader(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "bad.sql")
	if err := os.WriteFile(path, []byte("no header prefix here\nbody\n"), 0644); err != nil {
		t.Fatalf("write: %v", err)
	}
	if _, err := LoadSnapshot(path); err == nil {
		t.Errorf("expected error for missing hash header prefix")
	}

	// single line, no newline -> "no header line"
	path2 := filepath.Join(dir, "oneline.sql")
	if err := os.WriteFile(path2, []byte("oneline"), 0644); err != nil {
		t.Fatalf("write: %v", err)
	}
	if _, err := LoadSnapshot(path2); err == nil {
		t.Errorf("expected error for snapshot with no header line")
	}
}
