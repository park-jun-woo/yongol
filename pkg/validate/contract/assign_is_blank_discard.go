//ff:func feature=validate-contract type=util control=iteration dimension=1 topic=preserve-safety
//ff:what assignIsBlankDiscard — AssignStmt LHS 가 모두 blank identifier 인지 판정

package contract

import "go/ast"

// assignIsBlankDiscard reports whether every LHS of as is the blank
// identifier `_`. That is the canonical Go idiom for "I know this
// returns an error and I am intentionally dropping it" — we accept it
// as an explicit choice, no PRV-12/13 diagnostic emitted.
func assignIsBlankDiscard(as *ast.AssignStmt) bool {
	if as == nil || len(as.Lhs) == 0 {
		return false
	}
	for _, lhs := range as.Lhs {
		ident, ok := lhs.(*ast.Ident)
		if !ok || ident.Name != "_" {
			return false
		}
	}
	return true
}
