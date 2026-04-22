//ff:func feature=migration type=util control=sequence
//ff:what 제약 이름 자동 생성 — PostgreSQL 기본 규칙(<table>_<col>_pkey/key/fkey/check)을 재현
package migration

import "strings"

// PKName returns the canonical PK constraint name (PostgreSQL default).
func PKName(table string) string {
	return strings.ToLower(table) + "_pkey"
}

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

// CheckName returns the canonical CHECK constraint name bound to a
// single column (PostgreSQL uses <table>_<col>_check when inferrable).
func CheckName(table, column string) string {
	return strings.ToLower(table) + "_" + strings.ToLower(column) + "_check"
}
