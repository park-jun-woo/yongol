//ff:func feature=migration type=accessor control=sequence
//ff:what RenameColumn.SQL — ALTER TABLE RENAME COLUMN ... TO ...
package migration

import "fmt"

func (op RenameColumn) SQL() string {
	return fmt.Sprintf("ALTER TABLE %s RENAME COLUMN %s TO %s;", op.Table, op.From, op.To)
}
