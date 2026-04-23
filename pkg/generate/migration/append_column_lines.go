//ff:func feature=migration type=util control=iteration dimension=1
//ff:what appendColumnLines — CREATE TABLE 내부 컬럼 line 을 builder 에 추가
package migration

import "strings"

// appendColumnLines writes `\n    <col>` lines (with leading comma on
// subsequent items) for each column.
func appendColumnLines(b *strings.Builder, cols []*Column) {
	for i, c := range cols {
		if i == 0 {
			b.WriteString("\n    ")
		} else {
			b.WriteString(",\n    ")
		}
		b.WriteString(renderColumn(c))
	}
}
