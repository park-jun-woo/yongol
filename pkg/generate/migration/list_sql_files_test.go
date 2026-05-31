//ff:func feature=migration type=test control=iteration dimension=1
//ff:what io_helpers_unit_test — LoadSnapshot/WriteSnapshot/NextSequenceNumber/LoadDataMigrationSQL/listSQLFiles/loadPrevSnapshot 단위 테스트
package migration

import (
	"os"
	"path/filepath"
	"testing"
)

func TestListSQLFiles(t *testing.T) {
	// missing dir -> (nil, nil)
	if files, err := listSQLFiles(filepath.Join(t.TempDir(), "nope"), nil); err != nil || files != nil {
		t.Errorf("missing dir -> (%v,%v), want (nil,nil)", files, err)
	}

	dir := t.TempDir()
	for _, name := range []string{"b.sql", "a.sql", "skip.sql", "notsql.txt"} {
		if err := os.WriteFile(filepath.Join(dir, name), []byte("x"), 0644); err != nil {
			t.Fatalf("write %s: %v", name, err)
		}
	}
	files, err := listSQLFiles(dir, []string{"skip.sql"})
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if len(files) != 2 {
		t.Fatalf("got %d files, want 2: %v", len(files), files)
	}
	// sorted: a.sql before b.sql
	if filepath.Base(files[0]) != "a.sql" || filepath.Base(files[1]) != "b.sql" {
		t.Errorf("files not sorted/filtered correctly: %v", files)
	}
}
