//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what sortedColumnKeys — map[string]ddl.Column 의 키를 정렬된 []string로 반환

package sqlcpost

import (
	"sort"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// sortedColumnKeys returns the keys of a parsed Column map in ascending
// order — the Column-typed counterpart of sortedKeys, used as a
// deterministic fallback when a DDL table provides no explicit
// ColumnOrder.
func sortedColumnKeys(m map[string]ddl.Column) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
