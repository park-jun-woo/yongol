//ff:func feature=migration type=util control=sequence
//ff:what PKName — PostgreSQL 기본 PK 제약 이름 (<table>_pkey)
package migration

import "strings"

// PKName returns the canonical PK constraint name (PostgreSQL default).
func PKName(table string) string {
	return strings.ToLower(table) + "_pkey"
}
