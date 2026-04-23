//ff:func feature=validate type=util control=iteration dimension=1 topic=migration-hints
//ff:what Migration Table 에 지정된 컬럼이 있는지 검사

package migration

import (
	"github.com/park-jun-woo/yongol/pkg/generate/migration"
)

// hasColumn reports whether t carries a column with the given name.
func hasColumn(t *migration.Table, name string) bool {
	for _, c := range t.Columns {
		if c.Name == name {
			return true
		}
	}
	return false
}
