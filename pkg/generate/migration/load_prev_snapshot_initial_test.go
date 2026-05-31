//ff:func feature=migration type=test control=sequence
//ff:what io_helpers_unit_test — LoadSnapshot/WriteSnapshot/NextSequenceNumber/LoadDataMigrationSQL/listSQLFiles/loadPrevSnapshot 단위 테스트
package migration

import (
	"path/filepath"
	"testing"
)

func TestLoadPrevSnapshot_Initial(t *testing.T) {
	dir := t.TempDir()
	prev, mode, diags := loadPrevSnapshot(filepath.Join(dir, "nope.sql"), dir)
	if mode != ModeInitial {
		t.Errorf("missing snapshot -> mode %q, want initial", mode)
	}
	if prev == nil || prev.Tables == nil {
		t.Errorf("initial prev should be empty non-nil schema")
	}
	if len(diags) != 0 {
		t.Errorf("clean initial should produce no diags, got %v", diags)
	}
}
