//ff:func feature=manifest type=util control=iteration dimension=1
//ff:what parseInlineFK — 컬럼 정의에서 인라인 REFERENCES 파싱
package ddl

import "strings"

func parseInlineFK(colName string, parts []string) (ForeignKey, bool) {
	for i, p := range parts {
		if strings.ToUpper(p) != "REFERENCES" || i+1 >= len(parts) {
			continue
		}
		ref := strings.TrimSuffix(parts[i+1], ",")
		refTable, refCol := parseRef(ref)
		if refTable != "" {
			return ForeignKey{Column: colName, RefTable: refTable, RefColumn: refCol}, true
		}
	}
	return ForeignKey{}, false
}
