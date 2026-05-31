//ff:func feature=migration type=test control=iteration dimension=1
//ff:what TestMnemonicFor — Operation 별 snake_case mnemonic 생성
package migration

import (
	"testing"
)

func TestMnemonicFor(t *testing.T) {
	cases := []struct {
		name string
		op   Operation
		want string
	}{
		{"create table", CreateTable{Table: &Table{Name: "Users"}}, "create_users"},
		{"drop table", DropTable{Name: "users"}, "drop_users"},
		{"add column", AddColumn{Table: "users", Column: &Column{Name: "age"}}, "add_users_age"},
		{"alter type", AlterColumnType{Table: "users", Column: "age"}, "alter_users_age_type"},
		{"set nullable", AlterColumnNullable{Table: "users", Column: "age", To: true}, "nullable_users_age"},
		{"set not null", AlterColumnNullable{Table: "users", Column: "age", To: false}, "notnull_users_age"},
		{"default", AlterColumnDefault{Table: "users", Column: "age"}, "default_users_age"},
		{"add fk", AddForeignKey{FK: &ForeignKey{Name: "posts_user_id_fkey"}}, "add_fk_posts_user_id_fkey"},
		{"rename column", RenameColumn{Table: "users", From: "a", To: "b"}, "rename_users_a_to_b"},
		{"rename table", RenameTable{From: "a", To: "b"}, "rename_table_a_to_b"},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			if got := mnemonicFor(c.op); got != c.want {
				t.Errorf("mnemonicFor(%T) = %q, want %q", c.op, got, c.want)
			}
		})
	}
}
