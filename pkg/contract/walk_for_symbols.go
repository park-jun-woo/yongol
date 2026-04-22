//ff:func feature=contract type=util control=sequence
//ff:what walkForSymbols — body 를 두 번 walk: CallExpr selector 먼저, 그 외 SelectorExpr 는 field 로 분류

package contract

import (
	"go/ast"
	"go/token"
)

// walkForSymbols inspects every node in body and emits unique,
// alphabetically ordered symbol lists for the three categories.
// CallExpr are visited before their Fun so selectors used as call
// targets route to CallTargets / SqlcQueries rather than DDLFields.
func walkForSymbols(fset *token.FileSet, body *ast.BlockStmt) ExternalSymbols {
	queries := map[string]struct{}{}
	calls := map[string]struct{}{}
	fields := map[string]struct{}{}
	callSelectors := map[*ast.SelectorExpr]struct{}{}

	ast.Inspect(body, func(n ast.Node) bool {
		return collectCallSelector(n, queries, calls, callSelectors)
	})
	ast.Inspect(body, func(n ast.Node) bool {
		return collectFieldSelector(fset, n, callSelectors, fields)
	})

	return ExternalSymbols{
		SqlcQueries: toSortedSlice(queries),
		CallTargets: toSortedSlice(calls),
		DDLFields:   toSortedSlice(fields),
	}
}
