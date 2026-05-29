//ff:func feature=migration type=test control=sequence
//ff:what TestDropColumn_Description — table.column 설명 확인
package migration

import "testing"

func TestDropColumn_Description(t *testing.T) {
	op := DropColumn{Table: "users", Column: "age"}
	if got, want := op.Description(), "drop column users.age"; got != want {
		t.Errorf("Description() = %q, want %q", got, want)
	}
}
