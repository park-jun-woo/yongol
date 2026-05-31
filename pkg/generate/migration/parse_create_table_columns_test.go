//ff:func feature=migration type=test control=sequence
//ff:what parse_statements_unit_test — parseCreateTable/parseCreateIndex/parseAlterTable + 하위 parse 헬퍼 통합 단위 테스트
package migration

import (
	"testing"
)

func TestParseCreateTable_Columns(t *testing.T) {
	s := NewSchema()
	stmt := `CREATE TABLE users (
		id BIGINT NOT NULL,
		email VARCHAR(255) NOT NULL,
		status TEXT DEFAULT 'active',
		created_at TIMESTAMPTZ,
		PRIMARY KEY (id)
	)`
	if err := parseCreateTable(s, stmt); err != nil {
		t.Fatalf("parseCreateTable error: %v", err)
	}
	tbl := s.Tables["users"]
	if tbl == nil {
		t.Fatalf("users table not registered")
	}
	if len(tbl.Columns) != 4 {
		t.Fatalf("got %d columns, want 4: %+v", len(tbl.Columns), tbl.Columns)
	}
	id := tbl.Columns[0]
	if id.Name != "id" || id.Nullable {
		t.Errorf("id column wrong: %+v", id)
	}
	email := tbl.Columns[1]
	if email.Name != "email" || email.Type.Base != "VARCHAR" || email.Type.Length != 255 {
		t.Errorf("email column wrong: %+v (type %+v)", email, email.Type)
	}
	status := tbl.Columns[2]
	if status.Default != "'active'" {
		t.Errorf("status default = %q, want 'active'", status.Default)
	}
	createdAt := tbl.Columns[3]
	if !createdAt.Nullable {
		t.Errorf("created_at should be nullable")
	}
	if len(tbl.PrimaryKey) != 1 || tbl.PrimaryKey[0] != "id" {
		t.Errorf("PK = %v, want [id]", tbl.PrimaryKey)
	}
}
