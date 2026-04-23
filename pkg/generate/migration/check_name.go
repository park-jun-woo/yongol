//ff:func feature=migration type=util control=sequence
//ff:what CheckName — PostgreSQL 기본 CHECK 제약 이름 (<table>_<column>_check)
package migration

import "strings"

// CheckName returns the canonical CHECK constraint name bound to a
// single column (PostgreSQL uses <table>_<col>_check when inferrable).
func CheckName(table, column string) string {
	return strings.ToLower(table) + "_" + strings.ToLower(column) + "_check"
}
