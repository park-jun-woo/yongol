//ff:func feature=migration type=parser control=sequence
//ff:what applyInlineCheck — 컬럼 뒤의 CHECK (expr) 인라인 제약 수집
package migration

import "strings"

// applyInlineCheck handles CHECK (...) that appears inline on a column
// line. Returns the number of tokens consumed.
func applyInlineCheck(t *Table, col *Column, rest []string, i int) int {
	if i+1 >= len(rest) || !strings.HasPrefix(rest[i+1], "(") {
		return 1
	}
	t.Checks = append(t.Checks, &CheckConstraint{
		Name:       CheckName(t.Name, col.Name),
		Expression: innerParens(rest[i+1]),
	})
	return 2
}
