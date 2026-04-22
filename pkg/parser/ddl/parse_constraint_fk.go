//ff:func feature=manifest type=util control=sequence
//ff:what parseConstraintFK — 독립 FOREIGN KEY 절을 파싱
package ddl

import "strings"

func parseConstraintFK(line string) (ForeignKey, bool) {
	upper := strings.ToUpper(line)
	fkIdx := strings.Index(upper, "FOREIGN KEY")
	refIdx := strings.Index(upper, "REFERENCES")
	if fkIdx < 0 || refIdx < 0 {
		return ForeignKey{}, false
	}
	between := line[fkIdx+len("FOREIGN KEY") : refIdx]
	col := extractParenContent(between)
	if col == "" {
		return ForeignKey{}, false
	}
	after := strings.TrimSpace(line[refIdx+len("REFERENCES"):])
	after = strings.TrimSuffix(after, ",")
	refTable, refCol := parseRef(after)
	if refTable == "" {
		return ForeignKey{}, false
	}
	return ForeignKey{Column: col, RefTable: refTable, RefColumn: refCol}, true
}
