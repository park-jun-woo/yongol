//ff:func feature=migration type=test control=sequence
//ff:what helpers_lookup_unit_test — checkMap/columnMap/fkMap/indexMap/rename/setFKAction/newEmptyHints/NewSchema/collectTypeTokens 단위 테스트
package migration

import (
	"testing"
)

func TestRenamedTableName(t *testing.T) {
	rules := []RenameTableHint{{From: "old_t", To: "new_t"}}
	if got := renamedTableName("old_t", rules); got != "new_t" {
		t.Errorf("match -> %q, want new_t", got)
	}
	if got := renamedTableName("other", rules); got != "other" {
		t.Errorf("no match -> %q, want other", got)
	}
	if got := renamedTableName("x", nil); got != "x" {
		t.Errorf("nil rules -> %q, want x", got)
	}
}
