//ff:func feature=migration type=test control=sequence
//ff:what TestAlterColumnNullable_Description — To 값에 따라 drop/set 문구 분기
package migration

import (
	"testing"
)

func TestAlterColumnNullable_Description(t *testing.T) {
	if got, want := (AlterColumnNullable{Table: "t", Column: "c", To: true}).Description(), "drop NOT NULL on t.c"; got != want {
		t.Errorf("To=true: %q, want %q", got, want)
	}
	if got, want := (AlterColumnNullable{Table: "t", Column: "c", To: false}).Description(), "set NOT NULL on t.c"; got != want {
		t.Errorf("To=false: %q, want %q", got, want)
	}
}
