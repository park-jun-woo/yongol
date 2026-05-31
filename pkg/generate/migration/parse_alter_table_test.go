//ff:func feature=migration type=test control=sequence
//ff:what parse_statements_unit_test — parseCreateTable/parseCreateIndex/parseAlterTable + 하위 parse 헬퍼 통합 단위 테스트
package migration

import (
	"testing"
)

func TestParseAlterTable(t *testing.T) {
	s := NewSchema()
	stmt := "ALTER TABLE orders ADD CONSTRAINT fk_u FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE SET NULL ON UPDATE CASCADE"
	if err := parseAlterTable(s, stmt); err != nil {
		t.Fatalf("error: %v", err)
	}
	tbl := s.Tables["orders"]
	if len(tbl.ForeignKeys) != 1 {
		t.Fatalf("expected 1 FK, got %d", len(tbl.ForeignKeys))
	}
	fk := tbl.ForeignKeys[0]
	if fk.Name != "fk_u" || fk.RefTable != "users" || fk.OnDelete != "SET NULL" || fk.OnUpdate != "CASCADE" {
		t.Errorf("FK parsed wrong: %+v", fk)
	}
	if len(fk.Columns) != 1 || fk.Columns[0] != "user_id" {
		t.Errorf("FK columns wrong: %v", fk.Columns)
	}
}
