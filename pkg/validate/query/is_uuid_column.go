//ff:func feature=validate type=util control=sequence topic=query-structural
//ff:what isUUIDColumn — Column.RawType 의 head 가 UUID 인지 판정 (Q-12 filter)

package query

import "github.com/park-jun-woo/yongol/pkg/parser/ddl"

// isUUIDColumn reports whether col declares PG `UUID` (head-token match
// after stripping array marker / param list). Shared filter consumed by
// Q-12 and the matching ddl.Column scan in checkPgtypeOverride.
func isUUIDColumn(col ddl.Column) bool {
	return headTokenEquals(col.RawType, "UUID")
}
