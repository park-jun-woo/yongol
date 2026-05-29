//ff:func feature=migration type=test control=sequence
//ff:what TestAlterColumnType_Description — from→to 타입 표기 확인
package migration

import "testing"

func TestAlterColumnType_Description(t *testing.T) {
	op := AlterColumnType{Table: "t", Column: "c", From: CanonicalType{Base: "TEXT"}, To: CanonicalType{Base: "INTEGER"}}
	if got, want := op.Description(), "alter t.c type TEXT→INTEGER"; got != want {
		t.Errorf("Description() = %q, want %q", got, want)
	}
}
