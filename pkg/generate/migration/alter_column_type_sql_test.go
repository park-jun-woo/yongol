//ff:func feature=migration type=test control=sequence
//ff:what TestAlterColumnType_SQL — USING 명시 / 자동 ::cast 분기
package migration

import (
	"testing"
)

func TestAlterColumnType_SQL(t *testing.T) {
	t.Run("explicit using", func(t *testing.T) {
		op := AlterColumnType{Table: "t", Column: "c", To: CanonicalType{Base: "INTEGER"}, Using: "c::int"}
		want := "ALTER TABLE t ALTER COLUMN c TYPE INTEGER USING c::int;"
		if got := op.SQL(); got != want {
			t.Errorf("SQL() = %q, want %q", got, want)
		}
	})
	t.Run("auto cast", func(t *testing.T) {
		op := AlterColumnType{Table: "t", Column: "c", To: CanonicalType{Base: "INTEGER"}}
		want := "ALTER TABLE t ALTER COLUMN c TYPE INTEGER USING c::integer;"
		if got := op.SQL(); got != want {
			t.Errorf("SQL() = %q, want %q", got, want)
		}
	})
}
