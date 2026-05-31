//ff:func feature=migration type=test control=sequence
//ff:what io_helpers_unit_test — LoadSnapshot/WriteSnapshot/NextSequenceNumber/LoadDataMigrationSQL/listSQLFiles/loadPrevSnapshot 단위 테스트
package migration

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestWriteAndLoadSnapshot_RoundTrip(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "sub", ".latest_schema.sql")
	s := sampleSchema()

	if err := WriteSnapshot(path, s, "v1.2.3", time.Unix(0, 0)); err != nil {
		t.Fatalf("WriteSnapshot error: %v", err)
	}
	// file written + header present
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read written snapshot: %v", err)
	}
	if !strings.HasPrefix(string(data), SnapshotHashHeaderPrefix) {
		t.Errorf("snapshot missing hash header: %q", string(data)[:40])
	}

	loaded, err := LoadSnapshot(path)
	if err != nil {
		t.Fatalf("LoadSnapshot error: %v", err)
	}
	if loaded == nil || loaded.Tables["users"] == nil {
		t.Fatalf("loaded schema missing users table: %+v", loaded)
	}
	if len(loaded.Tables["users"].Columns) != 2 {
		t.Errorf("round-trip lost columns: %+v", loaded.Tables["users"].Columns)
	}
}
