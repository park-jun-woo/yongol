//ff:func feature=migration type=test control=sequence
//ff:what TestDropCheck_Description — "drop check <name>"
package migration

import "testing"

func TestDropCheck_Description(t *testing.T) {
	op := DropCheck{Table: "users", Name: "users_age_check"}
	if got, want := op.Description(), "drop check users_age_check"; got != want {
		t.Errorf("Description() = %q, want %q", got, want)
	}
}
