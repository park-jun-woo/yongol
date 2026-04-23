//ff:func feature=migration type=accessor control=selection
//ff:what AlterColumnNullable.SQL — SET/DROP NOT NULL (+ backfill UPDATE)
package migration

import "fmt"

func (op AlterColumnNullable) SQL() string {
	verb := "SET NOT NULL"
	if op.To {
		verb = "DROP NOT NULL"
	}
	if op.To == false && op.Backfill != "" {
		return fmt.Sprintf("UPDATE %s SET %s = %s WHERE %s IS NULL;\nALTER TABLE %s ALTER COLUMN %s SET NOT NULL;",
			op.Table, op.Column, op.Backfill, op.Column, op.Table, op.Column)
	}
	return fmt.Sprintf("ALTER TABLE %s ALTER COLUMN %s %s;", op.Table, op.Column, verb)
}
