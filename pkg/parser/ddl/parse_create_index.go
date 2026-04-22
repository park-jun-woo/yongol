//ff:func feature=manifest type=util control=iteration dimension=1
//ff:what parseCreateIndex — CREATE INDEX 문을 파싱
package ddl

import "strings"

func parseCreateIndex(line string, tables map[string]*Table) {
	upper := strings.ToUpper(line)
	onIdx := strings.Index(upper, " ON ")
	if onIdx < 0 {
		return
	}
	parts := strings.Fields(line[:onIdx])
	idxName := ""
	for i, p := range parts {
		if strings.ToUpper(p) == "INDEX" && i+1 < len(parts) {
			idxName = parts[i+1]
			break
		}
	}
	after := strings.TrimSpace(line[onIdx+4:])
	afterParts := strings.SplitN(after, "(", 2)
	if len(afterParts) < 2 {
		return
	}
	tableName := strings.TrimSpace(afterParts[0])
	cols := extractParenColumns("(" + afterParts[1])
	isUnique := strings.Contains(upper, "UNIQUE")
	if t := tables[tableName]; t != nil && len(cols) > 0 {
		t.Indexes = append(t.Indexes, Index{Name: idxName, Columns: cols, IsUnique: isUnique})
	}
}
