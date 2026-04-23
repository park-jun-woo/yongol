//ff:func feature=migration type=util control=iteration dimension=1
//ff:what UniqueName — PostgreSQL 기본 UNIQUE 제약 이름 (<table>_<col>_..._key)
package migration

import "strings"

// UniqueName returns the canonical UNIQUE constraint / index name.
// PostgreSQL joins each participating column with `_` and suffixes `_key`.
func UniqueName(table string, columns []string) string {
	b := strings.Builder{}
	b.WriteString(strings.ToLower(table))
	for _, c := range columns {
		b.WriteByte('_')
		b.WriteString(strings.ToLower(c))
	}
	b.WriteString("_key")
	return b.String()
}
