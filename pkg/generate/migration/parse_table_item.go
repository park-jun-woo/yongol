//ff:func feature=migration type=parser control=selection
//ff:what parseTableItem — CREATE TABLE(...) 내부 한 item(PRIMARY/UNIQUE/FK/CHECK/CONSTRAINT/컬럼) 분기
package migration

import "strings"

// parseTableItem routes a single comma-separated item inside
// CREATE TABLE(...) to its handler.
func parseTableItem(t *Table, it string) {
	if it == "" {
		return
	}
	upper := strings.ToUpper(it)
	switch {
	case strings.HasPrefix(upper, "PRIMARY KEY"):
		t.PrimaryKey = parseColumnList(afterKeyword(it, "PRIMARY KEY"))
	case strings.HasPrefix(upper, "UNIQUE"):
		cols := parseColumnList(afterKeyword(it, "UNIQUE"))
		t.Indexes = append(t.Indexes, &Index{
			Name: UniqueName(t.Name, cols), Columns: cols, Unique: true,
		})
	case strings.HasPrefix(upper, "FOREIGN KEY"):
		if fk := parseTableFK(t, it); fk != nil {
			t.ForeignKeys = append(t.ForeignKeys, fk)
		}
	case strings.HasPrefix(upper, "CHECK"):
		t.Checks = append(t.Checks, parseTableCheck(t, "", it))
	case strings.HasPrefix(upper, "CONSTRAINT"):
		parseNamedConstraint(t, it)
	default:
		parseColumn(t, it)
	}
}
