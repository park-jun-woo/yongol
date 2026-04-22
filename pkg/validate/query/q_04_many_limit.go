//ff:func feature=validate type=rule control=iteration dimension=1 topic=query-structural
//ff:what Q-04 — :many 쿼리에 LIMIT 누락 시 WARNING

package query

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// q04ManyLimit validates Q-04: :many queries without a LIMIT clause emit a
// WARNING. Unbounded list queries risk memory blowup for large datasets.
// Escape with `-- @no-pagination` when intentional (cursor-based or small
// fixed lists).
func q04ManyLimit(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, q := range fs.SQLcQueries {
		if d, ok := q04CheckManyLimit(q); ok {
			diags = append(diags, d)
		}
	}
	return diags
}
