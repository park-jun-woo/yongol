//ff:func feature=validate-contract type=util control=iteration dimension=1 topic=preserve-safety
//ff:what deferBlockClosesVar — defer func(){...}() 블록 body 내 Close 호출 존재 여부

package contract

import "go/ast"

// deferBlockClosesVar scans stmts (the body of a `defer func() { ... }()`
// literal) for a direct Close call on varName. Nested control flow
// (ifs, selects) is recursed into one level so the common pattern
// `defer func() { if x != nil { x.Close() } }()` is recognised.
func deferBlockClosesVar(stmts []ast.Stmt, varName string) bool {
	for _, s := range stmts {
		if exprStmtClosesVar(s, varName) {
			return true
		}
		if ifStmtClosesVar(s, varName) {
			return true
		}
	}
	return false
}
