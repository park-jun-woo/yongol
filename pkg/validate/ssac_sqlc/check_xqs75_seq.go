//ff:func feature=validate type=util control=sequence topic=ssac-sqlc
//ff:what checkXqs75Seq — 단일 시퀀스의 XQS-75 (@put/@delete + :one/:many) 위반 평가

package ssac_sqlc

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// checkXqs75Seq evaluates a single sequence against XQS-75.
// Returns (diag, true) when a mismatch is detected, otherwise (zero, false).
func checkXqs75Seq(
	fn ssacparser.ServiceFunc,
	seq ssacparser.Sequence,
	queryMap map[string]sqlcparser.QuerySpec,
) (diagnostic.Diagnostic, bool) {
	if seq.Type != ssacparser.SeqPut && seq.Type != ssacparser.SeqDelete {
		return diagnostic.Diagnostic{}, false
	}
	if seq.Package != "" {
		return diagnostic.Diagnostic{}, false
	}
	queryName := resolveQueryName(seq)
	q, ok := queryMap[queryName]
	if !ok {
		return diagnostic.Diagnostic{}, false
	}
	if q.Cardinality == "exec" || q.Cardinality == "execresult" {
		return diagnostic.Diagnostic{}, false
	}
	return diagnostic.Diagnostic{
		File:  fn.FileName,
		Line:  seq.Line,
		Phase: diagnostic.PhaseValidate,
		Level: diagnostic.LevelError,
		Message: fmt.Sprintf(
			"[XQS-75] @%s expects :exec query but %s is :%s",
			seq.Type, queryName, q.Cardinality,
		),
		Advice: "Change sqlc query to :exec (remove RETURNING), or use @get for :one queries",
	}, true
}
