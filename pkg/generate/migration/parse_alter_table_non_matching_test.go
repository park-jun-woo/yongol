//ff:func feature=migration type=test control=sequence
//ff:what parse_statements_unit_test — parseCreateTable/parseCreateIndex/parseAlterTable + 하위 parse 헬퍼 통합 단위 테스트
package migration

import (
	"testing"
)

func TestParseAlterTable_NonMatching(t *testing.T) {
	s := NewSchema()
	if err := parseAlterTable(s, "ALTER TABLE t ADD COLUMN x INTEGER"); err != nil {
		t.Errorf("non-FK ALTER should be skipped silently, got %v", err)
	}
	if len(s.Tables) != 0 {
		t.Errorf("non-matching ALTER should not create tables")
	}
}
