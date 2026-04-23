//ff:func feature=migration type=accessor control=sequence
//ff:what DropTable.SQL — DROP TABLE <name>;
package migration

import "fmt"

func (op DropTable) SQL() string { return fmt.Sprintf("DROP TABLE %s;", op.Name) }
