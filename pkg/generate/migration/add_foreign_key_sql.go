//ff:func feature=migration type=accessor control=sequence
//ff:what AddForeignKey.SQL — ALTER TABLE ADD CONSTRAINT ... FOREIGN KEY (...) REFERENCES ...
package migration

import (
	"fmt"
	"strings"
)

func (op AddForeignKey) SQL() string {
	b := strings.Builder{}
	fmt.Fprintf(&b, "ALTER TABLE %s ADD CONSTRAINT %s FOREIGN KEY (%s) REFERENCES %s (%s)",
		op.Table, op.FK.Name,
		strings.Join(op.FK.Columns, ", "),
		op.FK.RefTable, strings.Join(op.FK.RefColumns, ", "))
	if op.FK.OnDelete != "" {
		fmt.Fprintf(&b, " ON DELETE %s", op.FK.OnDelete)
	}
	if op.FK.OnUpdate != "" {
		fmt.Fprintf(&b, " ON UPDATE %s", op.FK.OnUpdate)
	}
	b.WriteByte(';')
	return b.String()
}
