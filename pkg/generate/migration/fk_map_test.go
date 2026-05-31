//ff:func feature=migration type=test control=sequence
//ff:what helpers_lookup_unit_test — checkMap/columnMap/fkMap/indexMap/rename/setFKAction/newEmptyHints/NewSchema/collectTypeTokens 단위 테스트
package migration

import (
	"testing"
)

func TestFKMap(t *testing.T) {
	fk := &ForeignKey{Name: "fk1"}
	m := fkMap([]*ForeignKey{fk})
	if m["fk1"] != fk {
		t.Errorf("fkMap did not key by name")
	}
	if got := fkMap(nil); len(got) != 0 {
		t.Errorf("nil -> len %d, want 0", len(got))
	}
}
