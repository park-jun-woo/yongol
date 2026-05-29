//ff:func feature=migration type=test control=selection dimension=2
//ff:what TestAddForeignKey_SQL — ON DELETE/UPDATE 절 유무 분기
package migration

import "testing"

func TestAddForeignKey_SQL(t *testing.T) {
	t.Run("no actions", func(t *testing.T) {
		op := AddForeignKey{Table: "posts", FK: &ForeignKey{
			Name: "posts_user_id_fkey", Columns: []string{"user_id"},
			RefTable: "users", RefColumns: []string{"id"},
		}}
		want := "ALTER TABLE posts ADD CONSTRAINT posts_user_id_fkey FOREIGN KEY (user_id) REFERENCES users (id);"
		if got := op.SQL(); got != want {
			t.Errorf("SQL() = %q, want %q", got, want)
		}
	})
	t.Run("with on delete and update", func(t *testing.T) {
		op := AddForeignKey{Table: "posts", FK: &ForeignKey{
			Name: "fk", Columns: []string{"user_id"},
			RefTable: "users", RefColumns: []string{"id"},
			OnDelete: "CASCADE", OnUpdate: "RESTRICT",
		}}
		want := "ALTER TABLE posts ADD CONSTRAINT fk FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE ON UPDATE RESTRICT;"
		if got := op.SQL(); got != want {
			t.Errorf("SQL() = %q, want %q", got, want)
		}
	})
}
