//ff:func feature=migration type=test control=sequence
//ff:what io_helpers_unit_test — LoadSnapshot/WriteSnapshot/NextSequenceNumber/LoadDataMigrationSQL/listSQLFiles/loadPrevSnapshot 단위 테스트
package migration

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadDataMigrationSQL(t *testing.T) {
	// nil hints / empty -> nil, nil
	if out, missing := LoadDataMigrationSQL("/tmp", nil); out != nil || missing != nil {
		t.Errorf("nil hints -> (%v,%v), want (nil,nil)", out, missing)
	}

	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "mig_users.sql"), []byte("UPDATE users SET x=1;"), 0644); err != nil {
		t.Fatalf("write: %v", err)
	}
	hints := &Hints{DataMigrations: map[string]string{
		"users":  "mig_users.sql", // relative -> resolved against specsDir
		"orders": "missing.sql",   // missing -> reported
	}}
	out, missing := LoadDataMigrationSQL(dir, hints)
	if out["users"] != "UPDATE users SET x=1;" {
		t.Errorf("users sidecar not loaded: %q", out["users"])
	}
	if len(missing) != 1 || missing[0] != "missing.sql" {
		t.Errorf("missing = %v, want [missing.sql]", missing)
	}
}
