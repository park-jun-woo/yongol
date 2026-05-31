//ff:func feature=migration type=test control=iteration dimension=1
//ff:what io_helpers_unit_test — LoadSnapshot/WriteSnapshot/NextSequenceNumber/LoadDataMigrationSQL/listSQLFiles/loadPrevSnapshot 단위 테스트
package migration

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNextSequenceNumber(t *testing.T) {
	// missing dir -> 1
	if n, err := NextSequenceNumber(filepath.Join(t.TempDir(), "nope")); err != nil || n != 1 {
		t.Errorf("missing dir -> (%d,%v), want (1,nil)", n, err)
	}

	dir := t.TempDir()
	// empty dir -> 1
	if n, err := NextSequenceNumber(dir); err != nil || n != 1 {
		t.Errorf("empty dir -> (%d,%v), want (1,nil)", n, err)
	}

	// with files: max 0003 .up.sql, plus noise -> 4
	for _, name := range []string{
		"0001_initial.up.sql",
		"0001_initial.down.sql", // down ignored
		"0003_add.up.sql",
		"readme.txt",      // non-sql ignored
		"noprefix.up.sql", // no leading number index 0 -> i<=0 skip
		"02x_bad.up.sql",  // non-numeric prefix -> skip
	} {
		if err := os.WriteFile(filepath.Join(dir, name), []byte("x"), 0644); err != nil {
			t.Fatalf("write %s: %v", name, err)
		}
	}
	n, err := NextSequenceNumber(dir)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if n != 4 {
		t.Errorf("NextSequenceNumber = %d, want 4", n)
	}
}
