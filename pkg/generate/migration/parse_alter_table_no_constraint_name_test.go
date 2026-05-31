//ff:func feature=migration type=test control=sequence
//ff:what parse_statements_unit_test — parseCreateTable/parseCreateIndex/parseAlterTable + 하위 parse 헬퍼 통합 단위 테스트
package migration

import (
	"testing"
)

func TestParseAlterTable_NoConstraintName(t *testing.T) {
	s := NewSchema()
	if err := parseAlterTable(s, "ALTER TABLE orders ADD FOREIGN KEY (user_id) REFERENCES users (id)"); err != nil {
		t.Fatalf("error: %v", err)
	}
	fk := s.Tables["orders"].ForeignKeys[0]
	if fk.Name == "" {
		t.Errorf("FK name should be auto-generated when omitted")
	}
}
