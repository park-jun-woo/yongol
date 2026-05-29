//ff:func feature=migration type=test control=sequence
//ff:what TestCreateTable_Description — "create table <name>"
package migration

import "testing"

func TestCreateTable_Description(t *testing.T) {
	op := CreateTable{Table: &Table{Name: "users"}}
	if got, want := op.Description(), "create table users"; got != want {
		t.Errorf("Description() = %q, want %q", got, want)
	}
}
