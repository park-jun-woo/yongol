//ff:func feature=migration type=accessor control=sequence
//ff:what DropForeignKey.SQL — ALTER TABLE DROP CONSTRAINT ...;
package migration

import "fmt"

func (op DropForeignKey) SQL() string {
	return fmt.Sprintf("ALTER TABLE %s DROP CONSTRAINT %s;", op.Table, op.Name)
}
