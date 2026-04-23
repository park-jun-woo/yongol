//ff:func feature=migration type=accessor control=sequence
//ff:what DropCheck.SQL — ALTER TABLE DROP CONSTRAINT ...;
package migration

import "fmt"

func (op DropCheck) SQL() string {
	return fmt.Sprintf("ALTER TABLE %s DROP CONSTRAINT %s;", op.Table, op.Name)
}
