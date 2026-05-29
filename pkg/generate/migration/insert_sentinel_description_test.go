//ff:func feature=migration type=test control=sequence
//ff:what TestInsertSentinel_Description — "insert sentinel <table>"
package migration

import "testing"

func TestInsertSentinel_Description(t *testing.T) {
	op := InsertSentinel{Table: "roles", Body: "INSERT INTO roles ...;"}
	if got, want := op.Description(), "insert sentinel roles"; got != want {
		t.Errorf("Description() = %q, want %q", got, want)
	}
}
