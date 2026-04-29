//ff:func feature=validate type=util control=sequence topic=query-structural
//ff:what isIntervalColumn — Column.RawType 가 INTERVAL 인지 판정 (Q-18 filter)

package query

import "github.com/park-jun-woo/yongol/pkg/parser/ddl"

func isIntervalColumn(col ddl.Column) bool {
	return headTokenEquals(col.RawType, "INTERVAL")
}
