//ff:func feature=migration type=test control=sequence
//ff:what helpers_lookup_unit_test — checkMap/columnMap/fkMap/indexMap/rename/setFKAction/newEmptyHints/NewSchema/collectTypeTokens 단위 테스트
package migration

import (
	"testing"
)

func TestCheckMap(t *testing.T) {
	a := &CheckConstraint{Name: "chk_a", Expression: "x > 0"}
	b := &CheckConstraint{Name: "chk_b", Expression: "y < 1"}
	m := checkMap([]*CheckConstraint{a, b})
	if len(m) != 2 {
		t.Fatalf("len = %d, want 2", len(m))
	}
	if m["chk_a"] != a || m["chk_b"] != b {
		t.Errorf("map did not key by name correctly")
	}
	if got := checkMap(nil); len(got) != 0 {
		t.Errorf("nil input -> len %d, want 0", len(got))
	}
}
