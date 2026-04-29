//ff:func feature=validate type=util control=sequence topic=query-structural
//ff:what isInetColumn — Column.RawType 가 INET/CIDR 인지 판정 (Q-17 filter)

package query

import "github.com/park-jun-woo/yongol/pkg/parser/ddl"

func isInetColumn(col ddl.Column) bool {
	return headTokenEquals(col.RawType, "INET") || headTokenEquals(col.RawType, "CIDR")
}
