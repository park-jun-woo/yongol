//ff:func feature=validate-contract type=util control=sequence topic=preserve-safety
//ff:what ifStmtClosesVar — IfStmt body 내에 varName 에 대한 Close 호출 존재 여부

package contract

import "go/ast"

// ifStmtClosesVar recurses into stmt.Body when stmt is an *ast.IfStmt,
// reusing deferBlockClosesVar for the nested walk. Non-if statements
// or ifs without bodies return false immediately. This split keeps
// deferBlockClosesVar itself at nesting depth 2 — required by Q1.
func ifStmtClosesVar(stmt ast.Stmt, varName string) bool {
	ifs, ok := stmt.(*ast.IfStmt)
	if !ok || ifs.Body == nil {
		return false
	}
	return deferBlockClosesVar(ifs.Body.List, varName)
}
