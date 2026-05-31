//ff:func feature=migration type=test control=sequence
//ff:what io_helpers_unit_test — LoadSnapshot/WriteSnapshot/NextSequenceNumber/LoadDataMigrationSQL/listSQLFiles/loadPrevSnapshot 단위 테스트
package migration

import (
	"path/filepath"
	"testing"
)

func TestLoadSnapshot_Missing(t *testing.T) {
	s, err := LoadSnapshot(filepath.Join(t.TempDir(), "nope.sql"))
	if err != nil {
		t.Errorf("missing file should yield (nil, nil), got err %v", err)
	}
	if s != nil {
		t.Errorf("missing file should yield nil schema, got %+v", s)
	}
}
