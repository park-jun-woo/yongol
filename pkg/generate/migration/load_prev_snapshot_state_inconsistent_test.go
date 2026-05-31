//ff:func feature=migration type=test control=sequence
//ff:what io_helpers_unit_test — LoadSnapshot/WriteSnapshot/NextSequenceNumber/LoadDataMigrationSQL/listSQLFiles/loadPrevSnapshot 단위 테스트
package migration

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadPrevSnapshot_StateInconsistent(t *testing.T) {
	dir := t.TempDir()
	migDir := filepath.Join(dir, MigrationsSubdir)
	if err := os.MkdirAll(migDir, 0755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	if err := os.WriteFile(filepath.Join(migDir, "0001_x.up.sql"), []byte("x"), 0644); err != nil {
		t.Fatalf("write: %v", err)
	}
	_, mode, diags := loadPrevSnapshot(filepath.Join(dir, "absent.sql"), dir)
	if mode != ModeInitial {
		t.Errorf("mode = %q, want initial", mode)
	}
	if len(diags) != 1 || !strings.Contains(diags[0].Message, "MIG-006") {
		t.Errorf("expected one MIG-006 diag, got %v", diags)
	}
}
