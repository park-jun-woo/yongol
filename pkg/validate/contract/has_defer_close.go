//ff:func feature=validate-contract type=util control=iteration dimension=1 topic=preserve-safety
//ff:what hasDeferClose — 블록 내에서 `defer <varName>.Close()` 존재 여부 판정

package contract

import "go/ast"

// hasDeferClose reports whether any stmt in stmts is
// `defer varName.Close()` or `defer func(){ varName.Close() }()`. It is
// intentionally lenient — any method call named "Close" on an
// identifier whose name equals varName counts, including chained calls
// such as `defer varName.Body.Close()` (Body is always safe to close
// when varName is non-nil — that is the stdlib http.Response idiom).
func hasDeferClose(stmts []ast.Stmt, varName string) bool {
	for _, s := range stmts {
		if deferStmtClosesVar(s, varName) {
			return true
		}
	}
	return false
}
