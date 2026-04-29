//ff:func feature=validate type=util control=sequence topic=query-structural
//ff:what isTimestampColumn — Column.RawType 가 TIMESTAMP (no TZ) 인지 판정 (Q-15 filter)

package query

import "github.com/park-jun-woo/yongol/pkg/parser/ddl"

func isTimestampColumn(col ddl.Column) bool {
	return headTokenEquals(col.RawType, "TIMESTAMP")
}
