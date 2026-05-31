//ff:func feature=validate-contract type=util control=sequence
//ff:what expectedExternalSymbols — Ground 에서 유효한 sqlc query / @call 대상 / DDL 컬럼(Pascal) 집합 구축

package contract

import (
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// expectedExternalSymbols returns three lookup sets used by PRV-02:
//
//   - queries   : valid sqlc method names (Queries.<Method> / qtx.<Method>).
//     Derived from fs.SQLcQueries (.Name — the raw `-- name:` ident).
//   - calls     : valid `<pkg>.<Func>` call targets. Union of fs.ProjectFuncSpecs
//   - fs.YongolPkgSpecs (canonical spec identity) and SSaC
//     `SSaC.callRef` (normalized) so user handlers can invoke any
//     func surface that SSaC / Func SSOT expose.
//   - ddlFields : valid DDL column names rendered in PascalCase (matches
//     oapi-codegen / sqlc-generated struct field exports).
//
// The symbol sets are intentionally permissive — false positives from
// local/temporary identifiers are acceptable; the function body may
// legitimately touch non-SSOT packages (slog, errors, sql) that the
// contract extractor already filters via callExprPkgDenylist.
func expectedExternalSymbols(fs *yongol.Fullstack, g *rule.Ground) (queries, calls, ddlFields map[string]bool) {
	queries = buildExpectedQueries(fs)
	calls = buildExpectedCalls(fs, g)
	ddlFields = buildExpectedDDLFields(fs)
	return queries, calls, ddlFields
}
