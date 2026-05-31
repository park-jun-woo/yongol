//ff:func feature=migration type=test control=sequence
//ff:what helpers_lookup_unit_test — checkMap/columnMap/fkMap/indexMap/rename/setFKAction/newEmptyHints/NewSchema/collectTypeTokens 단위 테스트
package migration

import (
	"testing"
)

func TestColumnMap(t *testing.T) {
	c := &Column{Name: "id"}
	m := columnMap([]*Column{c})
	if m["id"] != c {
		t.Errorf("columnMap did not key by name")
	}
	if got := columnMap(nil); len(got) != 0 {
		t.Errorf("nil -> len %d, want 0", len(got))
	}
}
