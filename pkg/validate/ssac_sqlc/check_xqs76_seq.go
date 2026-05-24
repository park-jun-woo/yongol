//ff:func feature=validate type=util control=sequence topic=ssac-sqlc
//ff:what checkXqs76Seq — 단일 시퀀스의 XQS-76 (@get/@post + :exec) 위반 평가

package ssac_sqlc

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// checkXqs76Seq evaluates a single sequence against XQS-76.
// Returns (diag, true) when a mismatch is detected, otherwise (zero, false).
func checkXqs76Seq(
	fn ssacparser.ServiceFunc,
	seq ssacparser.Sequence,
	queryMap map[string]sqlcparser.QuerySpec,
) (diagnostic.Diagnostic, bool) {
	if seq.Type != ssacparser.SeqGet && seq.Type != ssacparser.SeqPost {
		return diagnostic.Diagnostic{}, false
	}
	if seq.Package != "" {
		return diagnostic.Diagnostic{}, false
	}
	if seq.Result == nil {
		return diagnostic.Diagnostic{}, false
	}
	queryName := resolveQueryName(seq)
	q, ok := queryMap[queryName]
	if !ok {
		return diagnostic.Diagnostic{}, false
	}
	if q.Cardinality == "one" || q.Cardinality == "many" {
		return diagnostic.Diagnostic{}, false
	}
	return diagnostic.Diagnostic{
		File:  fn.FileName,
		Line:  seq.Line,
		Phase: diagnostic.PhaseValidate,
		Level: diagnostic.LevelError,
		Message: fmt.Sprintf(
			"[XQS-76] @%s expects :one/:many query but %s is :%s",
			seq.Type, queryName, q.Cardinality,
		),
		Advice: "Change sqlc query to :one (add RETURNING), or use @put for :exec queries",
	}, true
}
