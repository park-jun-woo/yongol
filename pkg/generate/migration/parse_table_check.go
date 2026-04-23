//ff:func feature=migration type=parser control=sequence
//ff:what parseTableCheck — CHECK (expr) 절을 CheckConstraint 로 변환
package migration

import "strings"

// parseTableCheck parses a CHECK (expr) clause. If `name` is empty, a
// default <table>_check name is used.
func parseTableCheck(t *Table, name, item string) *CheckConstraint {
	expr := innerParens(afterKeyword(item, "CHECK"))
	if name == "" {
		name = strings.ToLower(t.Name) + "_check"
	}
	return &CheckConstraint{Name: name, Expression: strings.TrimSpace(expr)}
}
