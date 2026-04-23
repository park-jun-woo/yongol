//ff:func feature=migration type=util control=iteration dimension=1
//ff:what CanonicalSQL — Schema AST → 정규 SQL (스냅샷 포맷, 테이블 알파벳 순서)
package migration

import (
	"sort"
	"strings"
)

// CanonicalSQL renders the entire Schema as deterministic DDL suitable
// for diff-based tooling. Tables are sorted alphabetically; within a
// table, column order is preserved (matches DDL original order).
func CanonicalSQL(s *Schema) string {
	if s == nil {
		return ""
	}
	names := make([]string, 0, len(s.Tables))
	for n := range s.Tables {
		names = append(names, n)
	}
	sort.Strings(names)
	var out strings.Builder
	for i, n := range names {
		if i > 0 {
			out.WriteByte('\n')
		}
		out.WriteString(renderTable(s.Tables[n]))
		out.WriteByte('\n')
	}
	return out.String()
}
