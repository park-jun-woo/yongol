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

func TestLoadSnapshot_Drift(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, ".latest_schema.sql")
	if err := WriteSnapshot(path, sampleSchema(), "v1", time.Unix(0, 0)); err != nil {
		t.Fatalf("WriteSnapshot: %v", err)
	}
	// tamper with the body so the stored hash no longer matches
	data, _ := os.ReadFile(path)
	tampered := string(data) + "\n-- injected drift\n"
	if err := os.WriteFile(path, []byte(tampered), 0644); err != nil {
		t.Fatalf("rewrite: %v", err)
	}
	if _, err := LoadSnapshot(path); err == nil || !strings.Contains(err.Error(), "drift") {
		t.Errorf("expected drift hash-mismatch error, got %v", err)
	}
}
