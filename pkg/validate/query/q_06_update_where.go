//ff:func feature=validate type=rule control=iteration dimension=1 topic=query-structural
//ff:what Q-06 — UPDATE 문에 WHERE 필수

package query

import (
	"regexp"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

var updateStartRe = regexp.MustCompile(`(?i)^\s*UPDATE\s+\w+`)

// q06UpdateWhere validates Q-06: every UPDATE must contain a WHERE clause.
// Full-table UPDATE is almost always unintended. Escape with `-- @allow-truncate`
// (reusing the Q-05 hatch semantics) when truly bulk.
func q06UpdateWhere(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, q := range fs.SQLcQueries {
		if d, ok := q06CheckUpdateWhere(q); ok {
			diags = append(diags, d)
		}
	}
	return diags
}
