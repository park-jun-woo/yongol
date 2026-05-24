//ff:func feature=orchestrator type=parser control=sequence
//ff:what parseOneSelectColumn — SELECT 항목 하나에서 실질적 컬럼 이름 추출 (alias/qualifier 처리)

package sqlc

import "strings"

// parseOneSelectColumn extracts the effective column name from a single
// SELECT-list item. Handles:
//   - `column_name` → "column_name"
//   - `table.column_name` → "column_name"
//   - `expr AS alias_name` → "alias_name"
//   - `*` embedded in a multi-column list → skipped (returns "")
func parseOneSelectColumn(item string) string {
	if item == "" || item == "*" {
		return ""
	}
	// Handle AS alias (case-insensitive).
	asIdx := strings.LastIndex(strings.ToUpper(item), " AS ")
	if asIdx >= 0 {
		alias := strings.TrimSpace(item[asIdx+4:])
		// Remove surrounding quotes if present.
		alias = strings.Trim(alias, `"`)
		return alias
	}
	// Remove table qualifier.
	if dotIdx := strings.LastIndex(item, "."); dotIdx >= 0 {
		item = item[dotIdx+1:]
	}
	// Remove surrounding whitespace and quotes.
	item = strings.TrimSpace(item)
	item = strings.Trim(item, `"`)
	// Reject complex expressions (contain spaces, parens, etc.) —
	// they are computed columns that need an AS alias to be resolvable.
	if strings.ContainsAny(item, " ()+/-") {
		return ""
	}
	return item
}
