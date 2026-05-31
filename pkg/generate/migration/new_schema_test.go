//ff:func feature=migration type=test control=sequence
//ff:what helpers_lookup_unit_test — checkMap/columnMap/fkMap/indexMap/rename/setFKAction/newEmptyHints/NewSchema/collectTypeTokens 단위 테스트
package migration

import (
	"testing"
)

func TestNewSchema(t *testing.T) {
	s := NewSchema()
	if s == nil || s.Tables == nil {
		t.Fatalf("NewSchema returned nil or nil Tables: %+v", s)
	}
	if len(s.Tables) != 0 {
		t.Errorf("Tables should be empty, got %d", len(s.Tables))
	}
}
