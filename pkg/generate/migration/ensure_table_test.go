//ff:func feature=migration type=test control=sequence
//ff:what parse_statements_unit_test — parseCreateTable/parseCreateIndex/parseAlterTable + 하위 parse 헬퍼 통합 단위 테스트
package migration

import (
	"testing"
)

func TestEnsureTable(t *testing.T) {
	s := NewSchema()
	a := ensureTable(s, "users")
	if a == nil || a.Name != "users" {
		t.Fatalf("ensureTable created wrong table: %+v", a)
	}
	b := ensureTable(s, "users")
	if a != b {
		t.Errorf("ensureTable should return the same instance for an existing name")
	}
	if len(s.Tables) != 1 {
		t.Errorf("schema should have exactly 1 table, got %d", len(s.Tables))
	}
}
