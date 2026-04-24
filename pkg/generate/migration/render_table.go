//ff:func feature=migration type=util control=sequence
//ff:what renderTable — 한 테이블을 스냅샷용 정규 CREATE TABLE + CREATE INDEX 로 렌더
package migration

import (
	"fmt"
	"strings"
)

// renderTable emits the canonical CREATE TABLE for t plus any
// CREATE INDEX lines (sorted by name) and — for snapshot fidelity —
// any `@sentinel` INSERT blocks attached to the table. Sentinel INSERTs
// are inserted between the closing `);` of CREATE TABLE and the
// CREATE INDEX lines, separated by blank lines, so the snapshot stays
// readable and any edit to a sentinel invalidates the snapshot hash.
func renderTable(t *Table) string {
	b := strings.Builder{}
	fmt.Fprintf(&b, "CREATE TABLE %s (", t.Name)
	appendColumnLines(&b, t.Columns)
	appendPrimaryKeyClause(&b, t.PrimaryKey)
	appendForeignKeyClauses(&b, t.ForeignKeys)
	appendCheckClauses(&b, t.Checks)
	b.WriteString("\n);\n")
	for _, s := range t.Sentinels {
		b.WriteByte('\n')
		body := strings.TrimRight(s.SQL, "\n")
		b.WriteString(body)
		b.WriteByte('\n')
	}
	appendIndexLines(&b, t)
	return b.String()
}
