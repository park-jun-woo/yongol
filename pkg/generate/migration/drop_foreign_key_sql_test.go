//ff:func feature=migration type=test control=sequence
//ff:what TestDropForeignKey_SQL — ALTER TABLE ... DROP CONSTRAINT ...;
package migration

import "testing"

func TestDropForeignKey_SQL(t *testing.T) {
	op := DropForeignKey{Table: "posts", Name: "posts_user_id_fkey"}
	want := "ALTER TABLE posts DROP CONSTRAINT posts_user_id_fkey;"
	if got := op.SQL(); got != want {
		t.Errorf("SQL() = %q, want %q", got, want)
	}
}
