//ff:func feature=migration type=parser control=selection
//ff:what parseNamedConstraint — CONSTRAINT <name> PRIMARY KEY/UNIQUE/FK/CHECK 분기
package migration

import "strings"

// parseNamedConstraint handles `CONSTRAINT <name> <body>` inside
// CREATE TABLE(...).
func parseNamedConstraint(t *Table, item string) {
	toks := tokenizeColumnDef(item)
	if len(toks) < 3 {
		return
	}
	name := canonIdent(toks[1])
	body := strings.Join(toks[2:], " ")
	upper := strings.ToUpper(body)
	switch {
	case strings.HasPrefix(upper, "PRIMARY KEY"):
		t.PrimaryKey = parseColumnList(afterKeyword(body, "PRIMARY KEY"))
	case strings.HasPrefix(upper, "UNIQUE"):
		cols := parseColumnList(afterKeyword(body, "UNIQUE"))
		t.Indexes = append(t.Indexes, &Index{Name: name, Columns: cols, Unique: true})
	case strings.HasPrefix(upper, "FOREIGN KEY"):
		if fk := parseTableFK(t, body); fk != nil {
			fk.Name = name
			t.ForeignKeys = append(t.ForeignKeys, fk)
		}
	case strings.HasPrefix(upper, "CHECK"):
		t.Checks = append(t.Checks, parseTableCheck(t, name, body))
	}
}
