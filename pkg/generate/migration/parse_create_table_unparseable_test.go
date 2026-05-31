//ff:func feature=migration type=test control=sequence
//ff:what parse_statements_unit_test — parseCreateTable/parseCreateIndex/parseAlterTable + 하위 parse 헬퍼 통합 단위 테스트
package migration

import (
	"testing"
)

func TestParseCreateTable_Unparseable(t *testing.T) {
	s := NewSchema()
	if err := parseCreateTable(s, "CREATE INDEX foo ON bar (x)"); err == nil {
		t.Errorf("expected error for non-CREATE-TABLE statement")
	}
}
