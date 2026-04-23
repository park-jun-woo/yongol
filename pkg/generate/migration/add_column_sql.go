//ff:func feature=migration type=accessor control=sequence
//ff:what AddColumn.SQL — ALTER TABLE <t> ADD COLUMN <col>;
package migration

import (
	"fmt"
	"strings"
)

func (op AddColumn) SQL() string {
	b := strings.Builder{}
	fmt.Fprintf(&b, "ALTER TABLE %s ADD COLUMN %s;",
		op.Table, renderColumn(op.Column))
	return b.String()
}
