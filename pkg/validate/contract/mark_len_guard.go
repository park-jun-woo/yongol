//ff:func feature=validate-contract type=util control=sequence topic=preserve-safety
//ff:what markLenGuard — CallExpr 가 len(x) 이면 x 를 guarded 집합에 추가, 반환값은 매칭 여부

package contract

import "go/ast"

// markLenGuard inspects call; when it is `len(x)` with x a plain
// identifier, x is marked in guarded. Returns true on match so the
// walking scanner can short-circuit nested traversal. Calls whose
// argument is not a simple ident (e.g. `len(m[k])`) are ignored
// because the indexed base — what PRV-14 actually cares about —
// cannot be resolved by name alone.
func markLenGuard(call *ast.CallExpr, guarded map[string]bool) bool {
	if call == nil || len(call.Args) != 1 {
		return false
	}
	fn, ok := call.Fun.(*ast.Ident)
	if !ok || fn.Name != "len" {
		return false
	}
	name := leftmostIdentName(call.Args[0])
	if name == "" {
		return false
	}
	guarded[name] = true
	return true
}
