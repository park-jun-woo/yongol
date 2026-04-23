//ff:func feature=migration type=accessor control=selection
//ff:what AlterColumnDefault.SQL — SET/DROP DEFAULT
package migration

import "fmt"

func (op AlterColumnDefault) SQL() string {
	if op.To == "" {
		return fmt.Sprintf("ALTER TABLE %s ALTER COLUMN %s DROP DEFAULT;", op.Table, op.Column)
	}
	return fmt.Sprintf("ALTER TABLE %s ALTER COLUMN %s SET DEFAULT %s;", op.Table, op.Column, op.To)
}
