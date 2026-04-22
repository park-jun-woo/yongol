//ff:func feature=validate-contract type=util control=selection topic=preserve-safety
//ff:what scanCallInStmt — stmt 가 Scan 호출이면 CallExpr/errName/kind 반환

package contract

import "go/ast"

// scanCallInStmt inspects stmt and reports whether it is a call to
// `<row|rows|r>.Scan(...)`. The match keys off the method name plus a
// heuristic on the receiver identifier — preserved handler bodies
// conventionally name row/rows after sqlc queries, so this stays
// precise without a full type-checker.
//
// Returns:
//   - nil CallExpr when the statement is not a Scan statement.
//   - kind=Discarded for `_ = row.Scan(...)` or the init of an
//     `if err := row.Scan(...); err != nil { ... }` statement.
//   - kind=Assigned with errName when assigned to a named err ident.
func scanCallInStmt(stmt ast.Stmt) (*ast.CallExpr, string, scanKind) {
	switch s := stmt.(type) {
	case *ast.IfStmt:
		return scanCallInIfStmt(s)
	case *ast.AssignStmt:
		return scanCallInAssignStmt(s)
	case *ast.ExprStmt:
		return scanCallInExprStmt(s)
	}
	return nil, "", scanKindUnknown
}
