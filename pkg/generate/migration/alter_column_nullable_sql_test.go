//ff:func feature=migration type=test control=selection dimension=3
//ff:what TestAlterColumnNullable_SQL — SET/DROP NOT NULL 및 backfill UPDATE 분기
package migration

import "testing"

func TestAlterColumnNullable_SQL(t *testing.T) {
	cases := []struct {
		name string
		op   AlterColumnNullable
		want string
	}{
		{
			"drop not null",
			AlterColumnNullable{Table: "t", Column: "c", To: true},
			"ALTER TABLE t ALTER COLUMN c DROP NOT NULL;",
		},
		{
			"set not null no backfill",
			AlterColumnNullable{Table: "t", Column: "c", To: false},
			"ALTER TABLE t ALTER COLUMN c SET NOT NULL;",
		},
		{
			"set not null with backfill",
			AlterColumnNullable{Table: "t", Column: "c", To: false, Backfill: "0"},
			"UPDATE t SET c = 0 WHERE c IS NULL;\nALTER TABLE t ALTER COLUMN c SET NOT NULL;",
		},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			if got := c.op.SQL(); got != c.want {
				t.Errorf("SQL() = %q, want %q", got, c.want)
			}
		})
	}
}
