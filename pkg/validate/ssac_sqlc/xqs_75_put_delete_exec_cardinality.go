//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-sqlc
//ff:what XQS-75 — @put/@delete 시퀀스가 :one/:many sqlc 쿼리를 참조하면 ERROR (assignment mismatch 방지)

package ssac_sqlc

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xqs75PutDeleteExecCardinality validates XQS-75: @put and @delete sequences
// generate `err = qtx.Method(...)` (single return value). If the referenced
// sqlc query has cardinality :one or :many, the generated code will have an
// assignment mismatch.
func xqs75PutDeleteExecCardinality(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil || len(fs.ServiceFuncs) == 0 || len(fs.SQLcQueries) == 0 {
		return nil
	}
	queryMap := buildQueryBodyMap(fs)
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if d, ok := checkXqs75Seq(fn, seq, queryMap); ok {
				diags = append(diags, d)
			}
		}
	}
	return diags
}
