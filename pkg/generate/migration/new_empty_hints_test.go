//ff:func feature=migration type=test control=sequence
//ff:what helpers_lookup_unit_test — checkMap/columnMap/fkMap/indexMap/rename/setFKAction/newEmptyHints/NewSchema/collectTypeTokens 단위 테스트
package migration

import (
	"testing"
)

func TestNewEmptyHints(t *testing.T) {
	h := newEmptyHints()
	if h.Casts == nil || h.Backfills == nil || h.DataMigrations == nil || h.AllowDestructive == nil {
		t.Fatalf("newEmptyHints left a nil map: %+v", h)
	}
	if len(h.RenameTables) != 0 || len(h.RenameColumns) != 0 {
		t.Errorf("rename slices should be empty initially")
	}
}
