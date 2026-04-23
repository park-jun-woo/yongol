//ff:func feature=migration type=util control=iteration dimension=1
//ff:what appendForeignKeyClauses — FK 제약을 inline 절로 렌더
package migration

import (
	"fmt"
	"strings"
)

// appendForeignKeyClauses writes one `,\n    CONSTRAINT <name> FOREIGN
// KEY (...) REFERENCES ... [ON ...]` line per FK.
func appendForeignKeyClauses(b *strings.Builder, fks []*ForeignKey) {
	for _, fk := range fks {
		fmt.Fprintf(b, ",\n    CONSTRAINT %s FOREIGN KEY (%s) REFERENCES %s (%s)",
			fk.Name,
			strings.Join(fk.Columns, ", "),
			fk.RefTable,
			strings.Join(fk.RefColumns, ", "))
		if fk.OnDelete != "" {
			fmt.Fprintf(b, " ON DELETE %s", fk.OnDelete)
		}
		if fk.OnUpdate != "" {
			fmt.Fprintf(b, " ON UPDATE %s", fk.OnUpdate)
		}
	}
}
