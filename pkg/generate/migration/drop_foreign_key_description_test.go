//ff:func feature=migration type=test control=sequence
//ff:what TestDropForeignKey_Description — "drop FK <name>"
package migration

import "testing"

func TestDropForeignKey_Description(t *testing.T) {
	op := DropForeignKey{Table: "posts", Name: "posts_user_id_fkey"}
	if got, want := op.Description(), "drop FK posts_user_id_fkey"; got != want {
		t.Errorf("Description() = %q, want %q", got, want)
	}
}
