//ff:func feature=migration type=accessor control=sequence
//ff:what DropColumn.SQL — ALTER TABLE <t> DROP COLUMN <col>;
package migration

import "fmt"

func (op DropColumn) SQL() string {
	return fmt.Sprintf("ALTER TABLE %s DROP COLUMN %s;", op.Table, op.Column)
}
