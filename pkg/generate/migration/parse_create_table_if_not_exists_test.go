//ff:func feature=migration type=test control=sequence
//ff:what parse_statements_unit_test — parseCreateTable/parseCreateIndex/parseAlterTable + 하위 parse 헬퍼 통합 단위 테스트
package migration

import (
	"testing"
)

func TestParseCreateTable_IfNotExists(t *testing.T) {
	s := NewSchema()
	if err := parseCreateTable(s, "CREATE TABLE IF NOT EXISTS t (id INTEGER NOT NULL)"); err != nil {
		t.Fatalf("error: %v", err)
	}
	if s.Tables["t"] == nil {
		t.Errorf("IF NOT EXISTS table not parsed")
	}
}
