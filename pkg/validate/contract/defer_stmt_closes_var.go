//ff:func feature=validate-contract type=util control=sequence topic=preserve-safety
//ff:what deferStmtClosesVar — 단일 stmt 가 defer Close (직접 또는 closure) 인지 판정

package contract

import "go/ast"

// deferStmtClosesVar returns true when stmt is either:
//
//   - `defer varName.Close()` (direct call), or
//   - `defer func() { ...; varName.Close(); ... }()` where the closure
//     body contains a direct Close on varName (nested ifs tolerated
//     via deferBlockClosesVar).
//
// Non-defer statements always return false.
func deferStmtClosesVar(stmt ast.Stmt, varName string) bool {
	def, ok := stmt.(*ast.DeferStmt)
	if !ok {
		return false
	}
	if callClosesVar(def.Call, varName) {
		return true
	}
	fn, ok := def.Call.Fun.(*ast.FuncLit)
	if !ok || fn.Body == nil {
		return false
	}
	return deferBlockClosesVar(fn.Body.List, varName)
}
