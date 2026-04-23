//ff:func feature=migration type=util control=iteration dimension=1
//ff:what FKName — PostgreSQL 기본 FK 제약 이름 (<table>_<col>_..._fkey)
package migration

import "strings"

// FKName returns the canonical FK constraint name.
func FKName(table string, columns []string) string {
	b := strings.Builder{}
	b.WriteString(strings.ToLower(table))
	for _, c := range columns {
		b.WriteByte('_')
		b.WriteString(strings.ToLower(c))
	}
	b.WriteString("_fkey")
	return b.String()
}
