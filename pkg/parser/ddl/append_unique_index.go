//ff:func feature=manifest type=util control=sequence
//ff:what appendUniqueIndex — UNIQUE 제약에서 unique index 추가
package ddl

import "strings"

func appendUniqueIndex(line string, t *Table) {
	cols := extractParenColumns(line)
	if len(cols) == 0 {
		return
	}
	t.Indexes = append(t.Indexes, Index{Name: "unique_" + strings.Join(cols, "_"), Columns: cols, IsUnique: true})
}
