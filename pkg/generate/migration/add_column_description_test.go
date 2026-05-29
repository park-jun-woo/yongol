//ff:func feature=migration type=test control=sequence
//ff:what TestAddColumn_Description — table.column 형식 설명 확인
package migration

import "testing"

func TestAddColumn_Description(t *testing.T) {
	op := AddColumn{Table: "users", Column: &Column{Name: "age"}}
	if got, want := op.Description(), "add column users.age"; got != want {
		t.Errorf("Description() = %q, want %q", got, want)
	}
}
