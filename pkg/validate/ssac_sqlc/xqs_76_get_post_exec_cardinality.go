//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-sqlc
//ff:what XQS-76 — @get/@post 시퀀스가 :exec/:execrows/:execresult sqlc 쿼리를 참조하면 ERROR (assignment mismatch 방지)

package ssac_sqlc

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xqs76GetPostExecCardinality validates XQS-76: @get and @post sequences
// generate `var, err := qtx.Method(...)` (two return values). If the referenced
// sqlc query has cardinality :exec, :execrows, or :execresult, the generated
// code will have an assignment mismatch.
func xqs76GetPostExecCardinality(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil || len(fs.ServiceFuncs) == 0 || len(fs.SQLcQueries) == 0 {
		return nil
	}
	queryMap := buildQueryBodyMap(fs)
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if d, ok := checkXqs76Seq(fn, seq, queryMap); ok {
				diags = append(diags, d)
			}
		}
	}
	return diags
}
