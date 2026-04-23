//ff:func feature=migration type=accessor control=sequence
//ff:what RenameTable.SQL — ALTER TABLE ... RENAME TO ...
package migration

import "fmt"

func (op RenameTable) SQL() string {
	return fmt.Sprintf("ALTER TABLE %s RENAME TO %s;", op.From, op.To)
}
