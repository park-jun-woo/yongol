//ff:func feature=validate type=rule control=iteration dimension=3 topic=ssac-sqlc
//ff:what XQS-20 — SSaC 선언 반환 타입 ↔ sqlc 쿼리 RETURNING shape cross-check (ERROR)

package ssac_sqlc

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xqs20ReturnTypeMatch validates XQS-20: when a SSaC `@get` / `@post` / `@put`
// sequence declares a return type, that declared type must match the sqlc
// emission shape for the underlying query. sqlc emits the model directly for
// `RETURNING *` (or full RETURNING), and a synthetic `<QueryName>Row` for
// partial RETURNING. A mismatch surfaces only at `go build` time without this
// rule.
//
// Skip conditions:
//   - seq.Type not in {@get, @post, @put}
//   - seq.Result == nil (no declared output, e.g. `@put` without binding)
//   - seq.Package != "" (`@call` etc. — non-sqlc target)
//   - query Cardinality == "exec" / "execresult" (no row returned)
//   - query lookup miss (XQS-19 / S-49 cover that)
func xqs20ReturnTypeMatch(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil || len(fs.ServiceFuncs) == 0 || len(fs.SQLcQueries) == 0 {
		return nil
	}
	queryMap := buildQueryBodyMap(fs)
	tableMap := buildDDLTableLookup(fs)
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			diag, ok := checkXqs20Seq(fn, seq, queryMap, tableMap)
			if ok {
				diags = append(diags, diag)
			}
		}
	}
	return diags
}
