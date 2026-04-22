//ff:func feature=migration type=util control=sequence
//ff:what CanonicalSQL — Schema AST → 정규 SQL (스냅샷 포맷, 테이블 알파벳 순서)
package migration

import (
	"fmt"
	"sort"
	"strings"
)

// CanonicalSQL renders the entire Schema as deterministic DDL suitable
// for diff-based tooling. Tables are sorted alphabetically; within a
// table, column order is preserved (matches DDL original order).
func CanonicalSQL(s *Schema) string {
	if s == nil {
		return ""
	}
	names := make([]string, 0, len(s.Tables))
	for n := range s.Tables {
		names = append(names, n)
	}
	sort.Strings(names)
	var out strings.Builder
	for i, n := range names {
		if i > 0 {
			out.WriteByte('\n')
		}
		out.WriteString(renderTable(s.Tables[n]))
		out.WriteByte('\n')
	}
	// Then CREATE INDEX / ALTER TABLE ADD FK at the end in deterministic
	// order. We already rendered table-local FK/checks/PK; standalone
	// indexes are emitted per-table.
	return out.String()
}

func renderTable(t *Table) string {
	b := strings.Builder{}
	fmt.Fprintf(&b, "CREATE TABLE %s (", t.Name)
	for i, c := range t.Columns {
		if i == 0 {
			b.WriteString("\n    ")
		} else {
			b.WriteString(",\n    ")
		}
		b.WriteString(renderColumn(c))
	}
	if len(t.PrimaryKey) > 0 {
		fmt.Fprintf(&b, ",\n    PRIMARY KEY (%s)", strings.Join(t.PrimaryKey, ", "))
	}
	// FK rendered inline so canonical is self-contained.
	for _, fk := range t.ForeignKeys {
		fmt.Fprintf(&b, ",\n    CONSTRAINT %s FOREIGN KEY (%s) REFERENCES %s (%s)",
			fk.Name,
			strings.Join(fk.Columns, ", "),
			fk.RefTable,
			strings.Join(fk.RefColumns, ", "))
		if fk.OnDelete != "" {
			fmt.Fprintf(&b, " ON DELETE %s", fk.OnDelete)
		}
		if fk.OnUpdate != "" {
			fmt.Fprintf(&b, " ON UPDATE %s", fk.OnUpdate)
		}
	}
	for _, chk := range t.Checks {
		fmt.Fprintf(&b, ",\n    CONSTRAINT %s CHECK (%s)", chk.Name, chk.Expression)
	}
	b.WriteString("\n);\n")
	// Non-inline indexes (UNIQUE inline is already a constraint — but
	// v1 stores UNIQUE as Index entries; render as CREATE UNIQUE INDEX
	// here for canonical consistency).
	// Sort for determinism.
	sortedIx := make([]*Index, len(t.Indexes))
	copy(sortedIx, t.Indexes)
	sort.SliceStable(sortedIx, func(i, j int) bool { return sortedIx[i].Name < sortedIx[j].Name })
	for _, idx := range sortedIx {
		uniq := ""
		if idx.Unique {
			uniq = "UNIQUE "
		}
		where := ""
		if idx.Where != "" {
			where = " WHERE " + idx.Where
		}
		fmt.Fprintf(&b, "CREATE %sINDEX %s ON %s (%s)%s;\n",
			uniq, idx.Name, t.Name, strings.Join(idx.Columns, ", "), where)
	}
	return b.String()
}
