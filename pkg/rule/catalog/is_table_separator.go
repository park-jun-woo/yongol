//ff:func feature=rule type=util control=iteration dimension=2 topic=catalog
//ff:what isTableSeparator — `|---|---|...|` markdown 테이블 구분자 검사
package catalog

import "strings"

// isTableSeparator recognises the `|---|---|...|` row immediately below a
// markdown table header.
func isTableSeparator(row string) bool {
	cells := splitRow(row)
	if len(cells) == 0 {
		return false
	}
	for _, c := range cells {
		c = strings.TrimSpace(c)
		if c == "" {
			return false
		}
		for _, r := range c {
			if r != '-' && r != ':' {
				return false
			}
		}
	}
	return true
}
