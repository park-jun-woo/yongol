//ff:func feature=manifest type=parser control=selection
//ff:what dispatchConstraint — DDL 라인 종류에 따라 제약조건/컬럼 파싱 분기
package ddl

import "strings"

func dispatchConstraint(line, upper string, t *Table, tables map[string]*Table, pendingArchived bool) {
	switch {
	case strings.HasPrefix(upper, "CONSTRAINT") || strings.HasPrefix(upper, "FOREIGN"):
		if fk, ok := parseConstraintFK(line); ok {
			t.ForeignKeys = append(t.ForeignKeys, fk)
		}
	case strings.HasPrefix(upper, "PRIMARY"):
		t.PrimaryKey = extractParenColumns(line)
	case strings.HasPrefix(upper, "UNIQUE"):
		appendUniqueIndex(line, t)
	case strings.HasPrefix(upper, "CHECK"):
		applyCheckEnum(line, "", t)
	case line != "":
		parseColumnDef(line, upper, t, pendingArchived)
	}
}
