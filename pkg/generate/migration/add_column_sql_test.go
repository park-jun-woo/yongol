//ff:func feature=migration type=test control=sequence
//ff:what TestAddColumn_SQL — ADD COLUMN 문장이 렌더된 컬럼 clause 를 포함
package migration

import "testing"

func TestAddColumn_SQL(t *testing.T) {
	op := AddColumn{Table: "users", Column: &Column{Name: "age", Type: CanonicalType{Base: "INTEGER"}, Nullable: true}}
	want := "ALTER TABLE users ADD COLUMN age INTEGER;"
	if got := op.SQL(); got != want {
		t.Errorf("SQL() = %q, want %q", got, want)
	}
}
