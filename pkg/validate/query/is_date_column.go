//ff:func feature=validate type=util control=sequence topic=query-structural
//ff:what isDateColumn — Column.RawType 가 DATE 인지 판정 (Q-16 filter)

package query

import "github.com/park-jun-woo/yongol/pkg/parser/ddl"

func isDateColumn(col ddl.Column) bool {
	return headTokenEquals(col.RawType, "DATE")
}
