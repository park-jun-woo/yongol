//ff:func feature=migration type=test control=sequence
//ff:what TestAddForeignKey_Description — 설명에 FK 이름 포함
package migration

import "testing"

func TestAddForeignKey_Description(t *testing.T) {
	op := AddForeignKey{FK: &ForeignKey{Name: "posts_user_id_fkey"}}
	if got, want := op.Description(), "add FK posts_user_id_fkey"; got != want {
		t.Errorf("Description() = %q, want %q", got, want)
	}
}
