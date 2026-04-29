//ff:func feature=validate type=util control=sequence topic=query-structural
//ff:what isTimestamptzColumn — Column.RawType 가 TIMESTAMPTZ 인지 판정 (Q-14 filter)

package query

import "github.com/park-jun-woo/yongol/pkg/parser/ddl"

func isTimestamptzColumn(col ddl.Column) bool {
	return headTokenEquals(col.RawType, "TIMESTAMPTZ")
}
