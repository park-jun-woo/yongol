//ff:func feature=migration type=accessor control=sequence
//ff:what DropIndex.SQL — DROP INDEX <name>;
package migration

import "fmt"

func (op DropIndex) SQL() string { return fmt.Sprintf("DROP INDEX %s;", op.Name) }
