//ff:func feature=validate type=util control=iteration dimension=2 topic=query-structural
//ff:what buildSensitiveColumnIndex — DDLTables → tableName(lower) → @sensitive 컬럼 정렬 목록

package query

import (
	"sort"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// buildSensitiveColumnIndex groups every DDL table's @sensitive columns into
// a lowercase-tableName keyed map used by Q-07 for SELECT * detection.
func buildSensitiveColumnIndex(tables []ddl.Table) map[string][]string {
	idx := make(map[string][]string)
	for _, t := range tables {
		if len(t.SensitiveColumns) == 0 {
			continue
		}
		var cols []string
		for col := range t.SensitiveColumns {
			cols = append(cols, col)
		}
		sort.Strings(cols)
		idx[strings.ToLower(t.Name)] = cols
	}
	return idx
}
