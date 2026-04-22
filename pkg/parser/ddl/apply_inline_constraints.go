//ff:func feature=manifest type=util control=sequence
//ff:what applyInlineConstraints — 컬럼 라인의 인라인 PK, UNIQUE, FK 반영
package ddl

import "strings"

func applyInlineConstraints(t *Table, upper, colName string, parts []string) {
	if strings.Contains(upper, "PRIMARY KEY") {
		t.PrimaryKey = append(t.PrimaryKey, colName)
	}
	if strings.Contains(upper, "UNIQUE") && !strings.Contains(upper, "PRIMARY") {
		t.Indexes = append(t.Indexes, Index{Name: colName + "_unique", Columns: []string{colName}, IsUnique: true})
	}
	if fk, ok := parseInlineFK(colName, parts); ok {
		t.ForeignKeys = append(t.ForeignKeys, fk)
	}
}
