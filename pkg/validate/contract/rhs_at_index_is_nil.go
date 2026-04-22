//ff:func feature=validate-contract type=util control=sequence topic=preserve-safety
//ff:what rhsAtIndexIsNil — AssignStmt RHS[idx] 가 literal nil 인지 판정

package contract

import "go/ast"

// rhsAtIndexIsNil returns true when RHS[idx] exists and is the bare
// identifier `nil`. Used by assignClobbersErr to tell a genuine
// "forget the error" rebinding (`err = ...anything...`) apart from
// the documented-discard pattern `err = nil`.
func rhsAtIndexIsNil(assign *ast.AssignStmt, idx int) bool {
	if assign == nil || idx < 0 || idx >= len(assign.Rhs) {
		return false
	}
	ident, ok := assign.Rhs[idx].(*ast.Ident)
	if !ok {
		return false
	}
	return ident.Name == "nil"
}
