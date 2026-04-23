//ff:func feature=migration type=util control=sequence
//ff:what renderTable — 한 테이블을 스냅샷용 정규 CREATE TABLE + CREATE INDEX 로 렌더
package migration

import (
	"fmt"
	"strings"
)

// renderTable emits the canonical CREATE TABLE for t plus any
// CREATE INDEX lines (sorted by name).
func renderTable(t *Table) string {
	b := strings.Builder{}
	fmt.Fprintf(&b, "CREATE TABLE %s (", t.Name)
	appendColumnLines(&b, t.Columns)
	appendPrimaryKeyClause(&b, t.PrimaryKey)
	appendForeignKeyClauses(&b, t.ForeignKeys)
	appendCheckClauses(&b, t.Checks)
	b.WriteString("\n);\n")
	appendIndexLines(&b, t)
	return b.String()
}
