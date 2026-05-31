//ff:func feature=migration type=test control=sequence
//ff:what helpers_lookup_unit_test — checkMap/columnMap/fkMap/indexMap/rename/setFKAction/newEmptyHints/NewSchema/collectTypeTokens 단위 테스트
package migration

import (
	"testing"
)

func TestIndexMap(t *testing.T) {
	ix := &Index{Name: "ix1"}
	m := indexMap([]*Index{ix})
	if m["ix1"] != ix {
		t.Errorf("indexMap did not key by name")
	}
	if got := indexMap(nil); len(got) != 0 {
		t.Errorf("nil -> len %d, want 0", len(got))
	}
}
