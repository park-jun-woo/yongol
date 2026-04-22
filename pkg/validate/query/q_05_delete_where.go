//ff:func feature=validate type=rule control=iteration dimension=1 topic=query-structural
//ff:what Q-05 — DELETE 문에 WHERE 필수

package query

import (
	"regexp"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

var deleteStartRe = regexp.MustCompile(`(?i)^\s*DELETE\s+FROM`)

// q05DeleteWhere validates Q-05: every DELETE must contain a WHERE clause.
// Full-table DELETE is usually an accidental production disaster; require
// explicit `-- @allow-truncate` escape when truly intended.
func q05DeleteWhere(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, q := range fs.SQLcQueries {
		if d, ok := q05CheckDeleteWhere(q); ok {
			diags = append(diags, d)
		}
	}
	return diags
}
