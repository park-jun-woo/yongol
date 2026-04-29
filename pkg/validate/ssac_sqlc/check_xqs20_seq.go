//ff:func feature=validate type=util control=sequence topic=ssac-sqlc
//ff:what checkXqs20Seq — 단일 SSaC 시퀀스에 대한 XQS-20 판정 (skip / pass / diag)

package ssac_sqlc

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// checkXqs20Seq evaluates a single sequence against XQS-20.
//
// Returns (diag, true) when a mismatch is detected, otherwise (zero, false).
// Skip cases are treated as "no diagnostic" (returns false).
func checkXqs20Seq(
	fn ssac.ServiceFunc,
	seq ssac.Sequence,
	queryMap map[string]sqlcparser.QuerySpec,
	tableMap map[string]*ddl.Table,
) (diagnostic.Diagnostic, bool) {
	if !xqs20EligibleSeqType(seq.Type) {
		return diagnostic.Diagnostic{}, false
	}
	if seq.Package != "" {
		return diagnostic.Diagnostic{}, false
	}
	if seq.Result == nil || seq.Result.Type == "" {
		return diagnostic.Diagnostic{}, false
	}
	// Wrappers (`Page[T]`, `Cursor[T]`, `[]T`) — list shapes — are out of
	// scope for this Phase. The element type alone is not enough to decide
	// whether the underlying query emits a Row or model wrapper, and
	// pagination wiring lives in a separate code path.
	if seq.Result.Wrapper != "" {
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

	clause := extractReturningClause(q.Body)
	tableName := modelToTableName(q.Model)
	shape := classifyReturningShape(clause, tableMap[tableName])
	// SELECT bodies without RETURNING (ShapeNone) are out of scope for this
	// rule: a SELECT may project the model or any subset (UserSummaryRow etc.)
	// and yongol cannot infer the intended emission shape from the query body
	// alone. XDS-12 covers result-type ↔ DDL coverage on the SELECT path.
	if shape == ShapeNone {
		return diagnostic.Diagnostic{}, false
	}
	expected := expectedSsacReturnType(shape, q.Model, queryName)
	declared := seq.Result.Type
	if declared == expected {
		return diagnostic.Diagnostic{}, false
	}
	reason := formatReturningReason(shape, clause)
	return buildXqs20Diag(fn, seq, declared, expected, queryName, shape, reason), true
}
