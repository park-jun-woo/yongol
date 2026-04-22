//ff:func feature=validate type=rule control=iteration dimension=1 topic=query-structural
//ff:what Q-09 — :exec 쿼리에 SELECT 반환 있으면 ERROR

package query

import (
	"regexp"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// topLevelSelectRe matches a top-level SELECT (first non-whitespace statement token).
var topLevelSelectRe = regexp.MustCompile(`(?i)^\s*SELECT\s+`)
var returningWordRe = regexp.MustCompile(`(?i)(?:^|\s)RETURNING\s+`)

// q09SelectOnExec validates Q-09: :exec queries must not return rows. The
// violation is a *top-level* SELECT (stand-alone query that fetches rows)
// or a RETURNING clause on DML — both produce output that :exec discards.
// INSERT ... SELECT is allowed: the inner SELECT feeds the INSERT rather
// than returning to the caller.
func q09SelectOnExec(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, q := range fs.SQLcQueries {
		if d, ok := q09CheckSelectOnExec(q); ok {
			diags = append(diags, d)
		}
	}
	return diags
}
