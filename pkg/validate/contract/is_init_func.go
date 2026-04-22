//ff:func feature=validate-contract type=util control=sequence topic=preserve-safety
//ff:what isInitFunc — *ast.FuncDecl 이 패키지 init() 함수인지 판정

package contract

import "go/ast"

// isInitFunc reports whether fn is the special `func init() { ... }`
// package-level initializer. init() is an allowlisted location for
// panic(...) — Go libraries conventionally fail fast there when
// configuration is invalid. Methods (non-nil Recv) are never init.
func isInitFunc(fn *ast.FuncDecl) bool {
	if fn == nil || fn.Name == nil {
		return false
	}
	if fn.Recv != nil {
		return false
	}
	return fn.Name.Name == "init"
}
