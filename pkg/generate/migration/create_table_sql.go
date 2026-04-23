//ff:func feature=migration type=accessor control=iteration dimension=1
//ff:what CreateTable.SQL — CREATE TABLE ... ( ... ) SQL 렌더링
package migration

import (
	"fmt"
	"strings"
)

func (op CreateTable) SQL() string {
	b := strings.Builder{}
	fmt.Fprintf(&b, "CREATE TABLE %s (", op.Table.Name)
	for i, c := range op.Table.Columns {
		if i == 0 {
			b.WriteString("\n    ")
		} else {
			b.WriteString(",\n    ")
		}
		b.WriteString(renderColumn(c))
	}
	if len(op.Table.PrimaryKey) > 0 {
		fmt.Fprintf(&b, ",\n    PRIMARY KEY (%s)", strings.Join(op.Table.PrimaryKey, ", "))
	}
	b.WriteString("\n);")
	return b.String()
}
