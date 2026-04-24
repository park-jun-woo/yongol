//ff:func feature=migration type=util control=iteration dimension=1
//ff:what appendIndexLines — 테이블 종속 인덱스를 CREATE INDEX 문으로 렌더 (이름 정렬)
package migration

import (
	"fmt"
	"sort"
	"strings"
)

// appendIndexLines writes one `CREATE [UNIQUE] INDEX ... ON t (...) [WHERE ...]`
// line per index in name-sorted order.
func appendIndexLines(b *strings.Builder, t *Table) {
	sorted := make([]*Index, len(t.Indexes))
	copy(sorted, t.Indexes)
	sort.SliceStable(sorted, func(i, j int) bool { return sorted[i].Name < sorted[j].Name })
	for _, idx := range sorted {
		uniq := ""
		if idx.Unique {
			uniq = "UNIQUE "
		}
		using := ""
		if idx.Method != "" {
			using = " USING " + idx.Method
		}
		where := ""
		if idx.Where != "" {
			where = " WHERE " + idx.Where
		}
		fmt.Fprintf(b, "CREATE %sINDEX %s ON %s%s (%s)%s;\n",
			uniq, idx.Name, t.Name, using, strings.Join(idx.Columns, ", "), where)
	}
}
