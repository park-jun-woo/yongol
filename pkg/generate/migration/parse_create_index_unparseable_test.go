//ff:func feature=migration type=test control=sequence
//ff:what parse_statements_unit_test — parseCreateTable/parseCreateIndex/parseAlterTable + 하위 parse 헬퍼 통합 단위 테스트
package migration

import (
	"testing"
)

func TestParseCreateIndex_Unparseable(t *testing.T) {
	s := NewSchema()
	if err := parseCreateIndex(s, "SELECT 1"); err != nil {
		t.Errorf("non-index statement should be skipped silently, got %v", err)
	}
}
