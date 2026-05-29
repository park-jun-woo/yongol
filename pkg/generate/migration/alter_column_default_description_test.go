//ff:func feature=migration type=test control=sequence
//ff:what TestAlterColumnDefault_Description — 헤더용 표기 확인
package migration

import "testing"

func TestAlterColumnDefault_Description(t *testing.T) {
	op := AlterColumnDefault{Table: "t", Column: "c", To: "0"}
	if got, want := op.Description(), "alter default t.c → 0"; got != want {
		t.Errorf("Description() = %q, want %q", got, want)
	}
}
