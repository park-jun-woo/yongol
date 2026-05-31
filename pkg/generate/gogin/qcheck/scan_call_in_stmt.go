//ff:func feature=gen-gogin type=util control=sequence
//ff:what scanCallInStmt — 단일 Stmt 에서 `<ident>.Scan(...)` 호출과 가드 여부 판정

package qcheck

import (
	"go/ast"
	"go/token"
)

// scanCallInStmt inspects stmt and returns a finding when a `.Scan(...)`
// call is unguarded. The same three shapes accepted by unmarshalInStmt
// are tolerated:
//   - `if err := x.Scan(...); err != nil { ... }`
//   - `err := x.Scan(...)` followed by err-guard
//   - `if err = x.Scan(...); err != nil { ... }` via assignCallsAndGuarded
//
// pkgName is empty — any receiver ident is accepted.
func scanCallInStmt(stmt ast.Stmt, blockList []ast.Stmt, i int, fset *token.FileSet) []DefensiveFinding {
	if ifStmt, ok := stmt.(*ast.IfStmt); ok && ifInitCalls(ifStmt, "", "Scan") {
		return nil
	}
	if assign, ok := stmt.(*ast.AssignStmt); ok && assignCallsAndGuarded(assign, "", "Scan", blockList, i) {
		return nil
	}
	call := findCallInStmt(stmt, "", "Scan")
	if call == nil {
		return nil
	}
	sel := call.Fun.(*ast.SelectorExpr)
	recv := sel.X.(*ast.Ident).Name
	return []DefensiveFinding{{
		Category: "DF-02",
		Detail:   recv + ".Scan",
		Pos:      fset.Position(call.Pos()),
	}}
}
