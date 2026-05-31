//ff:func feature=migration type=test control=sequence
//ff:what parse_statements_unit_test — parseCreateTable/parseCreateIndex/parseAlterTable + 하위 parse 헬퍼 통합 단위 테스트
package migration

import (
	"testing"
)

func TestParseCreateTable_NamedConstraints(t *testing.T) {
	s := NewSchema()
	stmt := `CREATE TABLE t (
		a INTEGER NOT NULL,
		b INTEGER NOT NULL,
		CONSTRAINT pk_t PRIMARY KEY (a, b),
		CONSTRAINT uq_b UNIQUE (b),
		CONSTRAINT chk_a CHECK (a > 0)
	)`
	if err := parseCreateTable(s, stmt); err != nil {
		t.Fatalf("error: %v", err)
	}
	tbl := s.Tables["t"]
	if len(tbl.PrimaryKey) != 2 {
		t.Errorf("PK = %v, want [a b]", tbl.PrimaryKey)
	}
	if len(tbl.Indexes) != 1 || tbl.Indexes[0].Name != "uq_b" {
		t.Errorf("named unique index wrong: %+v", tbl.Indexes)
	}
	if len(tbl.Checks) != 1 || tbl.Checks[0].Name != "chk_a" {
		t.Errorf("named check wrong: %+v", tbl.Checks)
	}
}
