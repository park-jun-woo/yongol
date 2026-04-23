//ff:func feature=migration type=test-helper control=sequence
//ff:what assertBasicUsersTable — TestBuildASTFromSQL_Basic 의 users 테이블 필드 검증
package migration

import "testing"

func assertBasicUsersTable(t *testing.T, tbl *Table) {
	t.Helper()
	if len(tbl.Columns) != 5 {
		t.Fatalf("expected 5 columns, got %d", len(tbl.Columns))
	}
	if tbl.Columns[0].Name != "id" || tbl.Columns[0].Type.Base != "BIGINT" {
		t.Errorf("id column: %+v", tbl.Columns[0])
	}
	if len(tbl.PrimaryKey) != 1 || tbl.PrimaryKey[0] != "id" {
		t.Errorf("PK: %+v", tbl.PrimaryKey)
	}
	email := tbl.Columns[1]
	if email.Type.Base != "VARCHAR" || email.Type.Length != 255 {
		t.Errorf("email type: %+v", email.Type)
	}
	if email.Nullable {
		t.Errorf("email should be NOT NULL")
	}
	if !tableHasUniqueIndexOnEmail(tbl) {
		t.Errorf("unique index on email not generated: %+v", tbl.Indexes)
	}
	if tbl.Columns[4].Default != "CURRENT_TIMESTAMP" {
		t.Errorf("created_at default: %q", tbl.Columns[4].Default)
	}
	if tbl.Columns[3].Default != "TRUE" {
		t.Errorf("active default: %q", tbl.Columns[3].Default)
	}
}
